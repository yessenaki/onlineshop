package helper

import "time"

type key string

const ContextDataKey key = "cdk"

type Auth struct {
	ID        int       `db:"id"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type ContextData struct {
	Auth    Auth
	ItemQnt int
}
