package gdb

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strings"

	"github.com/SyaibanAhmadRamadhan/gocatch/gcommon"
	"github.com/SyaibanAhmadRamadhan/gocatch/glog"
	"github.com/SyaibanAhmadRamadhan/gocatch/gstr"
	"github.com/SyaibanAhmadRamadhan/gocatch/gstruct"
)

// SpecifiationTable is a struct that contains the specification of the table
// schema name is optional
// TableName is table name in sql or collection name in nosql or index name in search db
type SpecifiationTable struct {
	TableName  string
	SchemaName string
}

// GeneratorModelForStructParam defines the parameters
// for generating methods for a struct
type GeneratorModelForStructParam struct {
	Src               any               // Src is the struct instance
	SpecifiationTable SpecifiationTable // SpecifiationTable is the struct instance that contains the specification of the table
	Tag               string            // by default tag is db tag
	FileName          string            // FileName is the name of generated source file without extension
}

// GeneratorModelFromStruct generates methods for given structs.
// This function expects each struct to contain a field `RQField []string`.
// The generated methods manipulate or make use of this field for various operations.
// The function does not support nested structs with name tag.
// if you want nested struct for this function, you can use "-" tag in your struct.
// for example check in folder Jdb/JOpg/exGen
func GeneratorModelFromStruct(params ...GeneratorModelForStructParam) {
	for _, param := range params {
		if param.SpecifiationTable.TableName == "" {
			panic("TableName is empty")
		}

		// if param.WithInterfaceRepo.Make {
		// 	if param.WithInterfaceRepo.InterfaceName == "" {
		// 		panic("InterfaceName is empty")
		// 	}
		//
		// 	if param.WithInterfaceRepo.Dir == "" {
		// 		panic("Dir is empty")
		// 	}
		// }
		//
		// if param.WithInterfaceImplRepo.Make {
		// 	if param.WithInterfaceImplRepo.ImplInterfaceName == "" {
		// 		panic("InterfaceName is empty")
		// 	}
		//
		// 	if param.WithInterfaceImplRepo.ImplLocationDir == "" {
		// 		panic("Dir is empty")
		// 	}
		// }

		buf := bytes.Buffer{}
		fn := func(str string) {
			_, err := buf.WriteString(str)
			gcommon.PanicIfError(err)
		}

		t := reflect.TypeOf(param.Src).Elem()
		caller := glog.CallerInfo(2)
		packageName := gstr.GetLastSubstring(caller.PackageName, "/")
		fn(`package ` + packageName + "\n\n")

		fn("// DO NOT EDIT, will be overwritten by https://github.com/SyaibanAhmadRamadhan/jolly/blob/main/Jdb/JOpg/postgres_generator.go. \n\n")
		// Prepare import statements
		res, _ := os.ReadFile(param.FileName + ".go")
		r, _ := regexp.Compile(`import \(([^)]+)\)`)
		importSrc := r.FindString(string(res))

		if importSrc == "" {
			importSrc = fmt.Sprintf(
				`import (%s%s`,
				"\n\t\"errors\"",
				// "\n\t\"fmt\"",
				// "\n\t\"strings\"",
				"\n)",
			)
		} else {
			customImport := fmt.Sprintf(
				`(%s%s`,
				"\n\t\"errors\"",
				// "\n\t\"fmt\"",
				// "\n\t\"strings\"",
				"\n",
			)
			importSrc = strings.Replace(importSrc, "(", customImport, 1)
		}

		fn(importSrc + "\n\n")

		fn(`// ` + t.Name() + "TableName this table or collection name\n")
		fn(`const ` + t.Name() + `TableName string = "` + param.SpecifiationTable.TableName + "\"\n\n")

		if param.SpecifiationTable.SchemaName != "" {
			fn("// " + t.Name() + "SchemaName is a variable schema name\n")
			fn(`const ` + t.Name() + `SchemaName string = "` + param.SpecifiationTable.SchemaName + "\"\n\n")
		}

		fn("// New" + t.Name() + " is a struct with pointer that represents the table " + t.Name() + " in the database.\n")
		fn(`func New` + t.Name() + `() *` + t.Name() + " {\n")
		fn(`	return &` + t.Name() + "{}\n")
		fn("}\n\n")

		fn("// New" + t.Name() + "WithOutPtr is a struct without pointer that represents the table " + t.Name() + " in the database.\n")
		fn(`func New` + t.Name() + `WithOutPtr() ` + t.Name() + " {\n")
		fn(`	return ` + t.Name() + "{}\n")
		fn("}\n\n")

		// fn("// TableName is a function to get table name\n")
		// fn(`func (` + gstr.LowercaseFirstChar(t.Name()) + ` *` + t.Name() + ") TableName() (table string) {\n")
		// fn(`	return "` + param.SpecifiationTable.TableName + "\"\n")
		// fn("}\n\n")

		// if param.SpecifiationTable.SchemaName != "" {
		// 	fn("// SchemaName is a function to get schema name\n")
		// 	fn(`func (` + gstr.LowercaseFirstChar(t.Name()) + ` *` + t.Name() + ") SchemaName() (schema string) {\n")
		// 	fn(`	return "` + param.SpecifiationTable.SchemaName + "\"\n")
		// 	fn("}\n\n")
		// }

		gsqlNullString := "gsql.NullString"
		gsqlNullBool := "gsql.NullBool"
		gsqlNullFloat64 := "gsql.NullFloat64"
		gsqlNullInt64 := "gsql.NullInt64"
		gsqlNullInt32 := "gsql.NullInt32"
		gsqlNullByte := "gsql.NullByte"
		gsqlNullTime := "gsql.NullTime"
		gsqlNullInt16 := "gsql.NullInt16"

		if param.Tag == "" {
			param.Tag = "db"
		}
		field := gstruct.ExtractStructTagsAndFields(param.Src, "", param.Tag)
		for k, v := range field {
			typeStruct := strings.Split(v, "|")[1]

			// FieldName
			fn("// Field" + k + " is a field or column in the table " + t.Name() + ".\n")
			fn(`func (` + gstr.LowercaseFirstChar(t.Name()) + ` *` + t.Name() + ") Field" + k + "() string {\n")
			fn(`	return "` + strings.Split(v, "|")[0] + "\"\n")
			fn("}\n\n")

			// SetField
			fn("// Set" + k + " is a setter for the field or column " + k + " in the table " + t.Name() + ".\n")
			switch typeStruct {
			case gsqlNullString:
				fn(`func (` + gstr.LowercaseFirstChar(t.Name()) + ` *` + t.Name() + ") Set" + k + "(param string) {\n")
			case gsqlNullBool:
				fn(`func (` + gstr.LowercaseFirstChar(t.Name()) + ` *` + t.Name() + ") Set" + k + "(param bool) {\n")
			case gsqlNullFloat64:
				fn(`func (` + gstr.LowercaseFirstChar(t.Name()) + ` *` + t.Name() + ") Set" + k + "(param float64) {\n")
			case gsqlNullInt64:
				fn(`func (` + gstr.LowercaseFirstChar(t.Name()) + ` *` + t.Name() + ") Set" + k + "(param int64) {\n")
			case gsqlNullInt32:
				fn(`func (` + gstr.LowercaseFirstChar(t.Name()) + ` *` + t.Name() + ") Set" + k + "(param int32) {\n")
			case gsqlNullByte:
				fn(`func (` + gstr.LowercaseFirstChar(t.Name()) + ` *` + t.Name() + ") Set" + k + "(param byte) {\n")
			case gsqlNullTime:
				fn(`func (` + gstr.LowercaseFirstChar(t.Name()) + ` *` + t.Name() + ") Set" + k + "(param time.Time) {\n")
			case gsqlNullInt16:
				fn(`func (` + gstr.LowercaseFirstChar(t.Name()) + ` *` + t.Name() + ") Set" + k + "(param int16) {\n")
			default:
				fn(`func (` + gstr.LowercaseFirstChar(t.Name()) + ` *` + t.Name() + ") Set" + k + "(param " + strings.Split(v, "|")[1] + ") {\n")
			}

			if typeStruct == gsqlNullString ||
				typeStruct == gsqlNullBool ||
				typeStruct == gsqlNullFloat64 ||
				typeStruct == gsqlNullInt64 ||
				typeStruct == gsqlNullInt32 ||
				typeStruct == gsqlNullByte ||
				typeStruct == gsqlNullTime ||
				typeStruct == gsqlNullInt16 {
				fn(`	` + gstr.LowercaseFirstChar(t.Name()) + "." + k + " = " + strings.ReplaceAll(typeStruct, ".", ".New") + "(&param)" + "\n")
			} else {
				fn(`	` + gstr.LowercaseFirstChar(t.Name()) + "." + k + " = param" + "\n")
			}
			fn("}\n\n")
		}

		fn("// AllField is a function to get all field or column in the table " + t.Name() + ".\n")
		fn(`func (` + gstr.LowercaseFirstChar(t.Name()) + ` *` + t.Name() + ") AllField() (str []string) {\n")
		fn(`	str = []string{ ` + "\n")
		for _, v := range field {
			fn("\t\t`" + strings.Split(v, "|")[0] + "`" + ",\n")
		}
		fn(`	}` + "\n")
		fn(`	return` + "\n")
		fn("}\n\n")

		order := gstruct.ExtractStructTagsAndFields(param.Src, "", "order")
		fn("// OrderFields is a function to get all field or column in the table " + t.Name() + ".\n")
		fn(`func (` + gstr.LowercaseFirstChar(t.Name()) + ` *` + t.Name() + ") OrderFields() (str []string) {\n")
		fn(`	str = []string{ ` + "\n")
		for k, v := range order {
			if strings.Split(v, "|")[0] == "true" {
				for k1, v1 := range field {
					if k1 == k {
						fn("\t\t`" + strings.Split(v1, "|")[0] + "`" + ",\n")
					}
				}
			}
		}
		fn(`	}` + "\n")
		fn(`	return` + "\n")
		fn("}\n\n")

		fn("// GetValuesByColums is a function to get all value by column in the table " + t.Name() + ".\n")
		fn(`func (` + gstr.LowercaseFirstChar(t.Name()) + ` *` + t.Name() + ") GetValuesByColums(columns ...string) []any {\n")
		fn(`	var values []any` + "\n")
		fn(`	for _, column := range columns {` + "\n")
		fn(`		switch column {` + "\n")
		for k := range field {
			fn(`		case ` + gstr.LowercaseFirstChar(t.Name()) + `.Field` + k + `():` + "\n")
			fn(`			values = append(values, ` + gstr.LowercaseFirstChar(t.Name()) + "." + k + ")" + "\n")
		}
		fn(`		}` + "\n")
		fn(`	}` + "\n")
		fn(`	return values` + "\n")
		fn("}\n\n")

		fn("// ScanMap is a function to scan the value with for rows.Value() from the database to the struct " + t.Name() + ".\n")
		fn(`func (` + gstr.LowercaseFirstChar(t.Name()) + ` *` + t.Name() + ") ScanMap(data map[string]any) (err error) {\n")
		fn(`	for key, value := range data {` + "\n")
		fn(`		switch key {` + "\n")
		for k, v := range field {
			typeStruct := strings.Split(v, "|")[1]
			fn(`		case ` + gstr.LowercaseFirstChar(t.Name()) + `.Field` + k + `():` + "\n")
			switch typeStruct {
			case gsqlNullString:
				fn(`			val, ok := value.(` + "string)\n")
				fn(`			` + "if !ok {\n")
				fn(`				` + `return errors.New("invalid type string. field ` + k + `")` + "\n")
			case gsqlNullBool:
				fn(`			val, ok := value.(` + "bool)\n")
				fn(`			` + "if !ok {\n")
				fn(`				` + `return errors.New("invalid type bool. field ` + k + `")` + "\n")
			case gsqlNullFloat64:
				fn(`			val, ok := value.(` + "float64)\n")
				fn(`			` + "if !ok {\n")
				fn(`				` + `return errors.New("invalid type float64. field ` + k + `")` + "\n")
			case gsqlNullInt64:
				fn(`			val, ok := value.(` + "int64)\n")
				fn(`			` + "if !ok {\n")
				fn(`				` + `return errors.New("invalid type int64. field ` + k + `")` + "\n")
			case gsqlNullInt32:
				fn(`			val, ok := value.(` + "int32)\n")
				fn(`			` + "if !ok {\n")
				fn(`				` + `return errors.New("invalid type int32. field ` + k + `")` + "\n")
			case gsqlNullByte:
				fn(`			val, ok := value.(` + "byte)\n")
				fn(`			` + "if !ok {\n")
				fn(`				` + `return errors.New("invalid type byte. field ` + k + `")` + "\n")
			case gsqlNullTime:
				fn(`			val, ok := value.(` + "time.Time)\n")
				fn(`			` + "if !ok {\n")
				fn(`				` + `return errors.New("invalid type time.Time. field ` + k + `")` + "\n")
			case gsqlNullInt16:
				fn(`			val, ok := value.(` + "int16)\n")
				fn(`			` + "if !ok {\n")
				fn(`				` + `return errors.New("invalid type int16. field ` + k + `")` + "\n")
			default:
				fn(`			val, ok := value.(` + typeStruct + ")\n")
				fn(`			` + "if !ok {\n")
				fn(`				` + `return errors.New("invalid type ` + typeStruct + `. field ` + k + `")` + "\n")
			}
			fn(`			` + "}\n")
			fn(`			` + gstr.LowercaseFirstChar(t.Name()) + `.Set` + k + "(val)\n")
		}
		fn(`		default:` + "\n")
		fn(`			return errors.New("invalid column")` + "\n")
		fn(`		}` + "\n")
		fn(`	}` + "\n")
		fn(`	return nil` + "\n")
		fn("}\n\n")

		// fn("// RQFieldSet is a function to set the column to RQField for that will be used in the query.\n")
		// fn(`func (` + gstr.LowercaseFirstChar(t.Name()) + ` *` + t.Name() + ") RQFieldSet(columns ...string) (err error) {\n")
		// fn(`	for _, column := range columns {` + "\n")
		// fn(`		switch column {` + "\n")
		// for k := range field {
		// 	fn(`		case ` + gstr.LowercaseFirstChar(t.Name()) + `.Field` + k + `():` + "\n")
		// }
		// fn(`		default:` + "\n")
		// fn(`			` + "return errors.New(\"invalid column\")\n")
		// fn(`		}` + "\n")
		// fn(`		cond := false` + "\n")
		// fn(`		for _, field := range ` + gstr.LowercaseFirstChar(t.Name()) + ".RQField " + "{\n")
		// fn(`			if column == field {` + "\n")
		// fn(`				cond = true` + "\n")
		// fn(`				break` + "\n")
		// fn(`			}` + "\n")
		// fn(`		}` + "\n")
		// fn(`		if cond == true {` + "\n")
		// fn(`			continue` + "\n")
		// fn(`		}` + "\n")
		// fn(`		` + gstr.LowercaseFirstChar(t.Name()) + ".RQField = append(" +
		// 	gstr.LowercaseFirstChar(t.Name()) + ".RQField, column)" + "\n")
		// fn(`	}` + "\n")
		// fn(`	` + "return nil\n")
		// fn("}\n\n")
		//
		// fn("// RQFieldDelete is a function to delete the column from RQField for that will be used in the query.\n")
		// fn(`func (` + gstr.LowercaseFirstChar(t.Name()) + ` *` + t.Name() + ") RQFieldDelete(elems ...string) (err error) {\n")
		// fn(`	var colums []string` + "\n")
		// fn(`	for _, v := range ` + gstr.LowercaseFirstChar(t.Name()) + ".RQField" + `{` + "\n")
		// fn(`		colums = append(colums, v)` + "\n")
		// fn(`	}` + "\n\n")
		// fn(`	for _, elem := range elems {` + "\n")
		// fn(`		index := -1` + "\n")
		// fn(`		for i, column := range ` + gstr.LowercaseFirstChar(t.Name()) + ".RQField" + `{` + "\n")
		// fn(`			if column == elem {` + "\n")
		// fn(`				index = i` + "\n")
		// fn(`				break` + "\n")
		// fn(`			}` + "\n")
		// fn(`		}` + "\n")
		// fn(`		if index == -1 {` + "\n")
		// fn(`			return fmt.Errorf("column %s not found", elem)` + "\n")
		// fn(`		}` + "\n")
		// fn(`		colums = append(colums[:index], colums[index+1:]...)` + "\n")
		// fn(`	}` + "\n")
		// fn(`	` + gstr.LowercaseFirstChar(t.Name()) + ".RQField" + ` = colums` + "\n")
		// fn(`	return nil` + "\n")
		// fn("}\n\n")
		//
		// fn("// RQFieldToString is a function to get the column format string from RQField for that will be used in the query.\n")
		// fn(`func (` + gstr.LowercaseFirstChar(t.Name()) + ` *` + t.Name() + ") RQFieldToString() (columnStr string) {\n")
		// fn(`	return strings.Join(` + gstr.LowercaseFirstChar(t.Name()) + ".RQField, \", \")\n")
		// fn("}\n\n")
		//
		// fn("// RQFieldReset is a function to reset RQField.\n")
		// fn(`func (` + gstr.LowercaseFirstChar(t.Name()) + ` *` + t.Name() + ") RQFieldReset() {\n")
		// fn(`	` + gstr.LowercaseFirstChar(t.Name()) + ".RQField = []string{}\n")
		// fn("}\n\n")
		//
		// fn("// FNamedArgsReset is a function to reset FNamedArgs.\n")
		// fn(`func (` + gstr.LowercaseFirstChar(t.Name()) + ` *` + t.Name() + ") FNamedArgsReset() {\n")
		// fn(`	` + gstr.LowercaseFirstChar(t.Name()) + ".FNamedArgs = nil\n")
		// fn("}\n\n")
		//
		// fn("// FNamedArgsSet is a function to set FNamedArgs.\n")
		// fn(`func (` + gstr.LowercaseFirstChar(t.Name()) + ` *` + t.Name() + ") FNamedArgsSet(param ...gsql.Filter) {\n")
		// fn(`	if len(` + gstr.LowercaseFirstChar(t.Name()) + ".FNamedArgs) < 1 {\n")
		// fn(`		if len(param) > 0` + " {\n")
		// fn(`			param[0].Logical = ""` + "\n")
		// fn("		}\n")
		// fn("	}\n")
		// fn(`	` + gstr.LowercaseFirstChar(t.Name()) + ".FNamedArgs = append(" + gstr.LowercaseFirstChar(t.Name()) + ".FNamedArgs, " + "param...)\n")
		// fn("}\n\n")

		err := os.WriteFile(param.FileName+"_GEN.go", buf.Bytes(), os.ModePerm)
		gcommon.PanicIfError(err)

		// if param.WithInterfaceRepo.Make {
		// 	makeInterfaceRepo(param.WithInterfaceRepo, param.Src, caller)
		// }
		// if param.WithInterfaceImplRepo.Make {
		// 	makeInterfaceRepoImpl(param, caller)
		// }
	}
}

