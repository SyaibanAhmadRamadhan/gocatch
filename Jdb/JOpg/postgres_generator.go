package JOpg

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"reflect"
	"regexp"
	"strings"

	"github.com/SyaibanAhmadRamadhan/jolly/Jtype/JOFD"

	"github.com/SyaibanAhmadRamadhan/jolly"
	"github.com/SyaibanAhmadRamadhan/jolly/Jlog"
	"github.com/SyaibanAhmadRamadhan/jolly/Jtype/JOstr"
	"github.com/SyaibanAhmadRamadhan/jolly/Jtype/JOstruct"
)

type SpecifiationTable struct {
	TableName  string
	SchemaName string
}

type WithInterfaceRepo struct {
	Make          bool
	Dir           string
	InterfaceName string
	Lib           string // use pgx or sqlx. if you want generates with sqlx, you must set Lib to "sqlx", default is "pgx"
}

type WithInterfaceImplRepo struct {
	Make              bool
	ImplLocationDir   string
	ImplInterfaceName string
}

// GeneratorModelForStructParam defines the parameters
// for generating methods for a struct
type GeneratorModelForStructParam struct {
	Src                   any                   // Src is the struct instance
	SpecifiationTable     SpecifiationTable     // SpecifiationTable is the struct instance that contains the specification of the table
	WithInterfaceRepo     WithInterfaceRepo     // WithInterfaceRepo is a flag to generate interface repository
	WithInterfaceImplRepo WithInterfaceImplRepo // WithInterfaceRepo is a flag to generate interface repository
	FileName              string                // FileName is the name of generated source file without extension
}

