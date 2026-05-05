package models

type TicketBuyer struct {
	Tickets int
	Index   int
}

func NewTicketBuyer(index int, tickets int) TicketBuyer {
	return TicketBuyer{
		Tickets: tickets,
		Index:   index,
	}
}

func (t *TicketBuyer) UpdateTicket() {
	t.Tickets--
}
