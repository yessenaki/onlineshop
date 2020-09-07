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

func createUser(user User) (int, error) {
	var lastInsertID int
	sqlStatement := "INSERT INTO users (first_name, last_name, email, password) VALUES ($1, $2, $3, $4) RETURNING id"
	err := config.DB.QueryRow(sqlStatement, user.FirstName, user.LastName, user.Email, user.Password).Scan(&lastInsertID)
	if err != nil {
		return 0, err
	}
	return lastInsertID, nil
}

func userExists(email string) (User, error) {
	user := User{}
	row := config.DB.QueryRow("SELECT id, email, password FROM users WHERE email = $1", email)
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, nil
		}
		return user, err
	}
	return user, nil
}

func createSession(sessionID string, userID int) error {
	_, err := config.DB.Exec("DELETE FROM sessions WHERE user_id = $1", userID)
	if err != nil {
		return err
	}

	_, err = config.DB.Exec("INSERT INTO sessions (session_id, user_id) VALUES ($1, $2)", sessionID, userID)
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