package user

import (
	"log"
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

func userExists(email string) (bool, error) {
	var result bool
	row := config.DB.QueryRow("SELECT EXISTS(SELECT id FROM users WHERE email = $1)", email)
	err := row.Scan(&result)
	if err != nil {
		return false, err
	}
	return result, nil
}

func createUser(user User) (int, error) {
	var lastInsertID int
	sqlStatement := "INSERT INTO users (first_name, last_name, email, password) VALUES ($1, $2, $3, $4) RETURNING id"
	err := config.DB.QueryRow(sqlStatement, user.FirstName, user.LastName, user.Email, user.Password).Scan(&lastInsertID)
	if err != nil {
		log.Fatalln(err)
		return 0, err
	}
	return lastInsertID, nil
}

func createSession(sessionID string, userID int) error {
	sqlStatement := "INSERT INTO sessions (session_id, user_id) VALUES ($1, $2)"
	_, err := config.DB.Exec(sqlStatement, sessionID, userID)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	return nil
}
