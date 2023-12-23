package example_gen

import (
	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra/gdb"
)

type Audit struct {
	CreatedAt gdb.SqlNullString `db:"created_at"`
	UpdatedAt gdb.SqlNullInt64  `db:"updated_at"`
	CreatedBy string            `db:"created_by"`
	UpdatedBy gdb.SqlNullString `db:"updated_by"`
	DeletedAt gdb.SqlNullInt64  `db:"deleted_at"`
	DeletedBy gdb.SqlNullString `db:"deleted_by"`
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
