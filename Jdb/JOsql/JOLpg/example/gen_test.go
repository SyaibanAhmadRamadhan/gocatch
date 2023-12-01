package example

import (
	"testing"

	"github.com/SyaibanAhmadRamadhan/jolly/Jdb/JOsql"
)

func TestGenerateStruct(t *testing.T) {
	userModel := JOsql.GeneratorModelForStructParam{
		Src:      &User{},
		FileName: "UserModel",
		SpecifiationTable: JOsql.SpecifiationTable{
			TableName:  "user",
			SchemaName: "public",
		},
		// WithInterfaceRepo: JOpg.WithInterfaceRepo{
		// 	Make:          true,
		// 	Dir:           "./repository",
		// 	InterfaceName: "UserRepo",
		// },
		// WithInterfaceImplRepo: JOpg.WithInterfaceImplRepo{
		// 	Make:              true,
		// 	ImplLocationDir:   "./repositoryimpl",
		// 	ImplInterfaceName: "userRepoImpl",
		// },
	}

	JOsql.GeneratorModelForStruct(userModel)
}
