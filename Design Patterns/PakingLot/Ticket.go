package main

type Ticket struct {
	ID            string
	ParkingRecord *ParkingRecord
}

func NewTicket(id string, record *ParkingRecord) *Ticket {
	return &Ticket{
		ID:            id,
		ParkingRecord: record,
	}
}
func (t *Ticket) GetID() string {
	return t.ID
}
