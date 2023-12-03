package genv

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
