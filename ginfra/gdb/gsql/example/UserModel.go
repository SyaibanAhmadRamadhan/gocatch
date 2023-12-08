package example

type Audit struct {
	CreatedAt Gsql.NullString `db:"created_at"`
	UpdatedAt Gsql.NullInt64  `db:"updated_at"`
	CreatedBy string          `db:"created_by"`
	UpdatedBy Gsql.NullString `db:"updated_by"`
	DeletedAt Gsql.NullInt64  `db:"deleted_at"`
	DeletedBy Gsql.NullString `db:"deleted_by"`
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
