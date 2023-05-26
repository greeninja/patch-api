package db

import (
	"log"

	"/home/ncampion/git/github/greeninja/patch-api/pkg/models"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	db, err = gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&models.Patches{})
	return db
}
