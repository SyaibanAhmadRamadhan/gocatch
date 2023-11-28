package example

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

// TableName is a function to get table name
func (u *User) TableName() (table string) {
	return "user"
}

// SchemaName is a function to get schema name
func (u *User) SchemaName() (schema string) {
	return "public"
}

// FieldRoleID is a field or column in the table User.
func (u *User) FieldRoleID() string {
	return "role_id"
}

// SetRoleID is a setter for the field or column RoleID in the table User.
func (u *User) SetRoleID(param int) {
	u.RoleID = param
}

// SetArgFieldRoleID sets the value, comparison operator, and logical operator for an argument field.
// The `namedArg` variadic parameter can take up to 1 argument:
//   - index 0 (named arg): Named Argument
//
// by default value is namedArg is: FieldRoleID()
// if variadic function > 1, it will be ignored and get index 0.
func (u *User) SetArgFieldRoleID(value any, comparasion Jsql.ComparisonOperator, logical Jsql.LogicalOperator, customNamedArg ...string) {
	namedArg := u.FieldRoleID()

	if customNamedArg[0] != "" {
		namedArg = customNamedArg[0]
	}

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
}

// SetArgFieldEmail sets the value, comparison operator, and logical operator for an argument field.
// The `namedArg` variadic parameter can take up to 1 argument:
//   - index 0 (named arg): Named Argument
//
// by default value is namedArg is: FieldEmail()
// if variadic function > 1, it will be ignored and get index 0.
func (u *User) SetArgFieldEmail(value any, comparasion Jsql.ComparisonOperator, logical Jsql.LogicalOperator, customNamedArg ...string) {
	namedArg := u.FieldEmail()

	if customNamedArg[0] != "" {
		namedArg = customNamedArg[0]
	}

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
}

