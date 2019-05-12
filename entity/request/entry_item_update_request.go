package request

type EntryItemUpdateRequest struct {
	ID 				uint		`json:"id"`
	AmountReceived	uint		`json:"amount_received"`
	Note 			string		`json:"note"`
}
