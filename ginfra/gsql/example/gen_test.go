package example

import (
	"testing"

	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra/gsql"
)

func TestGenerateStruct(t *testing.T) {
	userModel := gsql.GeneratorModelForStructParam{
		Src:      &User{},
		FileName: "UserModel",
		SpecifiationTable: gsql.SpecifiationTable{
			TableName:  "user",
			SchemaName: "public",
		},
	}

	gsql.GeneratorModelFromStruct(userModel)
}
