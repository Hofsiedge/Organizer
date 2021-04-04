package user

import "github.com/jackc/pgtype"

// User model
type User struct {
	Id           int64
	Email        pgtype.Text
	Login        pgtype.Text
	PasswordHash pgtype.Text
}
