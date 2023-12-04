package example

import (
	"github.com/SyaibanAhmadRamadhan/gocatch/gdb/gsql"
)

type Audit struct {
	CreatedAt gsql.NullString `db:"created_at"`
	UpdatedAt gsql.NullInt64  `db:"updated_at"`
	CreatedBy string          `db:"created_by"`
	UpdatedBy gsql.NullString `db:"updated_by"`
	DeletedAt gsql.NullInt64  `db:"deleted_at"`
	DeletedBy gsql.NullString `db:"deleted_by"`
}

type User struct {
	ID          string `db:"id"`
	RoleID      int    `db:"role_id"`
	Username    string `db:"username"`
	Email       string `db:"email"`
	Password    string `db:"password"`
	PhoneNumber string `db:"phone_number"`
	Audit       `db:"-"`
}
