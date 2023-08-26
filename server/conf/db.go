package conf

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"server/model"
)

type DBContact interface {
	ConnectDB() *gorm.DB
	GetDB() *gorm.DB
}

var DB *gorm.DB

func (a *App) GetDB() *gorm.DB {
	if DB == nil {
		return a.ConnectDB()
	}
	return DB
}

func (a *App) ConnectDB() *gorm.DB {
	var err error
	DB, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("failed to connect database")
	}

	// Migrate the schema
	err = DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("failed to migrate schema")
	}

	return DB
}
