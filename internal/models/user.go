package models

import "time"

type User struct {
	UUID         string    `db:"uuid"`
	Username     string    `db:"username"`
	Email        string    `db:"email"`
	PasswordHash string    `db:"password_hash,omitempty"`
	Active       bool      `db:"active"`
	CreatedAt    time.Time `db:"created_at"`
}
