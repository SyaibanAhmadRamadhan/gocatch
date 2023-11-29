package example

import (
	"errors"
	"fmt"
	"strings"

	"github.com/SyaibanAhmadRamadhan/jolly/Jtype/JOmap"

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

// TableName is a function to get table name
func (u *User) TableName() (table string) {
	return "user"
}

// SchemaName is a function to get schema name
func (u *User) SchemaName() (schema string) {
	return "public"
}

// FieldID is a field or column in the table User.
func (u *User) FieldID() string {
	return "id"
}

// SetID is a setter for the field or column ID in the table User.
func (u *User) SetID(param string) {
	u.ID = param
	cond := false
	for _, field := range u.QColumnFields {
		if u.FieldID() == field {
			cond = true
			break
		}
	}
	if !cond {
		u.QColumnFields = append(u.QColumnFields, u.FieldID())
	}
}

// SetArgFieldID sets the value, comparison operator, and logical operator for an argument field.
func (u *User) SetArgFieldID(value any, comparasion Jsql.ComparisonOperator, logical Jsql.LogicalOperator) {
	namedArg := u.FieldID() + "_where"

	u.QFilterNamedArgs = append(u.QFilterNamedArgs, Jsql.FilterNamedArg{
		Column:      u.FieldID(),
		Value:       value,
		NamedArg:    namedArg,
		Comparasion: comparasion,
		Logical:     logical,
	})

}

// FieldUsername is a field or column in the table User.
func (u *User) FieldUsername() string {
	return "username"
}

// SetUsername is a setter for the field or column Username in the table User.
func (u *User) SetUsername(param string) {
	u.Username = param
	cond := false
	for _, field := range u.QColumnFields {
		if u.FieldUsername() == field {
			cond = true
			break
		}
	}
	if !cond {
		u.QColumnFields = append(u.QColumnFields, u.FieldUsername())
	}
}

// SetArgFieldUsername sets the value, comparison operator, and logical operator for an argument field.
func (u *User) SetArgFieldUsername(value any, comparasion Jsql.ComparisonOperator, logical Jsql.LogicalOperator) {
	namedArg := u.FieldUsername() + "_where"

	u.QFilterNamedArgs = append(u.QFilterNamedArgs, Jsql.FilterNamedArg{
		Column:      u.FieldUsername(),
		Value:       value,
		NamedArg:    namedArg,
		Comparasion: comparasion,
		Logical:     logical,
	})

}

// FieldPhoneNumber is a field or column in the table User.
func (u *User) FieldPhoneNumber() string {
	return "phone_number"
}

// SetPhoneNumber is a setter for the field or column PhoneNumber in the table User.
func (u *User) SetPhoneNumber(param string) {
	u.PhoneNumber = param
	cond := false
	for _, field := range u.QColumnFields {
		if u.FieldPhoneNumber() == field {
			cond = true
			break
		}
	}
	if !cond {
		u.QColumnFields = append(u.QColumnFields, u.FieldPhoneNumber())
	}
}

// SetArgFieldPhoneNumber sets the value, comparison operator, and logical operator for an argument field.
func (u *User) SetArgFieldPhoneNumber(value any, comparasion Jsql.ComparisonOperator, logical Jsql.LogicalOperator) {
	namedArg := u.FieldPhoneNumber() + "_where"

	u.QFilterNamedArgs = append(u.QFilterNamedArgs, Jsql.FilterNamedArg{
		Column:      u.FieldPhoneNumber(),
		Value:       value,
		NamedArg:    namedArg,
		Comparasion: comparasion,
		Logical:     logical,
	})

}

// FieldUpdatedBy is a field or column in the table User.
func (u *User) FieldUpdatedBy() string {
	return "updated_by"
}

// SetUpdatedBy is a setter for the field or column UpdatedBy in the table User.
func (u *User) SetUpdatedBy(param string) {
	u.UpdatedBy = Jsql.NewNullString(&param)
	cond := false
	for _, field := range u.QColumnFields {
		if u.FieldUpdatedBy() == field {
			cond = true
			break
		}
	}
	if !cond {
		u.QColumnFields = append(u.QColumnFields, u.FieldUpdatedBy())
	}
}

// SetArgFieldUpdatedBy sets the value, comparison operator, and logical operator for an argument field.
func (u *User) SetArgFieldUpdatedBy(value any, comparasion Jsql.ComparisonOperator, logical Jsql.LogicalOperator) {
	namedArg := u.FieldUpdatedBy() + "_where"

	u.QFilterNamedArgs = append(u.QFilterNamedArgs, Jsql.FilterNamedArg{
		Column:      u.FieldUpdatedBy(),
		Value:       value,
		NamedArg:    namedArg,
		Comparasion: comparasion,
		Logical:     logical,
	})

}

// FieldCreatedBy is a field or column in the table User.
func (u *User) FieldCreatedBy() string {
	return "created_by"
}

// SetCreatedBy is a setter for the field or column CreatedBy in the table User.
func (u *User) SetCreatedBy(param string) {
	u.CreatedBy = param
	cond := false
	for _, field := range u.QColumnFields {
		if u.FieldCreatedBy() == field {
			cond = true
			break
		}
	}
	if !cond {
		u.QColumnFields = append(u.QColumnFields, u.FieldCreatedBy())
	}
}

// SetArgFieldCreatedBy sets the value, comparison operator, and logical operator for an argument field.
func (u *User) SetArgFieldCreatedBy(value any, comparasion Jsql.ComparisonOperator, logical Jsql.LogicalOperator) {
	namedArg := u.FieldCreatedBy() + "_where"

	u.QFilterNamedArgs = append(u.QFilterNamedArgs, Jsql.FilterNamedArg{
		Column:      u.FieldCreatedBy(),
		Value:       value,
		NamedArg:    namedArg,
		Comparasion: comparasion,
		Logical:     logical,
	})

}

// FieldRoleID is a field or column in the table User.
func (u *User) FieldRoleID() string {
	return "role_id"
}

// SetRoleID is a setter for the field or column RoleID in the table User.
func (u *User) SetRoleID(param int) {
	u.RoleID = param
	cond := false
	for _, field := range u.QColumnFields {
		if u.FieldRoleID() == field {
			cond = true
			break
		}
	}
	if !cond {
		u.QColumnFields = append(u.QColumnFields, u.FieldRoleID())
	}
}

// SetArgFieldRoleID sets the value, comparison operator, and logical operator for an argument field.
func (u *User) SetArgFieldRoleID(value any, comparasion Jsql.ComparisonOperator, logical Jsql.LogicalOperator) {
	namedArg := u.FieldRoleID() + "_where"

	u.QFilterNamedArgs = append(u.QFilterNamedArgs, Jsql.FilterNamedArg{
		Column:      u.FieldRoleID(),
		Value:       value,
		NamedArg:    namedArg,
		Comparasion: comparasion,
		Logical:     logical,
	})

}

// FieldEmail is a field or column in the table User.
func (u *User) FieldEmail() string {
	return "email"
}

// SetEmail is a setter for the field or column Email in the table User.
func (u *User) SetEmail(param string) {
	u.Email = param
	cond := false
	for _, field := range u.QColumnFields {
		if u.FieldEmail() == field {
			cond = true
			break
		}
	}
	if !cond {
		u.QColumnFields = append(u.QColumnFields, u.FieldEmail())
	}
}

// SetArgFieldEmail sets the value, comparison operator, and logical operator for an argument field.
func (u *User) SetArgFieldEmail(value any, comparasion Jsql.ComparisonOperator, logical Jsql.LogicalOperator) {
	namedArg := u.FieldEmail() + "_where"

	u.QFilterNamedArgs = append(u.QFilterNamedArgs, Jsql.FilterNamedArg{
		Column:      u.FieldEmail(),
		Value:       value,
		NamedArg:    namedArg,
		Comparasion: comparasion,
		Logical:     logical,
	})

}

// FieldPassword is a field or column in the table User.
func (u *User) FieldPassword() string {
	return "password"
}

// SetPassword is a setter for the field or column Password in the table User.
func (u *User) SetPassword(param string) {
	u.Password = param
	cond := false
	for _, field := range u.QColumnFields {
		if u.FieldPassword() == field {
			cond = true
			break
		}
	}
	if !cond {
		u.QColumnFields = append(u.QColumnFields, u.FieldPassword())
	}
}

// SetArgFieldPassword sets the value, comparison operator, and logical operator for an argument field.
func (u *User) SetArgFieldPassword(value any, comparasion Jsql.ComparisonOperator, logical Jsql.LogicalOperator) {
	namedArg := u.FieldPassword() + "_where"

	u.QFilterNamedArgs = append(u.QFilterNamedArgs, Jsql.FilterNamedArg{
		Column:      u.FieldPassword(),
		Value:       value,
		NamedArg:    namedArg,
		Comparasion: comparasion,
		Logical:     logical,
	})

}

// FieldDeletedAt is a field or column in the table User.
func (u *User) FieldDeletedAt() string {
	return "deleted_at"
}

// SetDeletedAt is a setter for the field or column DeletedAt in the table User.
func (u *User) SetDeletedAt(param int64) {
	u.DeletedAt = Jsql.NewNullInt64(&param)
	cond := false
	for _, field := range u.QColumnFields {
		if u.FieldDeletedAt() == field {
			cond = true
			break
		}
	}
	if !cond {
		u.QColumnFields = append(u.QColumnFields, u.FieldDeletedAt())
	}
}

// SetArgFieldDeletedAt sets the value, comparison operator, and logical operator for an argument field.
func (u *User) SetArgFieldDeletedAt(value any, comparasion Jsql.ComparisonOperator, logical Jsql.LogicalOperator) {
	namedArg := u.FieldDeletedAt() + "_where"

	u.QFilterNamedArgs = append(u.QFilterNamedArgs, Jsql.FilterNamedArg{
		Column:      u.FieldDeletedAt(),
		Value:       value,
		NamedArg:    namedArg,
		Comparasion: comparasion,
		Logical:     logical,
	})

}

// FieldDeletedBy is a field or column in the table User.
func (u *User) FieldDeletedBy() string {
	return "deleted_by"
}

// SetDeletedBy is a setter for the field or column DeletedBy in the table User.
func (u *User) SetDeletedBy(param string) {
	u.DeletedBy = Jsql.NewNullString(&param)
	cond := false
	for _, field := range u.QColumnFields {
		if u.FieldDeletedBy() == field {
			cond = true
			break
		}
	}
	if !cond {
		u.QColumnFields = append(u.QColumnFields, u.FieldDeletedBy())
	}
}

// SetArgFieldDeletedBy sets the value, comparison operator, and logical operator for an argument field.
func (u *User) SetArgFieldDeletedBy(value any, comparasion Jsql.ComparisonOperator, logical Jsql.LogicalOperator) {
	namedArg := u.FieldDeletedBy() + "_where"

	u.QFilterNamedArgs = append(u.QFilterNamedArgs, Jsql.FilterNamedArg{
		Column:      u.FieldDeletedBy(),
		Value:       value,
		NamedArg:    namedArg,
		Comparasion: comparasion,
		Logical:     logical,
	})

}

// FieldCreatedAt is a field or column in the table User.
func (u *User) FieldCreatedAt() string {
	return "created_at"
}

// SetCreatedAt is a setter for the field or column CreatedAt in the table User.
func (u *User) SetCreatedAt(param string) {
	u.CreatedAt = Jsql.NewNullString(&param)
	cond := false
	for _, field := range u.QColumnFields {
		if u.FieldCreatedAt() == field {
			cond = true
			break
		}
	}
	if !cond {
		u.QColumnFields = append(u.QColumnFields, u.FieldCreatedAt())
	}
}

// SetArgFieldCreatedAt sets the value, comparison operator, and logical operator for an argument field.
func (u *User) SetArgFieldCreatedAt(value any, comparasion Jsql.ComparisonOperator, logical Jsql.LogicalOperator) {
	namedArg := u.FieldCreatedAt() + "_where"

	u.QFilterNamedArgs = append(u.QFilterNamedArgs, Jsql.FilterNamedArg{
		Column:      u.FieldCreatedAt(),
		Value:       value,
		NamedArg:    namedArg,
		Comparasion: comparasion,
		Logical:     logical,
	})

}

// FieldUpdatedAt is a field or column in the table User.
func (u *User) FieldUpdatedAt() string {
	return "updated_at"
}

// SetUpdatedAt is a setter for the field or column UpdatedAt in the table User.
func (u *User) SetUpdatedAt(param int64) {
	u.UpdatedAt = Jsql.NewNullInt64(&param)
	cond := false
	for _, field := range u.QColumnFields {
		if u.FieldUpdatedAt() == field {
			cond = true
			break
		}
	}
	if !cond {
		u.QColumnFields = append(u.QColumnFields, u.FieldUpdatedAt())
	}
}

// SetArgFieldUpdatedAt sets the value, comparison operator, and logical operator for an argument field.
func (u *User) SetArgFieldUpdatedAt(value any, comparasion Jsql.ComparisonOperator, logical Jsql.LogicalOperator) {
	namedArg := u.FieldUpdatedAt() + "_where"

	u.QFilterNamedArgs = append(u.QFilterNamedArgs, Jsql.FilterNamedArg{
		Column:      u.FieldUpdatedAt(),
		Value:       value,
		NamedArg:    namedArg,
		Comparasion: comparasion,
		Logical:     logical,
	})

}

// AllField is a function to get all field or column in the table User.
func (u *User) AllField() (str string) {
	str += `
		deleted_at, 
		deleted_by, 
		created_at, 
		updated_at, 
		created_by, 
		role_id, 
		email, 
		password, 
		updated_by, 
		id, 
		username, 
		phone_number`
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
		case u.FieldCreatedBy():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field CreatedBy")
			}
			u.SetCreatedBy(val)
		case u.FieldID():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field ID")
			}
			u.SetID(val)
		case u.FieldUsername():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field Username")
			}
			u.SetUsername(val)
		case u.FieldPhoneNumber():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field PhoneNumber")
			}
			u.SetPhoneNumber(val)
		case u.FieldUpdatedBy():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field UpdatedBy")
			}
			u.SetUpdatedBy(val)
		default:
			return errors.New("invalid column")
		}
	}
	return nil
}