// GeneratorModelForStruct generates methods for given structs.
// This function expects each struct to contain a field `QColumnFields []string`.
// The generated methods manipulate or make use of this field for various operations.
// The function does not support nested structs with name tag.
// if you want nested struct for this function, you can use "-" tag in your struct.
// for example check in folder Jdb/JOpg/exGen
func GeneratorModelForStruct(params ...GeneratorModelForStructParam) {
	for _, param := range params {
		if param.SpecifiationTable.TableName == "" {
			panic("TableName is empty")
		}
		if param.SpecifiationTable.SchemaName == "" {
			panic("SchemaName is empty")
		}

		if param.WithInterfaceRepo.Make {
			if param.WithInterfaceRepo.InterfaceName == "" {
				panic("InterfaceName is empty")
			}

			if param.WithInterfaceRepo.Dir == "" {
				panic("Dir is empty")
			}
		}

		if param.WithInterfaceImplRepo.Make {
			if param.WithInterfaceImplRepo.ImplInterfaceName == "" {
				panic("InterfaceName is empty")
			}

			if param.WithInterfaceImplRepo.ImplLocationDir == "" {
				panic("Dir is empty")
			}
		}

		buf := bytes.Buffer{}
		fn := func(str string) {
			_, err := buf.WriteString(str)
			jolly.PanicIF(err)
		}

		t := reflect.TypeOf(param.Src).Elem()
		caller := Jlog.CallerInfo(2)
		packageName := JOstr.LastStringOfSubStr(caller.PackageName, "/")
		fn(`package ` + packageName + "\n\n")

		// Prepare import statements
		res, _ := os.ReadFile(param.FileName + ".go")
		r, _ := regexp.Compile(`import \(([^)]+)\)`)
		importSrc := r.FindString(string(res))

		customImport := fmt.Sprintf(
			`(%s%s%s%s%s`,
			"\n\t\"errors\"",
			"\n\t\"fmt\"",
			"\n\t\"strings\"",
			"\n",
			"\t\"github.com/SyaibanAhmadRamadhan/jolly/Jtype/JOmap\"\n",
		)

		importSrc = strings.Replace(importSrc, "(", customImport, 1)
		fn(importSrc + "\n\n")

		fn("// DO NOT EDIT, will be overwritten by https://github.com/SyaibanAhmadRamadhan/jolly/blob/main/Jdb/JOpg/postgres_generator.go. \n\n")

		fn("// New" + t.Name() + " is a struct with pointer that represents the table " + t.Name() + " in the database.\n")
		fn(`func New` + t.Name() + `() *` + t.Name() + " {\n")
		fn(`	return &` + t.Name() + "{}\n")
		fn("}\n\n")

		fn("// New" + t.Name() + "WithOutPtr is a struct without pointer that represents the table " + t.Name() + " in the database.\n")
		fn(`func New` + t.Name() + `WithOutPtr() ` + t.Name() + " {\n")
		fn(`	return ` + t.Name() + "{}\n")
		fn("}\n\n")

		fn("// TableName is a function to get table name\n")
		fn(`func (` + JOstr.FirstCharToLower(t.Name()) + ` *` + t.Name() + ") TableName() (table string) {\n")
		fn(`	return "` + param.SpecifiationTable.TableName + "\"\n")
		fn("}\n\n")

		fn("// SchemaName is a function to get schema name\n")
		fn(`func (` + JOstr.FirstCharToLower(t.Name()) + ` *` + t.Name() + ") SchemaName() (schema string) {\n")
		fn(`	return "` + param.SpecifiationTable.SchemaName + "\"\n")
		fn("}\n\n")

		field := JOstruct.GetTagAndFieldNameFromStruct(param.Src, "", "db")
		for k, v := range field {
			typeStruct := strings.Split(v, "|")[1]

			fn("// Field" + k + " is a field or column in the table " + t.Name() + ".\n")
			fn(`func (` + JOstr.FirstCharToLower(t.Name()) + ` *` + t.Name() + ") Field" + k + "() string {\n")
			fn(`	return "` + strings.Split(v, "|")[0] + "\"\n")
			fn("}\n\n")

			fn("// Set" + k + " is a setter for the field or column " + k + " in the table " + t.Name() + ".\n")
			switch typeStruct {
			case "Jsql.NullString":
				fn(`func (` + JOstr.FirstCharToLower(t.Name()) + ` *` + t.Name() + ") Set" + k + "(param string) {\n")
			case "Jsql.NullBool":
				fn(`func (` + JOstr.FirstCharToLower(t.Name()) + ` *` + t.Name() + ") Set" + k + "(param bool) {\n")
			case "Jsql.NullFloat64":
				fn(`func (` + JOstr.FirstCharToLower(t.Name()) + ` *` + t.Name() + ") Set" + k + "(param float64) {\n")
			case "Jsql.NullInt64":
				fn(`func (` + JOstr.FirstCharToLower(t.Name()) + ` *` + t.Name() + ") Set" + k + "(param int64) {\n")
			case "Jsql.NullInt32":
				fn(`func (` + JOstr.FirstCharToLower(t.Name()) + ` *` + t.Name() + ") Set" + k + "(param int32) {\n")
			case "Jsql.NullByte":
				fn(`func (` + JOstr.FirstCharToLower(t.Name()) + ` *` + t.Name() + ") Set" + k + "(param byte) {\n")
			case "Jsql.NullTime":
				fn(`func (` + JOstr.FirstCharToLower(t.Name()) + ` *` + t.Name() + ") Set" + k + "(param time.Time) {\n")
			case "Jsql.NullInt16":
				fn(`func (` + JOstr.FirstCharToLower(t.Name()) + ` *` + t.Name() + ") Set" + k + "(param int16) {\n")
			default:
				fn(`func (` + JOstr.FirstCharToLower(t.Name()) + ` *` + t.Name() + ") Set" + k + "(param " + strings.Split(v, "|")[1] + ") {\n")
			}

			if typeStruct == "Jsql.NullString" ||
				typeStruct == "Jsql.NullBool" ||
				typeStruct == "Jsql.NullFloat64" ||
				typeStruct == "Jsql.NullInt64" ||
				typeStruct == "Jsql.NullInt32" ||
				typeStruct == "Jsql.NullByte" ||
				typeStruct == "Jsql.NullTime" ||
				typeStruct == "Jsql.NullInt16" {
				fn(`	` + JOstr.FirstCharToLower(t.Name()) + "." + k + " = " + strings.ReplaceAll(typeStruct, ".", ".New") + "(&param)" + "\n")
			} else {
				fn(`	` + JOstr.FirstCharToLower(t.Name()) + "." + k + " = param" + "\n")
			}
			fn(`	cond := false` + "\n")
			fn(`	for _, field := range ` + JOstr.FirstCharToLower(t.Name()) + ".QColumnFields " + "{\n")
			fn(`		if ` + JOstr.FirstCharToLower(t.Name()) + ".Field" + k + "()" + ` == field {` + "\n")
			fn(`			cond = true` + "\n")
			fn(`			break` + "\n")
			fn(`		}` + "\n")
			fn(`	}` + "\n")
			fn(`	if !cond {` + "\n")
			fn(`		` + JOstr.FirstCharToLower(t.Name()) + ".QColumnFields = append(" +
				JOstr.FirstCharToLower(t.Name()) + ".QColumnFields, " + JOstr.FirstCharToLower(t.Name()) + ".Field" + k + "())\n")
			fn(`	}` + "\n")
			fn("}\n\n")

			fn("// SetArgField" + k + " sets the value, comparison operator, and logical operator for an argument field.\n")
			fn(`func (` + JOstr.FirstCharToLower(t.Name()) + ` *` + t.Name() + ") SetArgField" + k +
				"(value any, comparasion Jsql.ComparisonOperator, logical Jsql.LogicalOperator) {\n")
			fn(`	namedArg := ` + JOstr.FirstCharToLower(t.Name()) + ".Field" + k + "() + \"_where\"\n")
			fn("\n")
			fn(`	` + JOstr.FirstCharToLower(t.Name()) + `.QFilterNamedArgs = append(` + JOstr.FirstCharToLower(t.Name()) + `.QFilterNamedArgs, Jsql.FilterNamedArg{` + "\n")
			fn(`		Column:      ` + JOstr.FirstCharToLower(t.Name()) + ".Field" + k + "()," + "\n")
			fn(`		Value:       value,` + "\n")
			fn(`		NamedArg:    namedArg,` + "\n")
			fn(`		Comparasion: comparasion,` + "\n")
			fn(`		Logical:     logical,` + "\n")
			fn(`	})` + "\n")
			fn("\n")
			fn("}\n\n")
		}

		fn("// AllField is a function to get all field or column in the table " + t.Name() + ".\n")
		fn(`func (` + JOstr.FirstCharToLower(t.Name()) + ` *` + t.Name() + ") AllField() (str string) {\n")
		fn(`	str += `)

		var arrField []string
		for _, v := range field {
			arrField = append(arrField, strings.Split(v, "|")[0])
		}
		fn("`\n\t\t" + strings.Join(arrField, ", \n\t\t") + "`" + "\n")
		fn(`	return` + "\n")
		fn("}\n\n")

		fn("// Scan is a function to scan the value with for rows.Value() from the database to the struct " + t.Name() + ".\n")
		fn(`func (` + JOstr.FirstCharToLower(t.Name()) + ` *` + t.Name() + ") Scan(data map[string]any) (err error) {\n")
		fn(`	for key, value := range data {` + "\n")
		fn(`		switch key {` + "\n")
		for k, v := range field {
			typeStruct := strings.Split(v, "|")[1]
			fn(`		case ` + JOstr.FirstCharToLower(t.Name()) + `.Field` + k + `():` + "\n")
			switch typeStruct {
			case "Jsql.NullString":
				fn(`			val, ok := value.(` + "string)\n")
				fn(`			` + "if !ok {\n")
				fn(`				` + `return errors.New("invalid type string. field ` + k + `")` + "\n")
			case "Jsql.NullBool":
				fn(`			val, ok := value.(` + "bool)\n")
				fn(`			` + "if !ok {\n")
				fn(`				` + `return errors.New("invalid type bool. field ` + k + `")` + "\n")
			case "Jsql.NullFloat64":
				fn(`			val, ok := value.(` + "float64)\n")
				fn(`			` + "if !ok {\n")
				fn(`				` + `return errors.New("invalid type float64. field ` + k + `")` + "\n")
			case "Jsql.NullInt64":
				fn(`			val, ok := value.(` + "int64)\n")
				fn(`			` + "if !ok {\n")
				fn(`				` + `return errors.New("invalid type int64. field ` + k + `")` + "\n")
			case "Jsql.NullInt32":
				fn(`			val, ok := value.(` + "int32)\n")
				fn(`			` + "if !ok {\n")
				fn(`				` + `return errors.New("invalid type int32. field ` + k + `")` + "\n")
			case "Jsql.NullByte":
				fn(`			val, ok := value.(` + "byte)\n")
				fn(`			` + "if !ok {\n")
				fn(`				` + `return errors.New("invalid type byte. field ` + k + `")` + "\n")
			case "Jsql.NullTime":
				fn(`			val, ok := value.(` + "time.Time)\n")
				fn(`			` + "if !ok {\n")
				fn(`				` + `return errors.New("invalid type time.Time. field ` + k + `")` + "\n")
			case "Jsql.NullInt16":
				fn(`			val, ok := value.(` + "int16)\n")
				fn(`			` + "if !ok {\n")
				fn(`				` + `return errors.New("invalid type int16. field ` + k + `")` + "\n")
			default:
				fn(`			val, ok := value.(` + typeStruct + ")\n")
				fn(`			` + "if !ok {\n")
				fn(`				` + `return errors.New("invalid type ` + typeStruct + `. field ` + k + `")` + "\n")
			}
			fn(`			` + "}\n")
			fn(`			` + JOstr.FirstCharToLower(t.Name()) + `.Set` + k + "(val)\n")
		}
		fn(`		default:` + "\n")
		fn(`			return errors.New("invalid column")` + "\n")
		fn(`		}` + "\n")
		fn(`	}` + "\n")
		fn(`	return nil` + "\n")
		fn("}\n\n")

		fn("// SetColumn is a function to set the column to QColumnFields for that will be used in the query.\n")
		fn(`func (` + JOstr.FirstCharToLower(t.Name()) + ` *` + t.Name() + ") SetColumn(columns ...string) (err error) {\n")
		fn(`	for _, column := range columns {` + "\n")
		fn(`		switch column {` + "\n")
		for k := range field {
			fn(`		case ` + JOstr.FirstCharToLower(t.Name()) + `.Field` + k + `():` + "\n")
		}
		fn(`		default:` + "\n")
		fn(`			` + "return errors.New(\"invalid column\")\n")
		fn(`		}` + "\n")
		fn(`		cond := false` + "\n")
		fn(`		for _, field := range ` + JOstr.FirstCharToLower(t.Name()) + ".QColumnFields " + "{\n")
		fn(`			if column == field {` + "\n")
		fn(`				cond = true` + "\n")
		fn(`				break` + "\n")
		fn(`			}` + "\n")
		fn(`		}` + "\n")
		fn(`		if cond == true {` + "\n")
		fn(`			continue` + "\n")
		fn(`		}` + "\n")
		fn(`		` + JOstr.FirstCharToLower(t.Name()) + ".QColumnFields = append(" +
			JOstr.FirstCharToLower(t.Name()) + ".QColumnFields, column)" + "\n")
		fn(`	}` + "\n")
		fn(`	` + "return nil\n")
		fn("}\n\n")

		fn("// DeleteColumnFromQColumnFields is a function to delete the column from QColumnFields for that will be used in the query.\n")
		fn(`func (` + JOstr.FirstCharToLower(t.Name()) + ` *` + t.Name() + ") DeleteColumnFromQColumnFields(elems ...string) (err error) {\n")
		fn(`	var colums []string` + "\n")
		fn(`	for _, v := range ` + JOstr.FirstCharToLower(t.Name()) + ".QColumnFields" + `{` + "\n")
		fn(`		colums = append(colums, v)` + "\n")
		fn(`	}` + "\n\n")
		fn(`	for _, elem := range elems {` + "\n")
		fn(`		index := -1` + "\n")
		fn(`		for i, column := range ` + JOstr.FirstCharToLower(t.Name()) + ".QColumnFields" + `{` + "\n")
		fn(`			if column == elem {` + "\n")
		fn(`				index = i` + "\n")
		fn(`				break` + "\n")
		fn(`			}` + "\n")
		fn(`		}` + "\n")
		fn(`		if index == -1 {` + "\n")
		fn(`			return fmt.Errorf("column %s not found", elem)` + "\n")
		fn(`		}` + "\n")
		fn(`		colums = append(colums[:index], colums[index+1:]...)` + "\n")
		fn(`	}` + "\n")
		fn(`	` + JOstr.FirstCharToLower(t.Name()) + ".QColumnFields" + ` = colums` + "\n")
		fn(`	return nil` + "\n")
		fn("}\n\n")

		fn("// QColumnFieldToStrings is a function to get the column format string from QColumnFields for that will be used in the query.\n")
		fn(`func (` + JOstr.FirstCharToLower(t.Name()) + ` *` + t.Name() + ") QColumnFieldToStrings() (columnStr string) {\n")
		fn(`	return strings.Join(` + JOstr.FirstCharToLower(t.Name()) + ".QColumnFields, \", \")\n")
		fn("}\n\n")

		fn("// ResetQColumnFields is a function to reset QColumnFields.\n")
		fn(`func (` + JOstr.FirstCharToLower(t.Name()) + ` *` + t.Name() + ") ResetQColumnFields() {\n")
		fn(`	` + JOstr.FirstCharToLower(t.Name()) + ".QColumnFields = []string{}\n")
		fn("}\n\n")

		fn("// ResetQFilterNamedArgs is a function to reset QFilterNamedArgs.\n")
		fn(`func (` + JOstr.FirstCharToLower(t.Name()) + ` *` + t.Name() + ") ResetQFilterNamedArgs() {\n")
		fn(`	` + JOstr.FirstCharToLower(t.Name()) + ".QFilterNamedArgs = Jsql.QFilterNamedArgs{}\n")
		fn("}\n\n")

		fn("// Locking is a function to set locking method.\n")
		fn(`func (` + JOstr.FirstCharToLower(t.Name()) + ` *` + t.Name() + ") Locking(lockingOperator Jsql.LockingOperator) Jsql.LockingOperator {\n")
		fn(`	return lockingOperator` + "\n")
		fn("}\n\n")

		fn("// FieldAndValue is  function for get named arg for write query\n")
		fn(`func (` + JOstr.FirstCharToLower(t.Name()) + ` *` + t.Name() + ") FieldAndValue() JOmap.SA {\n")
		fn(`	sa := make(JOmap.SA)` + "\n")
		fn(`	for _, field := range u.QColumnFields {` + "\n")
		fn(`		switch field {` + "\n")
		for k, _ := range field {
			fn(`		case ` + JOstr.FirstCharToLower(t.Name()) + ".Field" + k + "():\n")
			fn(`			sa[field] = ` + JOstr.FirstCharToLower(t.Name()) + "." + k + "\n")
		}
		fn(`		}` + "\n")
		fn("	}\n")
		fn(`	return sa` + "\n")
		fn("}\n\n")

		fn("// FieldArgForUpdate is function get string to SET update\n")
		fn(`func (` + JOstr.FirstCharToLower(t.Name()) + ` *` + t.Name() + ") FieldArgForUpdate(prefix Jsql.PrefixNamedArgPG) string {\n")
		fn(`	str := ""` + "\n")
		fn(`	columns := ` + JOstr.FirstCharToLower(t.Name()) + ".FieldAndValue()\n")
		fn(`	i := 1` + "\n")
		fn(`	for k, _ := range columns {` + "\n")
		fn(`		if i == len(columns) {` + "\n")
		fn(`			str += k + " = " + string(prefix) + k` + "\n")
		fn(`		` + "} else {\n")
		fn(`			str += k + " = " + string(prefix) + k + ", "` + "\n")
		fn(`		` + "}\n")
		fn(`		i++` + "\n")
		fn(`	` + "}\n")
		fn(`	return str` + "\n")
		fn("}\n\n")
		err := os.WriteFile(param.FileName+"_GEN.go", buf.Bytes(), os.ModePerm)
		jolly.PanicIF(err)

		if param.WithInterfaceRepo.Make {
			makeInterfaceRepo(param.WithInterfaceRepo, param.Src, caller)
		}
		if param.WithInterfaceImplRepo.Make {
			makeInterfaceRepoImpl(param, caller)
		}
	}
}

