package models

type User struct {
	ID       int
	Username string
	Email    string
	Password string // Note: In production, always store hashed passwords
}