// SetColumn is a function to set the column to QColumnFields for that will be used in the query.
func (u *User) SetColumn(columns ...string) (err error) {
	for _, column := range columns {
		switch column {
		case u.FieldPassword():
		case u.FieldDeletedAt():
		case u.FieldDeletedBy():
		case u.FieldCreatedAt():
		case u.FieldUpdatedAt():
		case u.FieldCreatedBy():
		case u.FieldRoleID():
		case u.FieldEmail():
		case u.FieldPhoneNumber():
		case u.FieldUpdatedBy():
		case u.FieldID():
		case u.FieldUsername():
		default:
			return errors.New("invalid column")
		}
		cond := false
		for _, field := range u.QColumnFields {
			if column == field {
				cond = true
				break
			}
		}
		if cond == true {
			continue
		}
		u.QColumnFields = append(u.QColumnFields, column)
	}
	return nil
}

// DeleteColumnFromQColumnFields is a function to delete the column from QColumnFields for that will be used in the query.
func (u *User) DeleteColumnFromQColumnFields(elems ...string) (err error) {
	var colums []string
	for _, v := range u.QColumnFields {
		colums = append(colums, v)
	}

	for _, elem := range elems {
		index := -1
		for i, column := range u.QColumnFields {
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
	u.QColumnFields = colums
	return nil
}

// QColumnFieldToStrings is a function to get the column format string from QColumnFields for that will be used in the query.
func (u *User) QColumnFieldToStrings() (columnStr string) {
	return strings.Join(u.QColumnFields, ", ")
}

// ResetQColumnFields is a function to reset QColumnFields.
func (u *User) ResetQColumnFields() {
	u.QColumnFields = []string{}
}

// ResetQFilterNamedArgs is a function to reset QFilterNamedArgs.
func (u *User) ResetQFilterNamedArgs() {
	u.QFilterNamedArgs = Jsql.QFilterNamedArgs{}
}

// Locking is a function to set locking method.
func (u *User) Locking(lockingOperator Jsql.LockingOperator) Jsql.LockingOperator {
	return lockingOperator
}

// FieldAndValue is  function for get named arg for write query
func (u *User) FieldAndValue() JOmap.SA {
	sa := make(JOmap.SA)
	for _, field := range u.QColumnFields {
		switch field {
		case u.FieldCreatedBy():
			sa[field] = u.CreatedBy
		case u.FieldRoleID():
			sa[field] = u.RoleID
		case u.FieldEmail():
			sa[field] = u.Email
		case u.FieldPassword():
			sa[field] = u.Password
		case u.FieldDeletedAt():
			sa[field] = u.DeletedAt
		case u.FieldDeletedBy():
			sa[field] = u.DeletedBy
		case u.FieldCreatedAt():
			sa[field] = u.CreatedAt
		case u.FieldUpdatedAt():
			sa[field] = u.UpdatedAt
		case u.FieldID():
			sa[field] = u.ID
		case u.FieldUsername():
			sa[field] = u.Username
		case u.FieldPhoneNumber():
			sa[field] = u.PhoneNumber
		case u.FieldUpdatedBy():
			sa[field] = u.UpdatedBy
		}
	}
	return sa
}

// FieldArgForUpdate is function get string to SET update
func (u *User) FieldArgForUpdate(prefix Jsql.PrefixNamedArgPG) string {
	str := ""
	columns := u.FieldAndValue()
	i := 1
	for k, _ := range columns {
		if i == len(columns) {
			str += k + " = " + string(prefix) + k
		} else {
			str += k + " = " + string(prefix) + k + ", "
		}
		i++
	}
	return str
}
