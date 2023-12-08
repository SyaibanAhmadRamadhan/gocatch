package example

import (
	"testing"
)

func TestGenerateStruct(t *testing.T) {
	userModel := Gsql.GeneratorModelForStructParam{
		Src:      &User{},
		FileName: "UserModel",
		SpecifiationTable: Gsql.SpecifiationTable{
			TableName:  "user",
			SchemaName: "public",
		},
	}

	Gsql.GeneratorModelFromStruct(userModel)
}
