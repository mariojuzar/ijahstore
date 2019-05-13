package exception

type stockNotFoundException struct {

}

func (stockNotFoundException) Error() string {
	return "Stock Data Not Found"
}

func NewStockNotFoundException() error {
	return stockNotFoundException{}
}

type entryItemLogNotFoundException struct {

}

func (entryItemLogNotFoundException) Error() string  {
	return "Entry Stock Item Log Not Found"
}

func NewEntryItemLogNotFoundException() error  {
	return entryItemLogNotFoundException{}
}

type salesOrderItemNotFound struct {

}

func (salesOrderItemNotFound) Error() string {
	return "Sales Order Item Not Found"
}

func NewSalesOrderNotFoundException() error  {
	return salesOrderItemNotFound{}
}

type notMatchOrderIdWithStockItemException struct {

}

func (notMatchOrderIdWithStockItemException) Error() string {
	return "Order Id Not Match With Sale Stock Item"
}

func NewNotMatchOrderIdWithStockItemException() error {
	return notMatchOrderIdWithStockItemException{}
}