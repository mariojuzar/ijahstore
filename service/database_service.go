package service

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"ijahstore/dao/sqlite"
	"log"
)

type DatabaseService struct {
	db 				*gorm.DB
	IsInitialized 	bool
}

var databaseService DatabaseService

func init()  {
	db, err := gorm.Open("sqlite3", "./db/ijah_store.db")
	if err != nil {
		log.Fatal("Failed to init db: ", err)
	}
	db.LogMode(true)

	databaseService = DatabaseService{db:db, IsInitialized:true}

	db.AutoMigrate(
		&sqlite.StockItem{},
		&sqlite.EntryStockLog{},
		&sqlite.OutcomeStockLog{},
		&sqlite.CurrentStockItem{},
		&sqlite.SaleStock{},
		&sqlite.StockValue{},
		&sqlite.SaleOrder{},
		&sqlite.SaleReport{},
		&sqlite.ValueReport{})
}