package service

import (
	"ijahstore/dao/sqlite"
	"ijahstore/entity/request"
	"math/rand"
	"strconv"
	"time"
)

type OrderService interface {
	AddOrder(order []request.OrderCreationRequest) (sqlite.SaleOrder, error)
}

type orderService struct {

}

func NewOrderService() OrderService  {
	return orderService{}
}

func (orderService) AddOrder(orders []request.OrderCreationRequest) (sqlite.SaleOrder, error) {
	var sales []sqlite.SaleStock

	for _, order := range orders {
		var stockItem sqlite.StockItem
		var entries []sqlite.EntryStockLog

		databaseService.db.First(&stockItem, "item_id =?", order.StockId)
		databaseService.db.Find(&entries, "item_id =?", order.StockId)

		stock := sqlite.SaleStock{
			SaleStockId:	generateIdSaleStock(),
			StockItem: 		stockItem,
			Quantity: 		order.Quantity,
			SellPrice: 		order.SellPrice,
			TotalPrice: 	order.SellPrice * order.Quantity,
			PurchasedPrice:	uint(getPurchasePrice(entries)),
			Profit: 		order.SellPrice - uint(getPurchasePrice(entries)),
		}

		sales = append(sales, stock)

	}

	saleOrder := sqlite.SaleOrder{
		SaleOrderId:	generateIdSaleOrder(),
		OrderId: 		generateOrderId(),
		Time: 			time.Now(),
		SaleStock:		sales,
	}

	databaseService.db.Create(&saleOrder)

	if err := databaseService.db.GetErrors(); len(err) > 0 {
		return saleOrder, err[0]
	}

	return saleOrder, nil
}

func getPurchasePrice(entries []sqlite.EntryStockLog) int {
	sumPrice := 0
	n := 0

	for _, entry := range entries {
		sumPrice += int(entry.PurchasePrice)
		n += 1
	}

	return sumPrice/n
}

func generateIdSaleOrder() uint  {
	var order sqlite.SaleOrder

	row, _ := databaseService.db.Model(order).Select("MAX(sale_order_id) as sale_order_id").Find(&order).Rows()

	_ = databaseService.db.ScanRows(row, &order)

	return order.SaleOrderId + 1
}

func generateIdSaleStock() uint  {
	var order sqlite.SaleStock

	row, _ := databaseService.db.Model(order).Select("MAX(sale_stock_id) as sale_stock_id").Find(&order).Rows()

	_ = databaseService.db.ScanRows(row, &order)

	return order.SaleStockId + 1
}

func generateOrderId() string {
	today := time.Now()
	return "ID-" + strconv.Itoa(today.Year()) + strconv.Itoa(int(today.Month())) + strconv.Itoa(today.Day()) + "-" + randomID()
}

func randomID() string {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	return strconv.Itoa(r1.Intn(999999))
}