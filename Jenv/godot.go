package Jenv

import (
	"errors"

	"github.com/joho/godotenv"
)

func InitGodotEnv(envOverride bool) error {
	dir := []string{"./", "../", "../../", "../../../", "../../../../", "../../../../../"}
	for _, v := range dir {
		err := godotenv.Overload(v + ".env")
		if err == nil {
			if envOverride {
				envOverrideFile := v + ".env.override"
				err = godotenv.Overload(envOverrideFile)
				if err != nil {
					return err
				}
			}

			return nil
		}
	}

	return errors.New("cannot load environtment file")
}
