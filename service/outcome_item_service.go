package service

import (
	jsoniter "github.com/json-iterator/go"
	"ijahstore/dao/sqlite"
	"ijahstore/entity/request"
	"ijahstore/libraries/exception"
	"time"
)

type OutcomeItemService interface {
	AddOutcomeItem(request request.OutComeItemCreationRequest) (sqlite.OutcomeStockLog, error)
	GetAllOutcomeItem() ([]sqlite.OutcomeStockLog, error)
}

type outcomeItemService struct {

}

func NewOutcomeItemService() OutcomeItemService  {
	return outcomeItemService{}
}

func (outcomeItemService) AddOutcomeItem(request request.OutComeItemCreationRequest) (sqlite.OutcomeStockLog, error) {
	var salesOrder sqlite.SaleOrder
	var out sqlite.OutcomeStockLog

	databaseService.db.Model(salesOrder).Find(&salesOrder, "sale_order_id =?", request.SaleOrderId)

	if salesOrder.SaleOrderId == request.SaleOrderId {
		var saleStock sqlite.SaleStock
		var stockInSale []sqlite.SaleStock
		var json = jsoniter.ConfigCompatibleWithStandardLibrary

		_ = json.UnmarshalFromString(salesOrder.SaleStockString, &stockInSale)
		saleStock = getSaleStockItem(request.StockId, stockInSale)

		if saleStock.ItemID == request.StockId {
			out = sqlite.OutcomeStockLog{
				OutcomeId:		generateId(),
				Time: 			time.Now(),
				StockItem: 		saleStock.StockItem,
				OrderId:		salesOrder.OrderId,
				AmountDelivered:request.AmountDelivered,
				SellPrice: 		saleStock.SellPrice,
				TotalPrice: 	saleStock.SellPrice * request.AmountDelivered,
				Note:			generateNote(salesOrder.OrderId),
			}

			databaseService.db.Create(out)

			var current sqlite.CurrentStockItem
			databaseService.db.Model(current).Find(&current, "item_id =?", saleStock.ItemID)

			current.CurrentStock -= request.AmountDelivered

			databaseService.db.Model(current).Where("item_id =?", current.ItemID).Updates(current)

			if err := databaseService.db.GetErrors(); len(err) > 0 {
				return out, err[0]
			}

			return out, nil
		} else {
			return out, exception.NewNotMatchOrderIdWithStockItemException()
		}
	} else {
		return out, exception.NewSalesOrderNotFoundException()
	}
}

func (outcomeItemService) GetAllOutcomeItem() ([]sqlite.OutcomeStockLog, error) {
	var outs []sqlite.OutcomeStockLog

	databaseService.db.Find(&outs)

	if err := databaseService.db.GetErrors(); len(err) > 0 {
		return outs, err[0]
	}

	return outs, nil
}

func generateId() uint  {
	var out sqlite.OutcomeStockLog

	row, _ := databaseService.db.Model(out).Select("MAX(outcome_id) as outcome_id").Find(&out).Rows()

	_ = databaseService.db.ScanRows(row, &out)

	return out.OutcomeId + 1
}

func generateNote(id string) string  {
	return "Pesanan dari " + id
}

func getSaleStockItem(id uint, sales []sqlite.SaleStock) sqlite.SaleStock  {
	var ss sqlite.SaleStock
	for _, sale := range sales {
		if sale.ItemID == id {
			return sale
		}
	}
	return ss
}