func makeInterfaceRepo(param WithInterfaceRepo, src any, caller *Jlog.CallInfo) {
	structName := ""
	v := reflect.TypeOf(src)

	if v.Kind() == reflect.Ptr {
		structName = v.Elem().Name()
	} else {
		structName = v.Name()
	}

	callerPackageName := JOstr.LastStringOfSubStr(caller.PackageName, "/")

	thisCaller := Jlog.CallerInfo()
	lib := "pgx"

	if param.Lib != "" {
		lib = param.Lib
	}

	buf := bytes.Buffer{}
	fn := func(str string) {
		_, err := buf.WriteString(str)
		jolly.PanicIF(err)
	}

	dirSplit := strings.Split(param.Dir, "/")
	fn("package " + dirSplit[len(dirSplit)-1] + "\n\n")

	fn(`import (` + "\n")
	fn(`	"context"` + "\n")
	fn("\n")
	fn(`	"` + thisCaller.PackageName + `"` + "\n")
	fn(`	"` + caller.PackageName + `"` + "\n")
	fn(")\n\n")

	fn(`// ` + param.InterfaceName + `GEN is a interface for ` + param.InterfaceName + `.` + " \n")
	fn("// DO NOT EDIT, will be overwritten by https://github.com/SyaibanAhmadRamadhan/jolly/blob/main/Jdb/JOpg/postgres_generator.go. \n")
	fn(`type ` + param.InterfaceName + `GEN interface {` + "\n")

	fn(`	// Create And Update Is abstract for write command` + "\n")
	fn(`	Create(ctx context.Context, ` + JOstr.ToLower(structName) + ` *` +
		callerPackageName + "." + structName + `) (err error)` + "\n")
	fn(`	Update(ctx context.Context, ` + JOstr.ToLower(structName) + ` *` +
		callerPackageName + "." + structName + `) (err error)` + "\n\n")

	fn(`	// FindOne And FindAll And Count Is abstract for read query` + "\n")
	fn(`	FindOne(ctx context.Context, ` + JOstr.ToLower(structName) + ` *` +
		callerPackageName + "." + structName + `) ` +
		`(err error)` + "\n")

	fn(`	FindAll(ctx context.Context, ` + JOstr.ToLower(structName) + ` *` +
		callerPackageName + "." + structName + `) ` +
		`(` + JOstr.ToLower(structName) + `s []` + callerPackageName + "." + structName + `, err error)` + "\n")

	fn(`	Count(ctx context.Context, ` + JOstr.ToLower(structName) + ` *` +
		callerPackageName + "." + structName + `) ` +
		`(total int64 ,err error)` + "\n\n")

	fn(`	// WithTx return interface ` + param.InterfaceName + " for command or query with the same transaction" + "\n")
	fn(`	WithTx(` + path.Base(thisCaller.PackageName) + "." + "RDBMS" + lib + `) ` + param.InterfaceName + "\n\n")

	fn(`	// ` + "RDBMS" + lib + " contract for command or query with library https://github.com/jackc/pgx" + "\n")
	fn(`	` + path.Base(thisCaller.PackageName) + "." + "RDBMS" + lib + "\n")
	fn(`}` + "\n\n")

	fileLocation := param.Dir + "/" + param.InterfaceName + ".go"
	err := os.MkdirAll(param.Dir, os.ModePerm)
	jolly.PanicIF(err)

	if _, err := os.Stat(fileLocation); !os.IsNotExist(err) {
		data, err := os.ReadFile(fileLocation)
		if err != nil {
			panic(err)
		}
		contentToCheck := []byte(`type ` + param.InterfaceName + ` interface {`)
		if bytes.Contains(data, contentToCheck) {
			fn(`// ` + param.InterfaceName + ` is a interface for ` + param.InterfaceName + `.` + " your CAN EDIT. but not edit name interface\n")
			pattern := fmt.Sprintf(`(?ms:type %s interface\s*.*?})`, param.InterfaceName)
			r := regexp.MustCompile(pattern)
			buf.WriteString(r.FindString(string(data)))
		} else {
			fn(`// ` + param.InterfaceName + ` is a interface for ` + param.InterfaceName + `.` + " your CAN EDIT. but not edit name interface\n")
			fn(`type ` + param.InterfaceName + ` interface {` + "\n")
			fn(`	// ` + "You can customize the abstract below... \n")
			fn("\n	// " + param.InterfaceName +
				"GEN DO NOT DELETE THIS, will be overwritten by https://github.com/SyaibanAhmadRamadhan/jolly/blob/main/Jdb/JOpg/postgres_generator.go. \n")
			fn("" + `	` + param.InterfaceName + "GEN \n")
			fn(`}` + "\n\n")
		}
	} else {
		fn(`// ` + param.InterfaceName + ` is a interface for ` + param.InterfaceName + `.` + " your CAN EDIT. but not edit name interface\n")
		fn(`type ` + param.InterfaceName + ` interface {` + "\n")
		fn("\n	// " + param.InterfaceName +
			"GEN DO NOT DELETE THIS, will be overwritten by https://github.com/SyaibanAhmadRamadhan/jolly/blob/main/Jdb/JOpg/postgres_generator.go. \n")
		fn(`	` + param.InterfaceName + "GEN \n")
		fn(`}` + "\n\n")
	}

	err = os.WriteFile(fileLocation, buf.Bytes(), os.ModePerm)
	jolly.PanicIF(err)
}

