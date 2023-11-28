package exGen

import (
	"github.com/SyaibanAhmadRamadhan/jolly/Jsql"
)

type Audit struct {
	CreatedAt Jsql.NullString `db:"created_at"`
	UpdatedAt Jsql.NullInt64  `db:"updated_at"`
	CreatedBy string          `db:"created_by"`
	UpdatedBy Jsql.NullString `db:"updated_by"`
	DeletedAt Jsql.NullInt64  `db:"deleted_at"`
	DeletedBy Jsql.NullString `db:"deleted_by"`
}

type User struct {
	ID               string `db:"id"`
	RoleID           int    `db:"role_id"`
	Username         string `db:"username"`
	Email            string `db:"email"`
	Password         string `db:"password"`
	PhoneNumber      string `db:"phone_number"`
	Audit            `db:"-"`
	QColumnFields    []string
	QFilterNamedArgs Jsql.QFilterNamedArgs
}
