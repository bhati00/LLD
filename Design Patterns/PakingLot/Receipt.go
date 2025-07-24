package main

type Receipt struct {
	ID            string
	TicketId      string
	Fare          float64
	PaymentStatus bool
}

func NewReceipt(id string, ticketId string, fare float64, paymentStatus bool) *Receipt {
	return &Receipt{
		ID:            id,
		TicketId:      ticketId,
		Fare:          fare,
		PaymentStatus: paymentStatus,
	}
}
