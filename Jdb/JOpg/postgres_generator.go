package JOpg

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strings"

	"github.com/SyaibanAhmadRamadhan/jolly"
	"github.com/SyaibanAhmadRamadhan/jolly/Jlog"
	"github.com/SyaibanAhmadRamadhan/jolly/Jtype/Jstr"
	"github.com/SyaibanAhmadRamadhan/jolly/Jtype/Jstruct"
)

// GeneratorModelForStructParam defines the parameters
// for generating methods for a struct
type GeneratorModelForStructParam struct {
	Src      any    // Src is the struct instance
	FileName string // FileName is the name of generated source file without extension
}

// GeneratorModelForStruct generates methods for given structs.
// This function expects each struct to contain a field `QueryColumnFields []string`.
// The generated methods manipulate or make use of this field for various operations.
// The function does not support nested structs with name tag.
// if you want nested struct for this function, you can use "-" tag in your struct.
// for example check in folder Jdb/JOpg/exGen
func GeneratorModelForStruct(params ...GeneratorModelForStructParam) {
	for _, param := range params {
		buf := bytes.Buffer{}
		structOrm := func(str string) {
			_, err := buf.WriteString(str)
			jolly.PanicIF(err)
		}

		t := reflect.TypeOf(param.Src).Elem()

		packageName := Jstr.LastStringOfSubStr(Jlog.CallerInfo(2).PackageName, "/")
		fmt.Println(Jlog.CallerInfo(1))
		structOrm(`package ` + packageName + "\n\n")

		res, _ := os.ReadFile(param.FileName + ".go")
		r, _ := regexp.Compile(`import \(([^)]+)\)`)
		importSrc := r.FindString(string(res))
		importSrc = strings.Replace(importSrc, "(", "(\n\t\"errors\"\n\t\"fmt\"\n\t\"strings\"\n", 1)

		structOrm(importSrc + "\n\n")

		structOrm("// DO NOT EDIT, will be overwritten by https://github.com/SyaibanAhmadRamadhan/jolly/blob/main/Jdb/JOpg/postgres_generator.go. \n\n")

		structOrm("// New" + t.Name() + " is a struct with pointer that represents the table " + t.Name() + " in the database.\n")
		structOrm(`func New` + t.Name() + `() *` + t.Name() + " {\n")
		structOrm(`	return &` + t.Name() + "{}\n")
		structOrm("}\n\n")

		structOrm("// New" + t.Name() + "WithOutPtr is a struct without pointer that represents the table " + t.Name() + " in the database.\n")
		structOrm(`func New` + t.Name() + `WithOutPtr() ` + t.Name() + " {\n")
		structOrm(`	return ` + t.Name() + "{}\n")
		structOrm("}\n\n")

		field := Jstruct.GetTagAndFieldNameFromStruct(param.Src, "", "db")
		for k, v := range field {
			structOrm("// Field" + k + " is a field or column in the table " + t.Name() + ".\n")
			structOrm(`func (` + Jstr.FirstCharToLower(t.Name()) + ` *` + t.Name() + ") Field" + k + "() string {\n")
			structOrm(`	return "` + strings.Split(v, "|")[0] + "\"\n")
			structOrm("}\n\n")

			structOrm("// Set" + k + " is a setter for the field or column " + k + " in the table " + t.Name() + ".\n")
			structOrm(`func (` + Jstr.FirstCharToLower(t.Name()) + ` *` + t.Name() + ") Set" + k + "(param " + strings.Split(v, "|")[1] + ") {\n")
			structOrm(`	` + Jstr.FirstCharToLower(t.Name()) + "." + k + " = param" + "\n")
			structOrm("}\n\n")
		}

		structOrm("// AllField is a function to get all field or column in the table " + t.Name() + ".\n")
		structOrm(`func (` + Jstr.FirstCharToLower(t.Name()) + ` *` + t.Name() + ") AllField() (str string) {\n")
		structOrm(`	str += `)
		var arrField []string
		for _, v := range field {
			arrField = append(arrField, strings.Split(v, "|")[0])
		}
		structOrm("`\n\t\t" + strings.Join(arrField, ", \n\t\t") + "`" + "\n")
		structOrm(`	return` + "\n")
		structOrm("}\n\n")

		structOrm("// Scan is a function to scan the value with for rows.Value() from the database to the struct " + t.Name() + ".\n")
		structOrm(`func (` + Jstr.FirstCharToLower(t.Name()) + ` *` + t.Name() + ") Scan(value any, columns ...string) (err error) {\n")
		structOrm(`	for _, column := range columns {` + "\n")
		structOrm(`		switch column {` + "\n")
		for k, v := range field {
			typeStruct := strings.Split(v, "|")[1]
			structOrm(`		case ` + Jstr.FirstCharToLower(t.Name()) + `.Field` + k + `():` + "\n")
			structOrm(`			val, ok := value.(` + typeStruct + ")\n")
			structOrm(`			` + "if !ok {\n")
			structOrm(`				` + "return errors.New(\"invalid type\")\n")
			structOrm(`			` + "}\n")
			structOrm(`			` + Jstr.FirstCharToLower(t.Name()) + `.Set` + k + "(val)\n")
			structOrm(`			return nil` + "\n")
		}
		structOrm(`		}` + "\n")
		structOrm(`	}` + "\n")
		structOrm(`	` + "return errors.New(\"invalid column\")\n")
		structOrm("}\n\n")

		structOrm("// SetColumn is a function to set the column to QueryColumnFields for that will be used in the query.\n")
		structOrm(`func (` + Jstr.FirstCharToLower(t.Name()) + ` *` + t.Name() + ") SetColumn(columns ...string) (err error) {\n")
		structOrm(`	for _, column := range columns {` + "\n")
		structOrm(`		switch column {` + "\n")
		for k := range field {
			structOrm(`		case ` + Jstr.FirstCharToLower(t.Name()) + `.Field` + k + `():` + "\n")
		}
		structOrm(`		default:` + "\n")
		structOrm(`			` + "return errors.New(\"invalid column\")\n")
		structOrm(`		}` + "\n")
		structOrm(`	}` + "\n")
		structOrm(`	` + Jstr.FirstCharToLower(t.Name()) + ".QueryColumnFields = append(" +
			Jstr.FirstCharToLower(t.Name()) + ".QueryColumnFields, columns...)" + "\n")
		structOrm(`	` + "return nil\n")
		structOrm("}\n\n")

		structOrm("// DeleteColumnFromQueryColumnFields is a function to delete the column from QueryColumnFields for that will be used in the query.\n")
		structOrm(`func (` + Jstr.FirstCharToLower(t.Name()) + ` *` + t.Name() + ") DeleteColumnFromQueryColumnFields(elems ...string) (err error) {\n")
		structOrm(`	var colums []string` + "\n")
		structOrm(`	for _, v := range ` + Jstr.FirstCharToLower(t.Name()) + ".QueryColumnFields" + `{` + "\n")
		structOrm(`		colums = append(colums, v)` + "\n")
		structOrm(`	}` + "\n\n")
		structOrm(`	for _, elem := range elems {` + "\n")
		structOrm(`		index := -1` + "\n")
		structOrm(`		for i, column := range ` + Jstr.FirstCharToLower(t.Name()) + ".QueryColumnFields" + `{` + "\n")
		structOrm(`			if column == elem {` + "\n")
		structOrm(`				index = i` + "\n")
		structOrm(`				break` + "\n")
		structOrm(`			}` + "\n")
		structOrm(`		}` + "\n")
		structOrm(`		if index == -1 {` + "\n")
		structOrm(`			return fmt.Errorf("column %s not found", elem)` + "\n")
		structOrm(`		}` + "\n")
		structOrm(`		colums = append(colums[:index], colums[index+1:]...)` + "\n")
		structOrm(`	}` + "\n")
		structOrm(`	` + Jstr.FirstCharToLower(t.Name()) + ".QueryColumnFields" + ` = colums` + "\n")
		structOrm(`	return nil` + "\n")
		structOrm("}\n\n")

		structOrm("// QueryColumnFieldToStrings is a function to get the column format string from QueryColumnFields for that will be used in the query.\n")
		structOrm(`func (` + Jstr.FirstCharToLower(t.Name()) + ` *` + t.Name() + ") QueryColumnFieldToStrings() (columnStr string) {\n")
		structOrm(`	return strings.Join(` + Jstr.FirstCharToLower(t.Name()) + ".QueryColumnFields, \", \")\n")
		structOrm("}\n\n")

		structOrm("// ResetQueryColumnFields is a function to reset QueryColumnFields.\n")
		structOrm(`func (` + Jstr.FirstCharToLower(t.Name()) + ` *` + t.Name() + ") ResetQueryColumnFields() {\n")
		structOrm(`	` + Jstr.FirstCharToLower(t.Name()) + ".QueryColumnFields = []string{}\n")
		structOrm("}\n\n")

		err := os.WriteFile(param.FileName+"_GEN.go", buf.Bytes(), os.ModePerm)
		jolly.PanicIF(err)
	}
}
