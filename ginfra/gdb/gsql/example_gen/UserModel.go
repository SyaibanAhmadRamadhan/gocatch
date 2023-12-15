package example_gen

import (
	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra/gdb/gsql"
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
	ID          string `db:"id"        order:"true"`
	RoleID      int    `db:"role_id"   order:"true"`
	Username    string `db:"username"  order:"true"`
	Email       string `db:"email"     order:"false"`
	Password    string `db:"password"`
	PhoneNumber string `db:"phone_number"`
	Audit       `db:"-"`
}
