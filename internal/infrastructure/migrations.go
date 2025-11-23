package infrastructure

import (
	"16/internal/infrastructure/migrations"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Migrate() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open("host=localhost password=password user=postgres dbname=postgres port=5432"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&migrations.User{})
	return db, nil
}
