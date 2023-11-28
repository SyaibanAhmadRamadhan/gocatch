package example

import (
	"testing"

	"github.com/SyaibanAhmadRamadhan/jolly/Jdb/JOpg"
)

func TestGenerateStruct(t *testing.T) {
	userModel := JOpg.GeneratorModelForStructParam{
		Src:      &User{},
		FileName: "UserModel",
		SpecifiationTable: JOpg.SpecifiationTable{
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

	JOpg.GeneratorModelForStruct(userModel)
}
