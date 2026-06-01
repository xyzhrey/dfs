package config

import (
	"os"
)

func DBURL() string {
	return os.Getenv("DB_URL")
}