// func makeInterfaceRepo(param WithInterfaceRepo, src any, caller *Jlog.CallInfo) {
// 	structName := ""
// 	v := reflect.TypeOf(src)
//
// 	if v.Kind() == reflect.Ptr {
// 		structName = v.Elem().Name()
// 	} else {
// 		structName = v.Name()
// 	}
//
// 	callerPackageName := gstr.GetLastSubstring(caller.PackageName, "/")
//
// 	thisCaller := Jlog.CallerInfo()
// 	lib := "pgx"
//
// 	if param.Lib != "" {
// 		lib = param.Lib
// 	}
//
// 	buf := bytes.Buffer{}
// 	fn := func(str string) {
// 		_, err := buf.WriteString(str)
// 		gcommon.PanicIfError(err)
// 	}
//
// 	dirSplit := strings.Split(param.Dir, "/")
// 	fn("package " + dirSplit[len(dirSplit)-1] + "\n\n")
//
// 	fn(`import (` + "\n")
// 	fn(`	"context"` + "\n")
// 	fn("\n")
// 	fn(`	"` + thisCaller.PackageName + `"` + "\n")
// 	fn(`	"` + caller.PackageName + `"` + "\n")
// 	fn(")\n\n")
//
// 	fn(`// ` + param.InterfaceName + `GEN is an interface for ` + param.InterfaceName + `.` + " \n")
// 	fn("// DO NOT EDIT, will be overwritten by https://github.com/SyaibanAhmadRamadhan/jolly/blob/main/Jdb/JOpg/postgres_generator.go. \n")
// 	fn(`type ` + param.InterfaceName + `GEN interface {` + "\n")
//
// 	fn(`	// Create And Update Is abstract for write command` + "\n")
// 	fn(`	Create(ctx context.Context, ` + gstr.ToLowercase(structName) + ` *` +
// 		callerPackageName + "." + structName + `) (err error)` + "\n")
// 	fn(`	Update(ctx context.Context, ` + gstr.ToLowercase(structName) + ` *` +
// 		callerPackageName + "." + structName + `) (err error)` + "\n\n")
//
// 	fn(`	// FindOne And FindAll And Count Is abstract for read query` + "\n")
// 	fn(`	FindOne(ctx context.Context, ` + gstr.ToLowercase(structName) + ` *` +
// 		callerPackageName + "." + structName + `) ` +
// 		`(err error)` + "\n")
//
// 	fn(`	FindAll(ctx context.Context, ` + gstr.ToLowercase(structName) + ` *` +
// 		callerPackageName + "." + structName + `) ` +
// 		`(` + gstr.ToLowercase(structName) + `s []` + callerPackageName + "." + structName + `, err error)` + "\n")
//
// 	fn(`	Count(ctx context.Context, ` + gstr.ToLowercase(structName) + ` *` +
// 		callerPackageName + "." + structName + `) ` +
// 		`(total int64 ,err error)` + "\n\n")
//
// 	fn(`	// WithTx return interface ` + param.InterfaceName + " for command or query with the same transaction" + "\n")
// 	fn(`	WithTx(` + path.Base(thisCaller.PackageName) + "." + "Tx" + lib + `) ` + param.InterfaceName + "\n\n")
//
// 	fn(`	// ` + "Tx" + lib + " contract for command or query with library https://github.com/jackc/pgx" + "\n")
// 	fn(`	` + path.Base(thisCaller.PackageName) + "." + "Tx" + lib + "\n")
// 	fn(`}` + "\n\n")
//
// 	fileLocation := param.Dir + "/" + param.InterfaceName + ".go"
// 	err := os.MkdirAll(param.Dir, os.ModePerm)
// 	gcommon.PanicIfError(err)
//
// 	if _, err := os.Stat(fileLocation); !os.IsNotExist(err) {
// 		data, err := os.ReadFile(fileLocation)
// 		if err != nil {
// 			panic(err)
// 		}
// 		contentToCheck := []byte(`type ` + param.InterfaceName + ` interface {`)
// 		if bytes.Contains(data, contentToCheck) {
// 			fn(`// ` + param.InterfaceName + ` is an interface for ` + param.InterfaceName + `.` + " your CAN EDIT. but not edit name interface\n")
// 			pattern := fmt.Sprintf(`(?ms:type %s interface\s*.*?})`, param.InterfaceName)
// 			r := regexp.MustCompile(pattern)
// 			buf.WriteString(r.FindString(string(data)))
// 		} else {
// 			fn(`// ` + param.InterfaceName + ` is an interface for ` + param.InterfaceName + `.` + " your CAN EDIT. but not edit name interface\n")
// 			fn(`type ` + param.InterfaceName + ` interface {` + "\n")
// 			fn(`	// ` + "You can customize the abstract below... \n")
// 			fn("\n	// " + param.InterfaceName +
// 				"GEN DO NOT DELETE THIS, will be overwritten by https://github.com/SyaibanAhmadRamadhan/jolly/blob/main/Jdb/JOpg/postgres_generator.go. \n")
// 			fn("" + `	` + param.InterfaceName + "GEN \n")
// 			fn(`}` + "\n\n")
// 		}
// 	} else {
// 		fn(`// ` + param.InterfaceName + ` is an interface for ` + param.InterfaceName + `.` + " your CAN EDIT. but not edit name interface\n")
// 		fn(`type ` + param.InterfaceName + ` interface {` + "\n")
// 		fn("\n	// " + param.InterfaceName +
// 			"GEN DO NOT DELETE THIS, will be overwritten by https://github.com/SyaibanAhmadRamadhan/jolly/blob/main/Jdb/JOpg/postgres_generator.go. \n")
// 		fn(`	` + param.InterfaceName + "GEN \n")
// 		fn(`}` + "\n\n")
// 	}
//
// 	err = os.WriteFile(fileLocation, buf.Bytes(), os.ModePerm)
// 	gcommon.PanicIfError(err)
// }
//
// func makeInterfaceRepoImpl(param GeneratorModelForStructParam, caller *Jlog.CallInfo) {
// 	structName := ""
// 	v := reflect.TypeOf(param.Src)
//
// 	if v.Kind() == reflect.Ptr {
// 		structName = v.Elem().Name()
// 	} else {
// 		structName = v.Name()
// 	}
//
// 	fileDir, err := gdir.FindDirPathOfFileFromGoMod(param.WithInterfaceRepo.InterfaceName + ".go")
// 	gcommon.PanicIfError(err)
//
// 	funcNewName := strings.ToUpper(string(param.WithInterfaceImplRepo.ImplInterfaceName[0])) + param.WithInterfaceImplRepo.ImplInterfaceName[1:]
// 	interfaceRepoPackageName := gstr.GetLastSubstring(param.WithInterfaceRepo.Dir, "/")
// 	interfaceImplName := param.WithInterfaceImplRepo.ImplInterfaceName
// 	interfaceName := param.WithInterfaceRepo.InterfaceName
//
// 	callerPackageName := gstr.GetLastSubstring(caller.PackageName, "/")
//
// 	moduleName, err := gdir.GetModuleName()
// 	gcommon.PanicIfError(err)
//
// 	lib := "pgx"
// 	if param.WithInterfaceRepo.Lib != "" {
// 		lib = param.WithInterfaceRepo.Lib
// 	}
// 	thisCaller := Jlog.CallerInfo()
//
// 	buf := bytes.Buffer{}
// 	fn := func(str string) {
// 		_, err := buf.WriteString(str)
// 		gcommon.PanicIfError(err)
// 	}
//
// 	dirSplit := strings.Split(param.WithInterfaceImplRepo.ImplLocationDir, "/")
// 	fn(`package ` + dirSplit[len(dirSplit)-1] + "\n\n")
//
// 	fn(`import (` + "\n")
// 	fn(`	"context"` + "\n")
// 	fn("\n")
// 	fn(`	"` + thisCaller.PackageName + `"` + "\n")
// 	fn(`	"` + caller.PackageName + `"` + "\n")
// 	if param.WithInterfaceImplRepo.ImplLocationDir != param.WithInterfaceRepo.Dir {
// 		fn(`	"` + moduleName + "/" + fileDir + `"` + "\n")
// 	}
// 	fn(")\n\n")
//
// 	fn("// DO NOT EDIT, will be overwritten by https://github.com/SyaibanAhmadRamadhan/jolly/blob/main/Jdb/JOpg/postgres_generator.go. \n\n")
//
// 	fn(`type ` + param.WithInterfaceImplRepo.ImplInterfaceName + `GEN struct {` + "\n")
// 	fn(`	` + path.Base(thisCaller.PackageName) + "." + "Tx" + lib + "\n")
// 	fn(`}` + "\n\n")
//
// 	if param.WithInterfaceImplRepo.ImplLocationDir == param.WithInterfaceRepo.Dir {
// 		fn(`func New` + funcNewName + `GEN(` + gstr.LowercaseFirstChar(param.WithInterfaceImplRepo.ImplInterfaceName) + ` ` +
// 			path.Base(thisCaller.PackageName) + "." + "Tx" + lib + `) ` +
// 			param.WithInterfaceRepo.InterfaceName + `GEN {` + "\n")
// 	} else {
// 		fn(`func New` + funcNewName + `GEN(` + gstr.LowercaseFirstChar(param.WithInterfaceImplRepo.ImplInterfaceName) + ` ` +
// 			path.Base(thisCaller.PackageName) + "." + "Tx" + lib + `) ` +
// 			interfaceRepoPackageName + "." + param.WithInterfaceRepo.InterfaceName + `GEN {` + "\n")
// 	}
// 	fn(`	return &` + param.WithInterfaceImplRepo.ImplInterfaceName + `GEN{` + "\n")
// 	fn(`		Tx` + lib + ": " +
// 		gstr.LowercaseFirstChar(param.WithInterfaceImplRepo.ImplInterfaceName) + "," + "\n")
// 	fn(`	}` + "\n")
// 	fn(`}` + "\n\n")
//
// 	fn(`func (` + gstr.LowercaseFirstChar(interfaceImplName) + ` *` + interfaceImplName + `GEN) Create(ctx context.Context, ` +
// 		gstr.ToLowercase(structName) + ` *` + callerPackageName + "." + structName + `) (err error) {` + "\n")
// 	fn(`	return` + "\n")
// 	fn(`}` + "\n\n")
//
// 	fn(`func (` + gstr.LowercaseFirstChar(interfaceImplName) + ` *` + interfaceImplName + `GEN) Update(ctx context.Context, ` +
// 		gstr.ToLowercase(structName) + ` *` + callerPackageName + "." + structName + `) (err error) {` + "\n")
// 	fn(`	return` + "\n")
// 	fn(`}` + "\n\n")
//
// 	fn(`func (` + gstr.LowercaseFirstChar(interfaceImplName) + ` *` + interfaceImplName + `GEN) FindOne(ctx context.Context, ` +
// 		gstr.ToLowercase(structName) + ` *` + callerPackageName + "." + structName + `) (err error) {` + "\n")
// 	fn(`	return` + "\n")
// 	fn(`}` + "\n\n")
//
// 	fn(`func (` + gstr.LowercaseFirstChar(interfaceImplName) + ` *` + interfaceImplName + `GEN) FindAll(ctx context.Context, ` +
// 		gstr.ToLowercase(structName) + ` *` + callerPackageName + "." + structName + `) (` +
// 		gstr.ToLowercase(structName) + "s []" + callerPackageName + "." + structName + `, err error) {` + "\n")
// 	fn(`	return` + "\n")
// 	fn(`}` + "\n\n")
//
// 	fn(`func (` + gstr.LowercaseFirstChar(interfaceImplName) + ` *` + interfaceImplName + `GEN) Count(ctx context.Context, ` +
// 		gstr.ToLowercase(structName) + ` *` + callerPackageName + "." + structName + `) (total int64, err error) {` + "\n")
// 	fn(`	return` + "\n")
// 	fn(`}` + "\n\n")
//
// 	fn(`func (` + gstr.LowercaseFirstChar(interfaceImplName) + ` *` + interfaceImplName + `GEN) WithTx(j ` +
// 		path.Base(thisCaller.PackageName) + "." + "Tx" + lib + `) ` + interfaceRepoPackageName + "." + interfaceName + ` {` + "\n")
// 	fn(`	return &` + param.WithInterfaceImplRepo.ImplInterfaceName + `GEN{` + "\n")
// 	fn(`		Tx` + lib + ": " + "j," + "\n")
// 	fn(`	}` + "\n")
// 	fn(`}` + "\n\n")
//
// 	fileLocation := param.WithInterfaceImplRepo.ImplLocationDir + "/" + interfaceImplName + "_GEN.go"
// 	err = os.MkdirAll(param.WithInterfaceImplRepo.ImplLocationDir, os.ModePerm)
// 	gcommon.PanicIfError(err)
//
// 	err = os.WriteFile(fileLocation, buf.Bytes(), os.ModePerm)
// 	gcommon.PanicIfError(err)
// }

