package settings

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var Debug bool

func boolString(s string) bool {
	if s != "" && s != "false" && s != "FALSE" && s != "0" {
		return true
	}

	return false
}

func init() {
	_ = godotenv.Load()

	Debug = boolString(os.Getenv("DEBUG"))
}

func RequireEnvs(needEnvs []string) {
	for _, envKey := range needEnvs {
		if os.Getenv(envKey) == "" {
			log.Fatalf("env %s missed", envKey)
		}
	}
}
