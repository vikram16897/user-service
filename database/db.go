package database

import (
	"time"

	"github.com/user-service/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(dsn string) (*gorm.DB, error) {
	time.Sleep(60 * time.Second)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&model.User{}); err != nil {
		return nil, err
	}

	return db, nil
}
