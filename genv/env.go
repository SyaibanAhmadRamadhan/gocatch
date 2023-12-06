package genv

import (
	"os"
)

type EnvLib string

const DefaultEnvLib EnvLib = "godotenv"

// Initialize Function Method for EnvLib (godotenv, viper, etc)
// if envOverride is true, it will load .env.override file
// if envOverride is false, it will not load .env.override file
// by default godotenv
func Initialize(envLib EnvLib, envOverride bool) error {
	switch envLib {
	default:
		return godotEnv(envOverride)
	}
}

func GetEnv(key string, defaultValue ...string) string {
	v, ok := os.LookupEnv(key)
	if !ok {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return ""
	}

	return v
}
