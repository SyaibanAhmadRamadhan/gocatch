package Jenv

type EnvLib string

const GodotEnv EnvLib = "godotenv"

// Init Function Method for EnvLib (godotenv, viper, etc)
// if envOverride is true, it will load .env.override file
// if envOverride is false, it will not load .env.override file
// by default godotenv
func Init(envLib EnvLib, envOverride bool) error {
	switch envLib {
	case GodotEnv:
		return InitGodotEnv(envOverride)
	default:
		return InitGodotEnv(envOverride)
	}
}
