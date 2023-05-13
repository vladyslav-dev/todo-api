package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadENV() error {
	goEnv := os.Getenv("GO_ENV")
	fmt.Println(goEnv)
	if goEnv == "" || goEnv == "development" {
		err := godotenv.Load()
		if err != nil {
			return err
		}
	}

	return nil
}
