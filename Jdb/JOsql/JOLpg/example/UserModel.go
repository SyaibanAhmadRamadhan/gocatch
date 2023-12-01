package example

import (
	"github.com/SyaibanAhmadRamadhan/jolly/Jdb/JOsql"
	"github.com/SyaibanAhmadRamadhan/jolly/Jtype/JOmap"
)

type Audit struct {
	CreatedAt JOsql.NullString `db:"created_at"`
	UpdatedAt JOsql.NullInt64  `db:"updated_at"`
	CreatedBy string           `db:"created_by"`
	UpdatedBy JOsql.NullString `db:"updated_by"`
	DeletedAt JOsql.NullInt64  `db:"deleted_at"`
	DeletedBy JOsql.NullString `db:"deleted_by"`
}

type User struct {
	ID          string `db:"id"`
	RoleID      int    `db:"role_id"`
	Username    string `db:"username"`
	Email       string `db:"email"`
	Password    string `db:"password"`
	PhoneNumber string `db:"phone_number"`
	Audit       `db:"-"`
	RQField     []string
	WCField     JOmap.SA
	FNamedArgs  JOsql.QFilterNamedArgs
}
