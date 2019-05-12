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