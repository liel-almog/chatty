package configs

import (
	"log"
	"os"
)

func GetEnv(env string) (envVal string) {

	envVal, isEnv := os.LookupEnv(env)
	if !isEnv {
		log.Fatal("Environment variable not found")
	}

	return
}
