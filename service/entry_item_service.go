package service

import (
	"ijahstore/dao/sqlite"
	"ijahstore/entity/request"
	"ijahstore/libraries/exception"
	"time"
)

type EntryItemService interface {
	AddEntryItem(entry *request.EntryItemCreationRequest) (sqlite.EntryStockLog, error)
	UpdateEntryItem(entry *request.EntryItemUpdateRequest) (sqlite.EntryStockLog, error)
	GetAllEntryItem() ([]sqlite.EntryStockLog, error)
}

type entryItemService struct {

}

func NewEntryItemService() EntryItemService  {
	return entryItemService{}
}

func (entryItemService) AddEntryItem(entry *request.EntryItemCreationRequest) (sqlite.EntryStockLog, error) {
	var current sqlite.CurrentStockItem
	var temp sqlite.EntryStockLog

	databaseService.db.Model(current).First(&current, "item_id = ?", entry.StockId)

	if current.ItemID == entry.StockId {
		var stockItem sqlite.StockItem
		databaseService.db.Model(stockItem).First(&stockItem, "item_id = ?", current.ItemID)

		saveEntry := &sqlite.EntryStockLog{
			EntryId: 		generateID(),
			Time: 			time.Now(),
			AmountOrder: 	entry.AmountOrder,
			AmountReceived: entry.AmountReceived,
			PurchasePrice: 	entry.PurchasePrice,
			TotalPrice: 	entry.AmountOrder * entry.PurchasePrice,
			ReceiptNumber: 	entry.ReceiptNumber,
			Note: 			entry.Note,
			StockItem:		stockItem,
		}

		databaseService.db.Create(saveEntry)

		var savedEntry sqlite.EntryStockLog
		databaseService.db.Model(savedEntry).First(&savedEntry, "entry_id = ?", saveEntry.EntryId)

		newCurrent := &sqlite.CurrentStockItem{
			StockItem: 		current.StockItem,
			CurrentStock:	savedEntry.AmountReceived,
		}

		databaseService.db.Model(&newCurrent).Where("item_id =?", newCurrent.ItemID).Updates(newCurrent)

		if err := databaseService.db.GetErrors(); len(err) > 0 {
			return savedEntry, err[0]
		}

		return savedEntry, nil
	} else {
		return temp, exception.NewStockNotFoundException()
	}
}

func (entryItemService) UpdateEntryItem(entry *request.EntryItemUpdateRequest) (sqlite.EntryStockLog, error) {
	var oldEntry sqlite.EntryStockLog
	var oldEntry2 sqlite.EntryStockLog
	var temp sqlite.EntryStockLog

	databaseService.db.Model(oldEntry).First(&oldEntry, "entry_id = ?", entry.ID)
	databaseService.db.Model(oldEntry2).First(&oldEntry2, "entry_id = ?", entry.ID)

	if oldEntry.EntryId == entry.ID {
		var current sqlite.CurrentStockItem
		oldAmount := oldEntry2.AmountReceived

		databaseService.db.Model(current).First(&current, "item_id = ?", oldEntry.ItemID)

		oldEntry.AmountReceived = entry.AmountReceived
		oldEntry.Note = entry.Note

		databaseService.db.Model(&oldEntry).Where("entry_id = ?", entry.ID).Updates(oldEntry)

		var savedEntry sqlite.EntryStockLog
		databaseService.db.Model(savedEntry).First(&savedEntry, "entry_id = ?", entry.ID)

		newStock := current.CurrentStock + (entry.AmountReceived - oldAmount)

		newCurrent := &sqlite.CurrentStockItem{
			StockItem: 		current.StockItem,
			CurrentStock:	newStock,
		}

		databaseService.db.Model(&newCurrent).Where("item_id =?", newCurrent.ItemID).Updates(newCurrent)

		if err := databaseService.db.GetErrors(); len(err) > 0 {
			return oldEntry, err[0]
		}

		return oldEntry, nil
	} else {
		return temp, exception.NewEntryItemLogNotFoundException()
	}
}

func (entryItemService) GetAllEntryItem() ([]sqlite.EntryStockLog, error) {
	var entry []sqlite.EntryStockLog

	databaseService.db.Find(&entry)

	if err := databaseService.db.GetErrors(); len(err) > 0 {
		return entry, err[0]
	}

	return entry, nil
}

func generateID() uint  {
	var entry sqlite.EntryStockLog

	row, _ := databaseService.db.Model(sqlite.EntryStockLog{}).Select("MAX(entry_id) as entry_id").Find(&entry).Rows()

	_ = databaseService.db.ScanRows(row, &entry)

	return entry.EntryId + 1
}