func makeInterfaceRepoImpl(param GeneratorModelForStructParam, caller *Jlog.CallInfo) {
	structName := ""
	v := reflect.TypeOf(param.Src)

	if v.Kind() == reflect.Ptr {
		structName = v.Elem().Name()
	} else {
		structName = v.Name()
	}

	fileDir, err := JOFD.FindDirPathFileFromGoModule(param.WithInterfaceRepo.InterfaceName + ".go")
	jolly.PanicIF(err)

	funcNewName := strings.ToUpper(string(param.WithInterfaceImplRepo.ImplInterfaceName[0])) + param.WithInterfaceImplRepo.ImplInterfaceName[1:]
	interfaceRepoPackageName := JOstr.LastStringOfSubStr(param.WithInterfaceRepo.Dir, "/")
	interfaceImplName := param.WithInterfaceImplRepo.ImplInterfaceName
	interfaceName := param.WithInterfaceRepo.InterfaceName

	callerPackageName := JOstr.LastStringOfSubStr(caller.PackageName, "/")

	moduleName, err := JOFD.GetModuleName()
	jolly.PanicIF(err)

	lib := "pgx"
	if param.WithInterfaceRepo.Lib != "" {
		lib = param.WithInterfaceRepo.Lib
	}
	thisCaller := Jlog.CallerInfo()

	buf := bytes.Buffer{}
	fn := func(str string) {
		_, err := buf.WriteString(str)
		jolly.PanicIF(err)
	}

	dirSplit := strings.Split(param.WithInterfaceImplRepo.ImplLocationDir, "/")
	fn(`package ` + dirSplit[len(dirSplit)-1] + "\n\n")

	fn(`import (` + "\n")
	fn(`	"context"` + "\n")
	fn("\n")
	fn(`	"` + thisCaller.PackageName + `"` + "\n")
	fn(`	"` + caller.PackageName + `"` + "\n")
	if param.WithInterfaceImplRepo.ImplLocationDir != param.WithInterfaceRepo.Dir {
		fn(`	"` + moduleName + "/" + fileDir + `"` + "\n")
	}
	fn(")\n\n")

	fn("// DO NOT EDIT, will be overwritten by https://github.com/SyaibanAhmadRamadhan/jolly/blob/main/Jdb/JOpg/postgres_generator.go. \n\n")

	fn(`type ` + param.WithInterfaceImplRepo.ImplInterfaceName + `GEN struct {` + "\n")
	fn(`	` + path.Base(thisCaller.PackageName) + "." + "RDBMS" + lib + "\n")
	fn(`}` + "\n\n")

	if param.WithInterfaceImplRepo.ImplLocationDir == param.WithInterfaceRepo.Dir {
		fn(`func New` + funcNewName + `GEN(` + JOstr.FirstCharToLower(param.WithInterfaceImplRepo.ImplInterfaceName) + ` ` +
			path.Base(thisCaller.PackageName) + "." + "RDBMS" + lib + `) ` +
			param.WithInterfaceRepo.InterfaceName + `GEN {` + "\n")
	} else {
		fn(`func New` + funcNewName + `GEN(` + JOstr.FirstCharToLower(param.WithInterfaceImplRepo.ImplInterfaceName) + ` ` +
			path.Base(thisCaller.PackageName) + "." + "RDBMS" + lib + `) ` +
			interfaceRepoPackageName + "." + param.WithInterfaceRepo.InterfaceName + `GEN {` + "\n")
	}
	fn(`	return &` + param.WithInterfaceImplRepo.ImplInterfaceName + `GEN{` + "\n")
	fn(`		RDBMS` + lib + ": " +
		JOstr.FirstCharToLower(param.WithInterfaceImplRepo.ImplInterfaceName) + "," + "\n")
	fn(`	}` + "\n")
	fn(`}` + "\n\n")

	fn(`func (` + JOstr.FirstCharToLower(interfaceImplName) + ` *` + interfaceImplName + `GEN) Create(ctx context.Context, ` +
		JOstr.ToLower(structName) + ` *` + callerPackageName + "." + structName + `) (err error) {` + "\n")
	fn(`	return` + "\n")
	fn(`}` + "\n\n")

	fn(`func (` + JOstr.FirstCharToLower(interfaceImplName) + ` *` + interfaceImplName + `GEN) Update(ctx context.Context, ` +
		JOstr.ToLower(structName) + ` *` + callerPackageName + "." + structName + `) (err error) {` + "\n")
	fn(`	return` + "\n")
	fn(`}` + "\n\n")

	fn(`func (` + JOstr.FirstCharToLower(interfaceImplName) + ` *` + interfaceImplName + `GEN) FindOne(ctx context.Context, ` +
		JOstr.ToLower(structName) + ` *` + callerPackageName + "." + structName + `) (err error) {` + "\n")
	fn(`	return` + "\n")
	fn(`}` + "\n\n")

	fn(`func (` + JOstr.FirstCharToLower(interfaceImplName) + ` *` + interfaceImplName + `GEN) FindAll(ctx context.Context, ` +
		JOstr.ToLower(structName) + ` *` + callerPackageName + "." + structName + `) (` +
		JOstr.ToLower(structName) + "s []" + callerPackageName + "." + structName + `, err error) {` + "\n")
	fn(`	return` + "\n")
	fn(`}` + "\n\n")

	fn(`func (` + JOstr.FirstCharToLower(interfaceImplName) + ` *` + interfaceImplName + `GEN) Count(ctx context.Context, ` +
		JOstr.ToLower(structName) + ` *` + callerPackageName + "." + structName + `) (total int64, err error) {` + "\n")
	fn(`	return` + "\n")
	fn(`}` + "\n\n")

	fn(`func (` + JOstr.FirstCharToLower(interfaceImplName) + ` *` + interfaceImplName + `GEN) WithTx(j ` +
		path.Base(thisCaller.PackageName) + "." + "RDBMS" + lib + `) ` + interfaceRepoPackageName + "." + interfaceName + ` {` + "\n")
	fn(`	return &` + param.WithInterfaceImplRepo.ImplInterfaceName + `GEN{` + "\n")
	fn(`		RDBMS` + lib + ": " + "j," + "\n")
	fn(`	}` + "\n")
	fn(`}` + "\n\n")

	fileLocation := param.WithInterfaceImplRepo.ImplLocationDir + "/" + interfaceImplName + "_GEN.go"
	err = os.MkdirAll(param.WithInterfaceImplRepo.ImplLocationDir, os.ModePerm)
	jolly.PanicIF(err)

	err = os.WriteFile(fileLocation, buf.Bytes(), os.ModePerm)
	jolly.PanicIF(err)
}
