package config

import (
	"log"

	"github.com/a5512167086/url-shortener/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	cfg := LoadDBConfig()

	log.Println("Connecting to DB with DSN:", cfg.DSN())

	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&model.URLShortener{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
