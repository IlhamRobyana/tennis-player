package pg_storage

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var pgDB *gorm.DB

// GetPGClient is ...
func GetPGClient() (client *gorm.DB, err error) {
	err = godotenv.Load()
	if err != nil {
		panic(fmt.Sprintf("can't load .env : %v", err))
	}

	pgHost := os.Getenv("PG_HOST")
	pgPort := os.Getenv("PG_PORT")
	pgUsername := os.Getenv("PG_USERNAME")
	pgPassword := os.Getenv("PG_PASSWORD")
	pgDBName := os.Getenv("PG_DB_NAME")

	if pgDB == nil {
		pgDB, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", pgHost, pgPort, pgUsername, pgDBName, pgPassword))
		if err != nil {
			panic(fmt.Sprintf("failed to connect to database: %v", err))
		}
		client = pgDB
	}
	return
}
