package config

import (
	"log"

	"github.com/alandev/ssl-monitoring/types"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("ssl_monitoring.sqlite"), &gorm.Config{})
	if err != nil {
		log.Fatal("can't connect to db", err)
	}

	err = db.AutoMigrate(&types.Certificate{})
	if err != nil {
		log.Fatal("can't migrate", err)
	}
	return db, nil
}
