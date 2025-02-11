package migrate

import (
	"03/model"
	"log"

	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {
	log.Printf("Migrating database")
	db.Exec("CREATE SCHEMA IF NOT EXISTS intern_test")

	err := db.AutoMigrate(
		&model.Dialog{},
		&model.Word{},
		&model.WordDialog{},
	)

	if err != nil {
		log.Fatalf("Error migrating: %v", err)
	}
}
