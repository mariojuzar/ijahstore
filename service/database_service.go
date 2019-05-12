package service

import (
	"github.com/jinzhu/gorm"
	"ijahstore/dao/sqlite"
	"log"
)

type DatabaseService interface {

}

func init()  {
	db, err := gorm.Open("sqlite3", "./db/ijah_store.db")
	if err != nil {
		log.Fatal("Failed to init db:", err)
	}
	db.LogMode(true)

	db.AutoMigrate(
		&sqlite.Item{})
}