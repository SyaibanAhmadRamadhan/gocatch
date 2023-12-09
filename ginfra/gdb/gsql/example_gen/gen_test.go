package example_gen

import (
	"testing"

	"github.com/SyaibanAhmadRamadhan/gocatch/ginfra/gdb"
)

func TestGenerateStruct(t *testing.T) {
	userModel := gdb.GeneratorModelForStructParam{
		Src:      &User{},
		FileName: "UserModel",
		SpecifiationTable: gdb.SpecifiationTable{
			TableName:  "user",
			SchemaName: "public",
		},
	}

	gdb.GeneratorModelFromStruct(userModel)
}
