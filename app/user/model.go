package user

import (
	"database/sql"
	"net/http"
	"github.com/yesseneon/onlineshop/config"
	"github.com/yesseneon/onlineshop/helper"
	"regexp"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User struct
type User struct {
	ID        int       `db:"id"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	Role      int       `db:"role"`
	Password2 string
	Errors    map[string]string
}

func (u *User) validate(isRgs bool) (bool, error) {
	u.Errors = make(map[string]string)
	email := strings.TrimSpace(u.Email)
	var eu User
	var err error

	match := regexp.MustCompile(".+@.+\\..+").Match([]byte(email))
	if match {
		eu, err = getExistingUser(email)
		if err != nil {
			return false, err
		}
	} else {
		u.Errors["Email"] = "Please enter a valid email address"
	}

	// if it's registration
	if isRgs {
		fname := strings.TrimSpace(u.FirstName)
		lname := strings.TrimSpace(u.LastName)

		if fname == "" || len(fname) > 20 {
			u.Errors["FirstName"] = "The field First Name must be a string with a maximum length of 20"
		}

		if lname == "" || len(lname) > 20 {
			u.Errors["LastName"] = "The field Last Name must be a string with a maximum length of 20"
		}

		if eu.ID > 0 {
			u.Errors["Email"] = "The email address is already taken. Please choose another one"
		}

		if len(u.Password) < 6 || len(u.Password) > 20 {
			u.Errors["Password"] = "Your password must be 6-20 characters long"
		}

		if u.Password != u.Password2 {
			u.Errors["Password2"] = "The specified passwords do not match"
		}
	} else {
		u.ID = eu.ID

		if match && eu.ID == 0 {
			u.Errors["Email"] = "User not found with this email address"
		}

		// Does the entered password match the stored password?
		err := bcrypt.CompareHashAndPassword([]byte(eu.Password), []byte(u.Password))
		if err != nil {
			u.Errors["Password"] = "Password does not match, please try again"
		}
	}

	return len(u.Errors) == 0, nil
}

func (u *User) create() (int, error) {
	var id int
	stm := "INSERT INTO users (first_name, last_name, email, password, created_at, updated_at, role) VALUES ($1, $2, $3, $4, NOW()::timestamp(0), NOW()::timestamp(0), 0) RETURNING id"
	err := config.DB.QueryRow(stm, u.FirstName, u.LastName, u.Email, u.Password).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func getExistingUser(email string) (User, error) {
	var eu User // existing user
	row := config.DB.QueryRow("SELECT * FROM users WHERE email=$1", email)
	err := row.Scan(&eu.ID, &eu.FirstName, &eu.LastName, &eu.Email, &eu.Password, &eu.CreatedAt, &eu.UpdatedAt, &eu.Role)
	if err != nil && err != sql.ErrNoRows {
		return eu, err
	}
	return eu, nil
}

func createSession(sessionID string, userID int) error {
	_, err := config.DB.Exec("DELETE FROM sessions WHERE user_id=$1", userID)
	if err != nil {
		return err
	}

	stm := "INSERT INTO sessions (session_id, user_id, created_at, updated_at) VALUES ($1, $2, NOW()::timestamp(0), NOW()::timestamp(0))"
	_, err = config.DB.Exec(stm, sessionID, userID)
	if err != nil {
		return err
	}
	return nil
}

func deleteSession(sessionID string) error {
	_, err := config.DB.Exec("DELETE FROM sessions WHERE session_id=$1", sessionID)
	if err != nil {
		return err
	}
	return nil
}

func SessionExists(r *http.Request) (bool, error) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return false, nil
	}

	var result bool
	stm := "SELECT EXISTS(SELECT 1 FROM users WHERE id=(SELECT user_id FROM sessions WHERE session_id=$1 LIMIT 1))"
	err = config.DB.QueryRow(stm, cookie.Value).Scan(&result)
	if err != nil {
		return false, err
	}
	return result, nil
}

func GetAuthUser(r *http.Request) (helper.Auth, error) {
	auth := helper.Auth{}
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return auth, nil
	}

	row := config.DB.QueryRow("SELECT * FROM users WHERE id=(SELECT user_id FROM sessions WHERE session_id=$1)", cookie.Value)
	err = row.Scan(&auth.ID, &auth.FirstName, &auth.LastName, &auth.Email, &auth.Password, &auth.CreatedAt, &auth.UpdatedAt, &auth.Role)
	if err != nil && err != sql.ErrNoRows {
		return auth, err
	}

	return auth, nil
}