// SetArgFieldPassword sets the value, comparison operator, and logical operator for an argument field.
// The `namedArg` variadic parameter can take up to 1 argument:
//   - index 0 (named arg): Named Argument
//
// by default value is namedArg is: FieldPassword()
// if variadic function > 1, it will be ignored and get index 0.
func (u *User) SetArgFieldPassword(value any, comparasion Jsql.ComparisonOperator, logical Jsql.LogicalOperator, customNamedArg ...string) {
	namedArg := u.FieldPassword()

	if customNamedArg[0] != "" {
		namedArg = customNamedArg[0]
	}

	u.QFilterNamedArgs = append(u.QFilterNamedArgs, Jsql.FilterNamedArg{
		Column:      u.FieldPassword(),
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
}

// SetArgFieldCreatedAt sets the value, comparison operator, and logical operator for an argument field.
// The `namedArg` variadic parameter can take up to 1 argument:
//   - index 0 (named arg): Named Argument
//
// by default value is namedArg is: FieldCreatedAt()
// if variadic function > 1, it will be ignored and get index 0.
func (u *User) SetArgFieldCreatedAt(value any, comparasion Jsql.ComparisonOperator, logical Jsql.LogicalOperator, customNamedArg ...string) {
	namedArg := u.FieldCreatedAt()

	if customNamedArg[0] != "" {
		namedArg = customNamedArg[0]
	}

	u.QFilterNamedArgs = append(u.QFilterNamedArgs, Jsql.FilterNamedArg{
		Column:      u.FieldCreatedAt(),
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
}

// SetArgFieldDeletedAt sets the value, comparison operator, and logical operator for an argument field.
// The `namedArg` variadic parameter can take up to 1 argument:
//   - index 0 (named arg): Named Argument
//
// by default value is namedArg is: FieldDeletedAt()
// if variadic function > 1, it will be ignored and get index 0.
func (u *User) SetArgFieldDeletedAt(value any, comparasion Jsql.ComparisonOperator, logical Jsql.LogicalOperator, customNamedArg ...string) {
	namedArg := u.FieldDeletedAt()

	if customNamedArg[0] != "" {
		namedArg = customNamedArg[0]
	}

	u.QFilterNamedArgs = append(u.QFilterNamedArgs, Jsql.FilterNamedArg{
		Column:      u.FieldDeletedAt(),
		Value:       value,
		NamedArg:    namedArg,
		Comparasion: comparasion,
		Logical:     logical,
	})

}

// FieldID is a field or column in the table User.
func (u *User) FieldID() string {
	return "id"
}

// SetID is a setter for the field or column ID in the table User.
func (u *User) SetID(param string) {
	u.ID = param
}

// SetArgFieldID sets the value, comparison operator, and logical operator for an argument field.
// The `namedArg` variadic parameter can take up to 1 argument:
//   - index 0 (named arg): Named Argument
//
// by default value is namedArg is: FieldID()
// if variadic function > 1, it will be ignored and get index 0.
func (u *User) SetArgFieldID(value any, comparasion Jsql.ComparisonOperator, logical Jsql.LogicalOperator, customNamedArg ...string) {
	namedArg := u.FieldID()

	if customNamedArg[0] != "" {
		namedArg = customNamedArg[0]
	}

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
}

// SetArgFieldUsername sets the value, comparison operator, and logical operator for an argument field.
// The `namedArg` variadic parameter can take up to 1 argument:
//   - index 0 (named arg): Named Argument
//
// by default value is namedArg is: FieldUsername()
// if variadic function > 1, it will be ignored and get index 0.
func (u *User) SetArgFieldUsername(value any, comparasion Jsql.ComparisonOperator, logical Jsql.LogicalOperator, customNamedArg ...string) {
	namedArg := u.FieldUsername()

	if customNamedArg[0] != "" {
		namedArg = customNamedArg[0]
	}

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
}

// SetArgFieldPhoneNumber sets the value, comparison operator, and logical operator for an argument field.
// The `namedArg` variadic parameter can take up to 1 argument:
//   - index 0 (named arg): Named Argument
//
// by default value is namedArg is: FieldPhoneNumber()
// if variadic function > 1, it will be ignored and get index 0.
func (u *User) SetArgFieldPhoneNumber(value any, comparasion Jsql.ComparisonOperator, logical Jsql.LogicalOperator, customNamedArg ...string) {
	namedArg := u.FieldPhoneNumber()

	if customNamedArg[0] != "" {
		namedArg = customNamedArg[0]
	}

	u.QFilterNamedArgs = append(u.QFilterNamedArgs, Jsql.FilterNamedArg{
		Column:      u.FieldPhoneNumber(),
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
}

// SetArgFieldUpdatedAt sets the value, comparison operator, and logical operator for an argument field.
// The `namedArg` variadic parameter can take up to 1 argument:
//   - index 0 (named arg): Named Argument
//
// by default value is namedArg is: FieldUpdatedAt()
// if variadic function > 1, it will be ignored and get index 0.
func (u *User) SetArgFieldUpdatedAt(value any, comparasion Jsql.ComparisonOperator, logical Jsql.LogicalOperator, customNamedArg ...string) {
	namedArg := u.FieldUpdatedAt()

	if customNamedArg[0] != "" {
		namedArg = customNamedArg[0]
	}

	u.QFilterNamedArgs = append(u.QFilterNamedArgs, Jsql.FilterNamedArg{
		Column:      u.FieldUpdatedAt(),
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
}

// SetArgFieldCreatedBy sets the value, comparison operator, and logical operator for an argument field.
// The `namedArg` variadic parameter can take up to 1 argument:
//   - index 0 (named arg): Named Argument
//
// by default value is namedArg is: FieldCreatedBy()
// if variadic function > 1, it will be ignored and get index 0.
func (u *User) SetArgFieldCreatedBy(value any, comparasion Jsql.ComparisonOperator, logical Jsql.LogicalOperator, customNamedArg ...string) {
	namedArg := u.FieldCreatedBy()

	if customNamedArg[0] != "" {
		namedArg = customNamedArg[0]
	}

	u.QFilterNamedArgs = append(u.QFilterNamedArgs, Jsql.FilterNamedArg{
		Column:      u.FieldCreatedBy(),
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
}

// SetArgFieldUpdatedBy sets the value, comparison operator, and logical operator for an argument field.
// The `namedArg` variadic parameter can take up to 1 argument:
//   - index 0 (named arg): Named Argument
//
// by default value is namedArg is: FieldUpdatedBy()
// if variadic function > 1, it will be ignored and get index 0.
func (u *User) SetArgFieldUpdatedBy(value any, comparasion Jsql.ComparisonOperator, logical Jsql.LogicalOperator, customNamedArg ...string) {
	namedArg := u.FieldUpdatedBy()

	if customNamedArg[0] != "" {
		namedArg = customNamedArg[0]
	}

	u.QFilterNamedArgs = append(u.QFilterNamedArgs, Jsql.FilterNamedArg{
		Column:      u.FieldUpdatedBy(),
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
}

// SetArgFieldDeletedBy sets the value, comparison operator, and logical operator for an argument field.
// The `namedArg` variadic parameter can take up to 1 argument:
//   - index 0 (named arg): Named Argument
//
// by default value is namedArg is: FieldDeletedBy()
// if variadic function > 1, it will be ignored and get index 0.
func (u *User) SetArgFieldDeletedBy(value any, comparasion Jsql.ComparisonOperator, logical Jsql.LogicalOperator, customNamedArg ...string) {
	namedArg := u.FieldDeletedBy()

	if customNamedArg[0] != "" {
		namedArg = customNamedArg[0]
	}

	u.QFilterNamedArgs = append(u.QFilterNamedArgs, Jsql.FilterNamedArg{
		Column:      u.FieldDeletedBy(),
		Value:       value,
		NamedArg:    namedArg,
		Comparasion: comparasion,
		Logical:     logical,
	})

}

// AllField is a function to get all field or column in the table User.
func (u *User) AllField() (str string) {
	str += `
		updated_at, 
		created_by, 
		updated_by, 
		deleted_by, 
		id, 
		username, 
		phone_number, 
		created_at, 
		deleted_at, 
		role_id, 
		email, 
		password`
	return
}

// Scan is a function to scan the value with for rows.Value() from the database to the struct User.
func (u *User) Scan(data map[string]any) (err error) {
	for key, value := range data {
		switch key {
		case u.FieldCreatedAt():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field CreatedAt")
			}
			u.SetCreatedAt(val)
			return nil
		case u.FieldDeletedAt():
			val, ok := value.(int64)
			if !ok {
				return errors.New("invalid type int64. field DeletedAt")
			}
			u.SetDeletedAt(val)
			return nil
		case u.FieldRoleID():
			val, ok := value.(int)
			if !ok {
				return errors.New("invalid type int. field RoleID")
			}
			u.SetRoleID(val)
			return nil
		case u.FieldEmail():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field Email")
			}
			u.SetEmail(val)
			return nil
		case u.FieldPassword():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field Password")
			}
			u.SetPassword(val)
			return nil
		case u.FieldUpdatedAt():
			val, ok := value.(int64)
			if !ok {
				return errors.New("invalid type int64. field UpdatedAt")
			}
			u.SetUpdatedAt(val)
			return nil
		case u.FieldCreatedBy():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field CreatedBy")
			}
			u.SetCreatedBy(val)
			return nil
		case u.FieldUpdatedBy():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field UpdatedBy")
			}
			u.SetUpdatedBy(val)
			return nil
		case u.FieldDeletedBy():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field DeletedBy")
			}
			u.SetDeletedBy(val)
			return nil
		case u.FieldID():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field ID")
			}
			u.SetID(val)
			return nil
		case u.FieldUsername():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field Username")
			}
			u.SetUsername(val)
			return nil
		case u.FieldPhoneNumber():
			val, ok := value.(string)
			if !ok {
				return errors.New("invalid type string. field PhoneNumber")
			}
			u.SetPhoneNumber(val)
			return nil
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
		case u.FieldEmail():
		case u.FieldPassword():
		case u.FieldCreatedAt():
		case u.FieldDeletedAt():
		case u.FieldRoleID():
		case u.FieldUsername():
		case u.FieldPhoneNumber():
		case u.FieldUpdatedAt():
		case u.FieldCreatedBy():
		case u.FieldUpdatedBy():
		case u.FieldDeletedBy():
		case u.FieldID():
		default:
			return errors.New("invalid column")
		}
	}
	u.QColumnFields = append(u.QColumnFields, columns...)
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

// QueryColumnFieldToStrings is a function to get the column format string from QColumnFields for that will be used in the query.
func (u *User) QueryColumnFieldToStrings() (columnStr string) {
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
