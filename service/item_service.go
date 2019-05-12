package service

import (
	"fmt"
	"ijahstore/dao/sqlite"
	"ijahstore/libraries/exception"
)

type ItemService interface {
	AddStockItem(item *sqlite.StockItem) error
	UpdateStockItem(item *sqlite.StockItem) error
	GetStockItem(id uint) (sqlite.StockItem, error)
	GetAllStockItem() ([]sqlite.StockItem, error)
	GenerateSKUID() uint
}

type itemService struct {

}

func NewItemService() ItemService  {
	return itemService{}
}

func (itemService) AddStockItem(item *sqlite.StockItem) (err error) {
	var stockItem sqlite.StockItem
	databaseService.db.First(&stockItem, "name = ?", item.Name)

	if stockItem.Name != item.Name {
		databaseService.db.Create(item)

		databaseService.db.First(&stockItem, "name = ?", item.Name)

		current := &sqlite.CurrentStockItem{
			CurrentStock: 0,
			StockItem: stockItem,
		}

		databaseService.db.Create(current)

		if err := databaseService.db.GetErrors(); len(err) > 0 {
			return err[0]
		}

		return
	}

	return fmt.Errorf("%s", "duplicate entry")
}

func (itemService) UpdateStockItem(item *sqlite.StockItem) (err error) {
	var stockItem sqlite.StockItem

	databaseService.db.Model(sqlite.StockItem{}).Find(&stockItem, "id = ?", item.ID)

	if stockItem.ID == item.ID {
		stockItem.Colour = item.Colour
		stockItem.Size = item.Size
		stockItem.Name = item.Name
		item.SKUID = stockItem.SKUID
		item.SKU = stockItem.SKU
		databaseService.db.Model(&stockItem).Updates(stockItem)

		if err := databaseService.db.GetErrors(); len(err) > 0 {
			return err[0]
		}

		return
	} else {
		return exception.NewNotFoundException()
	}
}

func (itemService) GetStockItem(id uint) (sqlite.StockItem, error) {
	var stockItem sqlite.StockItem

	databaseService.db.Model(sqlite.StockItem{}).Find(&stockItem, "id = ?", id)

	if err := databaseService.db.GetErrors(); len(err) > 0 {
		return stockItem, err[0]
	}

	if stockItem.ID != id {
		return stockItem, exception.NewNotFoundException()
	}

	return stockItem, nil
}

func (itemService) GetAllStockItem() ([]sqlite.StockItem, error) {
	var stockItem []sqlite.StockItem

	databaseService.db.Find(&stockItem)

	if err := databaseService.db.GetErrors(); len(err) > 0 {
		return stockItem, err[0]
	}

	return stockItem, nil
}

func (itemService) GenerateSKUID() uint {
	var stockItem sqlite.StockItem

	row, _ := databaseService.db.Model(sqlite.StockItem{}).Select("MAX(sk_uid) as sk_uid").Find(&stockItem).Rows()

	_ = databaseService.db.ScanRows(row, &stockItem)

	return stockItem.SKUID + 1
}