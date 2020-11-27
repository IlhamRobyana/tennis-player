package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	"github.com/ilhamrobyana/tennis-player/entity"
	"github.com/ilhamrobyana/tennis-player/pg_storage"
)

func main() {
	client, err := pg_storage.GetPGClient()
	if err != nil {
		fmt.Println(err)
	}
	migrateScheme(client)
}

func migrateScheme(DB *gorm.DB) {
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprintf("can't load .env : %v", err))
	}

	isTableDropped := os.Getenv("DROP_TABLE")
	if isTableDropped == "true" {
		DB.DropTableIfExists(
			&entity.Player{},
		)
	}

	DB.AutoMigrate(
		&entity.Player{},
	)
}
