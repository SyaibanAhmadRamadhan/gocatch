package example

import (
	"errors"
	"fmt"
	"strings"

	"github.com/SyaibanAhmadRamadhan/jolly/Jsql"
	"github.com/SyaibanAhmadRamadhan/jolly/Jtype/JOmap"
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

// TableName is a function to get table name
func (u *User) TableName() (table string) {
	return "user"
}

// SchemaName is a function to get schema name
func (u *User) SchemaName() (schema string) {
	return "public"
}

// FieldUpdatedAt is a field or column in the table User.
func (u *User) FieldUpdatedAt() string {
	return "updated_at"
}

// SetUpdatedAt is a setter for the field or column UpdatedAt in the table User.
func (u *User) SetUpdatedAt(param int64) {
	u.UpdatedAt = Jsql.NewNullInt64(&param)
	if u.WCField == nil {
		u.WCField = make(JOmap.SA)
	}
	u.WCField[u.FieldUpdatedAt()] = param
}

// FNamedArgsUpdatedAt sets the value, comparison operator, and logical operator for an argument field.
func (u *User) FNamedArgsUpdatedAt(value any, comparasion Jsql.ComparisonOperator, logical Jsql.LogicalOperator) Jsql.FilterNamedArg {
	namedArg := u.FieldUpdatedAt() + "_namedarg"

	return Jsql.FilterNamedArg{
		Column:      u.FieldUpdatedAt(),
		Value:       value,
		NamedArg:    namedArg,
		Comparasion: comparasion,
		Logical:     logical,
	}

}

// FieldUpdatedBy is a field or column in the table User.
func (u *User) FieldUpdatedBy() string {
	return "updated_by"
}

// SetUpdatedBy is a setter for the field or column UpdatedBy in the table User.
func (u *User) SetUpdatedBy(param string) {
	u.UpdatedBy = Jsql.NewNullString(&param)
	if u.WCField == nil {
		u.WCField = make(JOmap.SA)
	}
	u.WCField[u.FieldUpdatedBy()] = param
}

// FNamedArgsUpdatedBy sets the value, comparison operator, and logical operator for an argument field.
func (u *User) FNamedArgsUpdatedBy(value any, comparasion Jsql.ComparisonOperator, logical Jsql.LogicalOperator) Jsql.FilterNamedArg {
	namedArg := u.FieldUpdatedBy() + "_namedarg"

	return Jsql.FilterNamedArg{
		Column:      u.FieldUpdatedBy(),
		Value:       value,
		NamedArg:    namedArg,
		Comparasion: comparasion,
		Logical:     logical,
	}

}

// FieldRoleID is a field or column in the table User.
func (u *User) FieldRoleID() string {
	return "role_id"
}

// SetRoleID is a setter for the field or column RoleID in the table User.
func (u *User) SetRoleID(param int) {
	u.RoleID = param
	if u.WCField == nil {
		u.WCField = make(JOmap.SA)
	}
	u.WCField[u.FieldRoleID()] = param
}

// FNamedArgsRoleID sets the value, comparison operator, and logical operator for an argument field.
func (u *User) FNamedArgsRoleID(value any, comparasion Jsql.ComparisonOperator, logical Jsql.LogicalOperator) Jsql.FilterNamedArg {
	namedArg := u.FieldRoleID() + "_namedarg"

	return Jsql.FilterNamedArg{
		Column:      u.FieldRoleID(),
		Value:       value,
		NamedArg:    namedArg,
		Comparasion: comparasion,
		Logical:     logical,
	}

}

// FieldUsername is a field or column in the table User.
func (u *User) FieldUsername() string {
	return "username"
}

// SetUsername is a setter for the field or column Username in the table User.
func (u *User) SetUsername(param string) {
	u.Username = param
	if u.WCField == nil {
		u.WCField = make(JOmap.SA)
	}
	u.WCField[u.FieldUsername()] = param
}

// FNamedArgsUsername sets the value, comparison operator, and logical operator for an argument field.
func (u *User) FNamedArgsUsername(value any, comparasion Jsql.ComparisonOperator, logical Jsql.LogicalOperator) Jsql.FilterNamedArg {
	namedArg := u.FieldUsername() + "_namedarg"

	return Jsql.FilterNamedArg{
		Column:      u.FieldUsername(),
		Value:       value,
		NamedArg:    namedArg,
		Comparasion: comparasion,
		Logical:     logical,
	}

}

// FieldCreatedAt is a field or column in the table User.
func (u *User) FieldCreatedAt() string {
	return "created_at"
}

// SetCreatedAt is a setter for the field or column CreatedAt in the table User.
func (u *User) SetCreatedAt(param string) {
	u.CreatedAt = Jsql.NewNullString(&param)
	if u.WCField == nil {
		u.WCField = make(JOmap.SA)
	}
	u.WCField[u.FieldCreatedAt()] = param
}

// FNamedArgsCreatedAt sets the value, comparison operator, and logical operator for an argument field.
func (u *User) FNamedArgsCreatedAt(value any, comparasion Jsql.ComparisonOperator, logical Jsql.LogicalOperator) Jsql.FilterNamedArg {
	namedArg := u.FieldCreatedAt() + "_namedarg"

	return Jsql.FilterNamedArg{
		Column:      u.FieldCreatedAt(),
		Value:       value,
		NamedArg:    namedArg,
		Comparasion: comparasion,
		Logical:     logical,
	}

}

// FieldPhoneNumber is a field or column in the table User.
func (u *User) FieldPhoneNumber() string {
	return "phone_number"
}

// SetPhoneNumber is a setter for the field or column PhoneNumber in the table User.
func (u *User) SetPhoneNumber(param string) {
	u.PhoneNumber = param
	if u.WCField == nil {
		u.WCField = make(JOmap.SA)
	}
	u.WCField[u.FieldPhoneNumber()] = param
}

// FNamedArgsPhoneNumber sets the value, comparison operator, and logical operator for an argument field.
func (u *User) FNamedArgsPhoneNumber(value any, comparasion Jsql.ComparisonOperator, logical Jsql.LogicalOperator) Jsql.FilterNamedArg {
	namedArg := u.FieldPhoneNumber() + "_namedarg"

	return Jsql.FilterNamedArg{
		Column:      u.FieldPhoneNumber(),
		Value:       value,
		NamedArg:    namedArg,
		Comparasion: comparasion,
		Logical:     logical,
	}

}

// FieldCreatedBy is a field or column in the table User.
func (u *User) FieldCreatedBy() string {
	return "created_by"
}

// SetCreatedBy is a setter for the field or column CreatedBy in the table User.
func (u *User) SetCreatedBy(param string) {
	u.CreatedBy = param
	if u.WCField == nil {
		u.WCField = make(JOmap.SA)
	}
	u.WCField[u.FieldCreatedBy()] = param
}

// FNamedArgsCreatedBy sets the value, comparison operator, and logical operator for an argument field.
func (u *User) FNamedArgsCreatedBy(value any, comparasion Jsql.ComparisonOperator, logical Jsql.LogicalOperator) Jsql.FilterNamedArg {
	namedArg := u.FieldCreatedBy() + "_namedarg"

	return Jsql.FilterNamedArg{
		Column:      u.FieldCreatedBy(),
		Value:       value,
		NamedArg:    namedArg,
		Comparasion: comparasion,
		Logical:     logical,
	}

}

// FieldDeletedAt is a field or column in the table User.
func (u *User) FieldDeletedAt() string {
	return "deleted_at"
}

// SetDeletedAt is a setter for the field or column DeletedAt in the table User.
func (u *User) SetDeletedAt(param int64) {
	u.DeletedAt = Jsql.NewNullInt64(&param)
	if u.WCField == nil {
		u.WCField = make(JOmap.SA)
	}
	u.WCField[u.FieldDeletedAt()] = param
}

// FNamedArgsDeletedAt sets the value, comparison operator, and logical operator for an argument field.
func (u *User) FNamedArgsDeletedAt(value any, comparasion Jsql.ComparisonOperator, logical Jsql.LogicalOperator) Jsql.FilterNamedArg {
	namedArg := u.FieldDeletedAt() + "_namedarg"

	return Jsql.FilterNamedArg{
		Column:      u.FieldDeletedAt(),
		Value:       value,
		NamedArg:    namedArg,
		Comparasion: comparasion,
		Logical:     logical,
	}

}

// FieldDeletedBy is a field or column in the table User.
func (u *User) FieldDeletedBy() string {
	return "deleted_by"
}

// SetDeletedBy is a setter for the field or column DeletedBy in the table User.
func (u *User) SetDeletedBy(param string) {
	u.DeletedBy = Jsql.NewNullString(&param)
	if u.WCField == nil {
		u.WCField = make(JOmap.SA)
	}
	u.WCField[u.FieldDeletedBy()] = param
}

// FNamedArgsDeletedBy sets the value, comparison operator, and logical operator for an argument field.
func (u *User) FNamedArgsDeletedBy(value any, comparasion Jsql.ComparisonOperator, logical Jsql.LogicalOperator) Jsql.FilterNamedArg {
	namedArg := u.FieldDeletedBy() + "_namedarg"

	return Jsql.FilterNamedArg{
		Column:      u.FieldDeletedBy(),
		Value:       value,
		NamedArg:    namedArg,
		Comparasion: comparasion,
		Logical:     logical,
	}

}

// FieldID is a field or column in the table User.
func (u *User) FieldID() string {
	return "id"
}

// SetID is a setter for the field or column ID in the table User.
func (u *User) SetID(param string) {
	u.ID = param
	if u.WCField == nil {
		u.WCField = make(JOmap.SA)
	}
	u.WCField[u.FieldID()] = param
}

// FNamedArgsID sets the value, comparison operator, and logical operator for an argument field.
func (u *User) FNamedArgsID(value any, comparasion Jsql.ComparisonOperator, logical Jsql.LogicalOperator) Jsql.FilterNamedArg {
	namedArg := u.FieldID() + "_namedarg"

	return Jsql.FilterNamedArg{
		Column:      u.FieldID(),
		Value:       value,
		NamedArg:    namedArg,
		Comparasion: comparasion,
		Logical:     logical,
	}

}

// FieldEmail is a field or column in the table User.
func (u *User) FieldEmail() string {
	return "email"
}

// SetEmail is a setter for the field or column Email in the table User.
func (u *User) SetEmail(param string) {
	u.Email = param
	if u.WCField == nil {
		u.WCField = make(JOmap.SA)
	}
	u.WCField[u.FieldEmail()] = param
}

// FNamedArgsEmail sets the value, comparison operator, and logical operator for an argument field.
func (u *User) FNamedArgsEmail(value any, comparasion Jsql.ComparisonOperator, logical Jsql.LogicalOperator) Jsql.FilterNamedArg {
	namedArg := u.FieldEmail() + "_namedarg"

	return Jsql.FilterNamedArg{
		Column:      u.FieldEmail(),
		Value:       value,
		NamedArg:    namedArg,
		Comparasion: comparasion,
		Logical:     logical,
	}

}

// FieldPassword is a field or column in the table User.
func (u *User) FieldPassword() string {
	return "password"
}

// SetPassword is a setter for the field or column Password in the table User.
func (u *User) SetPassword(param string) {
	u.Password = param
	if u.WCField == nil {
		u.WCField = make(JOmap.SA)
	}
	u.WCField[u.FieldPassword()] = param
}

// FNamedArgsPassword sets the value, comparison operator, and logical operator for an argument field.
func (u *User) FNamedArgsPassword(value any, comparasion Jsql.ComparisonOperator, logical Jsql.LogicalOperator) Jsql.FilterNamedArg {
	namedArg := u.FieldPassword() + "_namedarg"

	return Jsql.FilterNamedArg{
		Column:      u.FieldPassword(),
		Value:       value,
		NamedArg:    namedArg,
		Comparasion: comparasion,
		Logical:     logical,
	}

}

// AllField is a function to get all field or column in the table User.
func (u *User) AllField() (str []string) {
	str = []string{ 
		`username`,
		`created_at`,
		`updated_at`,
		`updated_by`,
		`role_id`,
		`email`,
		`password`,
		`phone_number`,
		`created_by`,
		`deleted_at`,
		`deleted_by`,
		`id`,
	}
	return
}

// Scan is a function to scan the value with for rows.Value() from the database to the struct User.
func (u *User) Scan(data map[string]any) (err error) {
	for key, value := range data {
		switch key {
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
		case u.FieldCreatedAt():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field CreatedAt")
			}
			u.SetCreatedAt(val)
		case u.FieldUpdatedAt():
			val, ok := value.(int64)
			if !ok {
				return errors.New("invalid type int64. field UpdatedAt")
			}
			u.SetUpdatedAt(val)
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
		case u.FieldCreatedBy():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field CreatedBy")
			}
			u.SetCreatedBy(val)
		default:
			return errors.New("invalid column")
		}
	}
	return nil
}

