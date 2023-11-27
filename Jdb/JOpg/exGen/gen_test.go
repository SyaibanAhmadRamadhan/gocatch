package exGen

import (
	"testing"

	"github.com/SyaibanAhmadRamadhan/jolly/Jdb/JOpg"
)

func TestGenerateStruct(t *testing.T) {
	userModel := JOpg.GeneratorModelForStructParam{
		Src:      &User{},
		FileName: "UserModel",
	}
	// addressModel := JOpg.GeneratorModelForStructParam{
	// 	Src:      &Address{},
	// 	FileName: "Address",
	// }

	JOpg.GeneratorModelForStruct(userModel)
}
