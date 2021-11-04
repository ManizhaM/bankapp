package types

//Money - денежная сумма в мин.еденицах (центы, копейки, дирамы и тд)
type Money int64

// Currency - код валюты
type Currency string

// Коды валют
const (
	TJS Currency = "TJS"
	USD Currency = "USD"
	RUB Currency = "RUB"
)

// PAN - номер карты
type PAN string

//Card - информация о платежной карте
type Card struct{
	ID 			int
	PAN			PAN
	Balance 	Money
	MinBalance	Money
	Currency 	Currency
	Color 		string
	Name		string
	Active		bool
}

//Payment information about payment
type Payment struct{
	ID 		int
	Amount	Money
}