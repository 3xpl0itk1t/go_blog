package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var (
	PORT      string = os.Getenv("PORT")
	MONGO_URL string = os.Getenv("MONGO_URL")
	SECRET    string = os.Getenv("SECRET")
)
