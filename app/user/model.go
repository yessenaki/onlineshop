package user

import (
	"database/sql"
	"net/http"
	"onlineshop/config"
)

// User struct
type User struct {
	ID        int    `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

// Validator struct
type Validator struct {
	User   *User
	Errors map[string]string
}

func (u *User) createUser() (int, error) {
	var lastInsertID int
	sqlStatement := "INSERT INTO users (first_name, last_name, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, NOW()::timestamp(0), NOW()::timestamp(0)) RETURNING id"
	err := config.DB.QueryRow(sqlStatement, u.FirstName, u.LastName, u.Email, u.Password).Scan(&lastInsertID)
	if err != nil {
		return 0, err
	}
	return lastInsertID, nil
}

func (u *User) exists() (bool, error) {
	row := config.DB.QueryRow("SELECT id, password FROM users WHERE email = $1", u.Email)
	err := row.Scan(&u.ID, &u.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func createSession(sessionID string, userID int) error {
	_, err := config.DB.Exec("DELETE FROM sessions WHERE user_id = $1", userID)
	if err != nil {
		return err
	}

	sqlStatement := "INSERT INTO sessions (session_id, user_id, created_at, updated_at) VALUES ($1, $2, NOW()::timestamp(0), NOW()::timestamp(0))"
	_, err = config.DB.Exec(sqlStatement, sessionID, userID)
	if err != nil {
		return err
	}
	return nil
}

func deleteSession(sessionID string) error {
	_, err := config.DB.Exec("DELETE FROM sessions WHERE session_id = $1", sessionID)
	if err != nil {
		return err
	}
	return nil
}

func sessionExists(r *http.Request) (bool, error) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return false, nil
	}

	var result bool
	sqlStatement := "SELECT EXISTS(SELECT 1 FROM users WHERE id = (SELECT user_id FROM sessions WHERE session_id = $1))"
	err = config.DB.QueryRow(sqlStatement, cookie.Value).Scan(&result)
	if err != nil {
		return false, err
	}
	return result, nil
}
