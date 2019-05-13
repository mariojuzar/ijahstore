package service

import "ijahstore/dao/sqlite"

type StockService interface {
	GetAllCurrentStock() ([]sqlite.CurrentStockItem, error)
}

type stockService struct {

}

func (stockService) GetAllCurrentStock() ([]sqlite.CurrentStockItem, error) {
	var current []sqlite.CurrentStockItem

	databaseService.db.Find(&current)

	if err := databaseService.db.GetErrors(); len(err) > 0 {
		return current, err[0]
	}

	return current, nil
}

func NewStockService() StockService  {
	return stockService{}
}