package Configs

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SQLDatabaseConnection(name string) (db *gorm.DB, err error) {
	db, err = gorm.Open(sqlite.Open(name), &gorm.Config{})
	return db, err
}
