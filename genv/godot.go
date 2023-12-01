package genv

import (
	"errors"

	"github.com/joho/godotenv"
)

/*
godotEnv loads environment variables from '.env' files using godotenv.
It searches multiple directories for the '.env' file. If 'envOverride' is true, it also loads '.env.override'.
On success, it returns nil. On failure, it returns an error.
*/
func godotEnv(envOverride bool) error {
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
