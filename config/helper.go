package config

import "github.com/joho/godotenv"
import "log"
import "os"

func DotEnvVariable(key string) string {
  
	if err := godotenv.Load(".env"); err != nil {
	  log.Fatalf("Error loading .env file")
	}
  
	return os.Getenv(key)
}