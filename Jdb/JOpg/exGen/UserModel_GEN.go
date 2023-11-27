package exGen

import (
	"errors"
	"fmt"
	"strings"

	"github.com/SyaibanAhmadRamadhan/jolly/Jsql"
)

// DO NOT EDIT, will be overwritten by https://github.com/SyaibanAhmadRamadhan/jolly/blob/main/Jdb/JOpg/postgres_generator.go. 

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
func (u *User) SetID(param string) {
	u.ID = param
}

// FieldUsername is a field or column in the table User.
func (u *User) FieldUsername() string {
	return "username"
}

// SetUsername is a setter for the field or column Username in the table User.
func (u *User) SetUsername(param string) {
	u.Username = param
}

// FieldEmail is a field or column in the table User.
func (u *User) FieldEmail() string {
	return "email"
}

// SetEmail is a setter for the field or column Email in the table User.
func (u *User) SetEmail(param string) {
	u.Email = param
}

// FieldPassword is a field or column in the table User.
func (u *User) FieldPassword() string {
	return "password"
}

// SetPassword is a setter for the field or column Password in the table User.
func (u *User) SetPassword(param string) {
	u.Password = param
}

// FieldUpdatedAt is a field or column in the table User.
func (u *User) FieldUpdatedAt() string {
	return "updated_at"
}

// SetUpdatedAt is a setter for the field or column UpdatedAt in the table User.
func (u *User) SetUpdatedAt(param Jsql.NullInt64) {
	u.UpdatedAt = param
}

// FieldUpdatedBy is a field or column in the table User.
func (u *User) FieldUpdatedBy() string {
	return "updated_by"
}

// SetUpdatedBy is a setter for the field or column UpdatedBy in the table User.
func (u *User) SetUpdatedBy(param Jsql.NullString) {
	u.UpdatedBy = param
}

// FieldDeletedAt is a field or column in the table User.
func (u *User) FieldDeletedAt() string {
	return "deleted_at"
}

// SetDeletedAt is a setter for the field or column DeletedAt in the table User.
func (u *User) SetDeletedAt(param Jsql.NullInt64) {
	u.DeletedAt = param
}

// FieldRoleID is a field or column in the table User.
func (u *User) FieldRoleID() string {
	return "role_id"
}

// SetRoleID is a setter for the field or column RoleID in the table User.
func (u *User) SetRoleID(param int) {
	u.RoleID = param
}

// FieldPhoneNumber is a field or column in the table User.
func (u *User) FieldPhoneNumber() string {
	return "phone_number"
}

// SetPhoneNumber is a setter for the field or column PhoneNumber in the table User.
func (u *User) SetPhoneNumber(param string) {
	u.PhoneNumber = param
}

// FieldCreatedAt is a field or column in the table User.
func (u *User) FieldCreatedAt() string {
	return "created_at"
}

// SetCreatedAt is a setter for the field or column CreatedAt in the table User.
func (u *User) SetCreatedAt(param Jsql.NullString) {
	u.CreatedAt = param
}

// FieldCreatedBy is a field or column in the table User.
func (u *User) FieldCreatedBy() string {
	return "created_by"
}

// SetCreatedBy is a setter for the field or column CreatedBy in the table User.
func (u *User) SetCreatedBy(param string) {
	u.CreatedBy = param
}

// FieldDeletedBy is a field or column in the table User.
func (u *User) FieldDeletedBy() string {
	return "deleted_by"
}

// SetDeletedBy is a setter for the field or column DeletedBy in the table User.
func (u *User) SetDeletedBy(param Jsql.NullString) {
	u.DeletedBy = param
}

// AllField is a function to get all field or column in the table User.
func (u *User) AllField() (str string) {
	str += `
		updated_by, 
		deleted_at, 
		id, 
		username, 
		email, 
		password, 
		updated_at, 
		role_id, 
		phone_number, 
		created_at, 
		created_by, 
		deleted_by`
	return
}

// Scan is a function to scan the value with for rows.Value() from the database to the struct User.
func (u *User) Scan(value any, columns ...string) (err error) {
	for _, column := range columns {
		switch column {
		case u.FieldDeletedBy():
			val, ok := value.(Jsql.NullString)
			if !ok {
				return errors.New("invalid type")
			}
			u.SetDeletedBy(val)
			return nil
		case u.FieldRoleID():
			val, ok := value.(int)
			if !ok {
				return errors.New("invalid type")
			}
			u.SetRoleID(val)
			return nil
		case u.FieldPhoneNumber():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type")
			}
			u.SetPhoneNumber(val)
			return nil
		case u.FieldCreatedAt():
			val, ok := value.(Jsql.NullString)
			if !ok {
				return errors.New("invalid type")
			}
			u.SetCreatedAt(val)
			return nil
		case u.FieldCreatedBy():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type")
			}
			u.SetCreatedBy(val)
			return nil
		case u.FieldUpdatedAt():
			val, ok := value.(Jsql.NullInt64)
			if !ok {
				return errors.New("invalid type")
			}
			u.SetUpdatedAt(val)
			return nil
		case u.FieldUpdatedBy():
			val, ok := value.(Jsql.NullString)
			if !ok {
				return errors.New("invalid type")
			}
			u.SetUpdatedBy(val)
			return nil
		case u.FieldDeletedAt():
			val, ok := value.(Jsql.NullInt64)
			if !ok {
				return errors.New("invalid type")
			}
			u.SetDeletedAt(val)
			return nil
		case u.FieldID():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type")
			}
			u.SetID(val)
			return nil
		case u.FieldUsername():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type")
			}
			u.SetUsername(val)
			return nil
		case u.FieldEmail():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type")
			}
			u.SetEmail(val)
			return nil
		case u.FieldPassword():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type")
			}
			u.SetPassword(val)
			return nil
		}
	}
	return errors.New("invalid column")
}

// SetColumn is a function to set the column to QueryColumnFields for that will be used in the query.
func (u *User) SetColumn(columns ...string) (err error) {
	for _, column := range columns {
		switch column {
		case u.FieldUpdatedBy():
		case u.FieldDeletedAt():
		case u.FieldID():
		case u.FieldUsername():
		case u.FieldEmail():
		case u.FieldPassword():
		case u.FieldUpdatedAt():
		case u.FieldRoleID():
		case u.FieldPhoneNumber():
		case u.FieldCreatedAt():
		case u.FieldCreatedBy():
		case u.FieldDeletedBy():
		default:
			return errors.New("invalid column")
		}
	}
	u.QueryColumnFields = append(u.QueryColumnFields, columns...)
	return nil
}

// DeleteColumnFromQueryColumnFields is a function to delete the column from QueryColumnFields for that will be used in the query.
func (u *User) DeleteColumnFromQueryColumnFields(elems ...string) (err error) {
	var colums []string
	for _, v := range u.QueryColumnFields{
		colums = append(colums, v)
	}

	for _, elem := range elems {
		index := -1
		for i, column := range u.QueryColumnFields{
			if column == elem {
				index = i
				break
			}
		}
		if index == -1 {
			return fmt.Errorf("column %s not found", elem)
		}
		colums = append(colums[:index], colums[index+1:]...)
	}
	u.QueryColumnFields = colums
	return nil
}

// QueryColumnFieldToStrings is a function to get the column format string from QueryColumnFields for that will be used in the query.
func (u *User) QueryColumnFieldToStrings() (columnStr string) {
	return strings.Join(u.QueryColumnFields, ", ")
}

// ResetQueryColumnFields is a function to reset QueryColumnFields.
func (u *User) ResetQueryColumnFields() {
	u.QueryColumnFields = []string{}
}

