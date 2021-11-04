package card

import (
	"bank/pkg/bank/types"
)

func IssueCard(currency types.Currency, color string, name string) types.Card {
	card := types.Card{
		ID:       1000,
		PAN:      "5058 xxxx xxxx 0001",
		Balance:  0,
		Currency: currency,
		Color:    color,
		Name:     name,
		Active:   true,
	}
	return card
}

func Withdraw(card *types.Card, amount types.Money) {

	const WithdrawLimit = 20_000_00
	if !card.Active {
		return
	}
	if amount > card.Balance {
		return
	}
	if amount < 0 {
		return
	}
	if amount > WithdrawLimit {
		return
	}
	card.Balance -= amount
}

func Deposit(card *types.Card, amount types.Money) {
	const DepositLimit = 5000000 //лимит зачисления средств
	if !card.Active {
		return
	}
	if amount > DepositLimit {
		return
	}
	if amount <= 0 {
		return
	}
	card.Balance += amount
}

func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int) {
	if (!card.Active) {
		return
	}
	if (card.Balance < 0) {
		return
	}
	var maxBonus types.Money = 500000
	bonus := card.MinBalance * types.Money(percent)/100 * types.Money(daysInMonth) / types.Money(daysInYear)
	if bonus > maxBonus {
		bonus = 0
	}
	if bonus < 0 {
		return
	}
	card.Balance += types.Money(int(bonus))
}

//Function count sum on card balances 
func Total(cards []types.Card)types.Money{
	var sum types.Money = 0;
	for _, card := range cards {
		if(card.Balance>=0 && card.Active){
			sum += card.Balance;
		}
	}
	return sum
}



func PaymentSources(cards []types.Card) []types.PaymentSource{
	var PaymentSource []types.PaymentSource
	for index := range cards {
		if(cards[index].Active && cards[index].Balance>0){
		PaymentSource = append(PaymentSource, types.PaymentSource{
			Type: "card", 
			Number: cards[index].PAN, 
			Balance: cards[index].Balance,
			}) 
		}
	}
	return PaymentSource
}