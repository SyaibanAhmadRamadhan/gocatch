// DO NOT EDIT, will be overwritten by https://github.com/SyaibanAhmadRamadhan/gocatch/blob/main/ginfra/gdb/generator.go. 

package example_gen

import (
	"errors"

	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra/gdb"
)

// UserTableName this table or collection name
const UserTableName string = "user"

// UserSchemaName is a variable schema name
const UserSchemaName string = "public"

// NewUser is a struct with pointer that represents the table User in the database.
func NewUser() *User {
	return &User{}
}

// NewUserWithOutPtr is a struct without pointer that represents the table User in the database.
func NewUserWithOutPtr() User {
	return User{}
}

// FieldID is a field or column in the table User.
func (u *User) FieldID() string {
	return "id"
}

// SetID is a setter for the field or column ID in the table User.
func (u *User) SetID(param string) string {
	u.ID = param
	return "id"
}

// FieldEmail is a field or column in the table User.
func (u *User) FieldEmail() string {
	return "email"
}

// SetEmail is a setter for the field or column Email in the table User.
func (u *User) SetEmail(param string) string {
	u.Email = param
	return "email"
}

// FieldCreatedAt is a field or column in the table User.
func (u *User) FieldCreatedAt() string {
	return "created_at"
}

// SetCreatedAt is a setter for the field or column CreatedAt in the table User.
func (u *User) SetCreatedAt(param string) string {
	u.CreatedAt = gdb.NewSqlNullString(&param)
	return "created_at"
}

// FieldUpdatedBy is a field or column in the table User.
func (u *User) FieldUpdatedBy() string {
	return "updated_by"
}

// SetUpdatedBy is a setter for the field or column UpdatedBy in the table User.
func (u *User) SetUpdatedBy(param string) string {
	u.UpdatedBy = gdb.NewSqlNullString(&param)
	return "updated_by"
}

// FieldDeletedAt is a field or column in the table User.
func (u *User) FieldDeletedAt() string {
	return "deleted_at"
}

// SetDeletedAt is a setter for the field or column DeletedAt in the table User.
func (u *User) SetDeletedAt(param int64) string {
	u.DeletedAt = gdb.NewSqlNullInt64(&param)
	return "deleted_at"
}

// FieldDeletedBy is a field or column in the table User.
func (u *User) FieldDeletedBy() string {
	return "deleted_by"
}

// SetDeletedBy is a setter for the field or column DeletedBy in the table User.
func (u *User) SetDeletedBy(param string) string {
	u.DeletedBy = gdb.NewSqlNullString(&param)
	return "deleted_by"
}

// FieldRoleID is a field or column in the table User.
func (u *User) FieldRoleID() string {
	return "role_id"
}

// SetRoleID is a setter for the field or column RoleID in the table User.
func (u *User) SetRoleID(param int) string {
	u.RoleID = param
	return "role_id"
}

// FieldUsername is a field or column in the table User.
func (u *User) FieldUsername() string {
	return "username"
}

// SetUsername is a setter for the field or column Username in the table User.
func (u *User) SetUsername(param string) string {
	u.Username = param
	return "username"
}

// FieldPassword is a field or column in the table User.
func (u *User) FieldPassword() string {
	return "password"
}

// SetPassword is a setter for the field or column Password in the table User.
func (u *User) SetPassword(param string) string {
	u.Password = param
	return "password"
}

// FieldPhoneNumber is a field or column in the table User.
func (u *User) FieldPhoneNumber() string {
	return "phone_number"
}

// SetPhoneNumber is a setter for the field or column PhoneNumber in the table User.
func (u *User) SetPhoneNumber(param string) string {
	u.PhoneNumber = param
	return "phone_number"
}

// FieldUpdatedAt is a field or column in the table User.
func (u *User) FieldUpdatedAt() string {
	return "updated_at"
}

// SetUpdatedAt is a setter for the field or column UpdatedAt in the table User.
func (u *User) SetUpdatedAt(param int64) string {
	u.UpdatedAt = gdb.NewSqlNullInt64(&param)
	return "updated_at"
}

// FieldCreatedBy is a field or column in the table User.
func (u *User) FieldCreatedBy() string {
	return "created_by"
}

// SetCreatedBy is a setter for the field or column CreatedBy in the table User.
func (u *User) SetCreatedBy(param string) string {
	u.CreatedBy = param
	return "created_by"
}

// AllField is a function to get all field or column in the table User.
func (u *User) AllField() (str []string) {
	str = []string{ 
		`deleted_by`,
		`id`,
		`email`,
		`created_at`,
		`updated_by`,
		`deleted_at`,
		`created_by`,
		`role_id`,
		`username`,
		`password`,
		`phone_number`,
		`updated_at`,
	}
	return
}

// OrderFields is a function to get all field or column in the table User.
func (u *User) OrderFields() (str []string) {
	str = []string{ 
		`id`,
		`role_id`,
		`username`,
	}
	return
}

// GetValuesByColums is a function to get all value by column in the table User.
func (u *User) GetValuesByColums(columns ...string) []any {
	var values []any
	for _, column := range columns {
		switch column {
		case u.FieldDeletedAt():
			values = append(values, u.DeletedAt)
		case u.FieldDeletedBy():
			values = append(values, u.DeletedBy)
		case u.FieldID():
			values = append(values, u.ID)
		case u.FieldEmail():
			values = append(values, u.Email)
		case u.FieldCreatedAt():
			values = append(values, u.CreatedAt)
		case u.FieldUpdatedBy():
			values = append(values, u.UpdatedBy)
		case u.FieldUpdatedAt():
			values = append(values, u.UpdatedAt)
		case u.FieldCreatedBy():
			values = append(values, u.CreatedBy)
		case u.FieldRoleID():
			values = append(values, u.RoleID)
		case u.FieldUsername():
			values = append(values, u.Username)
		case u.FieldPassword():
			values = append(values, u.Password)
		case u.FieldPhoneNumber():
			values = append(values, u.PhoneNumber)
		}
	}
	return values
}

// ScanMap is a function to scan the value with for rows.Value() from the database to the struct User.
func (u *User) ScanMap(data map[string]any) (err error) {
	for key, value := range data {
		switch key {
		case u.FieldCreatedAt():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field CreatedAt")
			}
			u.SetCreatedAt(val)
		case u.FieldUpdatedBy():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field UpdatedBy")
			}
			u.SetUpdatedBy(val)
		case u.FieldDeletedAt():
			val, ok := value.(int64)
			if !ok {
				return errors.New("invalid type int64. field DeletedAt")
			}
			u.SetDeletedAt(val)
		case u.FieldDeletedBy():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field DeletedBy")
			}
			u.SetDeletedBy(val)
		case u.FieldID():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field ID")
			}
			u.SetID(val)
		case u.FieldEmail():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field Email")
			}
			u.SetEmail(val)
		case u.FieldPassword():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field Password")
			}
			u.SetPassword(val)
		case u.FieldPhoneNumber():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field PhoneNumber")
			}
			u.SetPhoneNumber(val)
		case u.FieldUpdatedAt():
			val, ok := value.(int64)
			if !ok {
				return errors.New("invalid type int64. field UpdatedAt")
			}
			u.SetUpdatedAt(val)
		case u.FieldCreatedBy():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field CreatedBy")
			}
			u.SetCreatedBy(val)
		case u.FieldRoleID():
			val, ok := value.(int)
			if !ok {
				return errors.New("invalid type int. field RoleID")
			}
			u.SetRoleID(val)
		case u.FieldUsername():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field Username")
			}
			u.SetUsername(val)
		default:
			return errors.New("invalid column")
		}
	}
	return nil
}