// fn("// FieldAndValue is  function for get named arg for write query\n")
// fn(`func (` + gstr.LowercaseFirstChar(t.Name()) + ` *` + t.Name() + ") FieldAndValue() gmap.StrAny {\n")
// fn(`	sa := make(gmap.StrAny)` + "\n")
// fn(`	for _, field := range ` + gstr.LowercaseFirstChar(t.Name()) + `.RQField {` + "\n")
// fn(`		switch field {` + "\n")
// for k, _ := range field {
// 	fn(`		case ` + gstr.LowercaseFirstChar(t.Name()) + ".Field" + k + "():\n")
// 	fn(`			sa[field] = ` + gstr.LowercaseFirstChar(t.Name()) + "." + k + "\n")
// }
// fn(`		}` + "\n")
// fn("	}\n")
// fn(`	return sa` + "\n")
// fn("}\n\n")

// fn("// FieldArgForUpdate is function get string to SET update\n")
// fn(`func (` + gstr.LowercaseFirstChar(t.Name()) + ` *` + t.Name() + ") FieldArgForUpdate(prefix gsql.PrefixNamedArgPG) string {\n")
// fn(`	str := ""` + "\n")
// fn(`	columns := ` + gstr.LowercaseFirstChar(t.Name()) + ".FieldAndValue()\n")
// fn(`	i := 1` + "\n")
// fn(`	for k, _ := range columns {` + "\n")
// fn(`		if i == len(columns) {` + "\n")
// fn(`			str += k + " = " + string(prefix) + k` + "\n")
// fn(`		` + "} else {\n")
// fn(`			str += k + " = " + string(prefix) + k + ", "` + "\n")
// fn(`		` + "}\n")
// fn(`		i++` + "\n")
// fn(`	` + "}\n")
// fn(`	return str` + "\n")
// fn("}\n\n")