// RQFieldSet is a function to set the column to RQField for that will be used in the query.
func (u *User) RQFieldSet(columns ...string) (err error) {
	for _, column := range columns {
		switch column {
		case u.FieldRoleID():
		case u.FieldUsername():
		case u.FieldCreatedAt():
		case u.FieldUpdatedAt():
		case u.FieldUpdatedBy():
		case u.FieldDeletedAt():
		case u.FieldDeletedBy():
		case u.FieldID():
		case u.FieldEmail():
		case u.FieldPassword():
		case u.FieldPhoneNumber():
		case u.FieldCreatedBy():
		default:
			return errors.New("invalid column")
		}
		cond := false
		for _, field := range u.RQField {
			if column == field {
				cond = true
				break
			}
		}
		if cond == true {
			continue
		}
		u.RQField = append(u.RQField, column)
	}
	return nil
}

// RQFieldDelete is a function to delete the column from RQField for that will be used in the query.
func (u *User) RQFieldDelete(elems ...string) (err error) {
	var colums []string
	for _, v := range u.RQField{
		colums = append(colums, v)
	}

	for _, elem := range elems {
		index := -1
		for i, column := range u.RQField{
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
	u.RQField = colums
	return nil
}

// RQFieldToString is a function to get the column format string from RQField for that will be used in the query.
func (u *User) RQFieldToString() (columnStr string) {
	return strings.Join(u.RQField, ", ")
}

// RQFieldReset is a function to reset RQField.
func (u *User) RQFieldReset() {
	u.RQField = []string{}
}

// FNamedArgsSetReset is a function to reset FNamedArgs.
func (u *User) FNamedArgsSetReset() {
	u.FNamedArgs = nil
}

// FNamedArgsSet is a function to set locking method.
func (u *User) FNamedArgsSet(param ...Jsql.FilterNamedArg) {
	u.FNamedArgs = append(u.FNamedArgs, param...)
}

// Locking is a function to set locking method.
func (u *User) Locking(lockingOperator Jsql.LockingOperator) Jsql.LockingOperator {
	return lockingOperator
}

