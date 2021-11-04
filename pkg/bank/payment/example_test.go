package payment

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleMax() {
	payments := []types.Payment{
		{
			Amount:0,
			ID: 0,
		},
		{
			Amount:2000,
			ID: 1,
		},
		{
			Amount:2000,
			ID: 2,
		},
		{
			Amount:2000,
			ID: 3,
		},
	}
	max:=Max(payments)
	fmt.Println(max.ID)

	//Output: 1
}


func ExamplePaymentSources(){
	cards := []types.Card{
		{
			PAN: "5058 xxxx xxxx 8881",
			Balance: 0,
			Active: true,
		},
		{
			PAN: "5058 xxxx xxxx 8882",
			Balance: 2000,
			Active: true,
		},
		{
			PAN: "5058 xxxx xxxx 8883",
			Balance: -1000,
			Active: true,
		},
		{
			PAN: "5058 xxxx xxxx 8884",
			Balance: 3000,
			Active: false,
		},
		{
			PAN: "5058 xxxx xxxx 8885",
			Balance: 3000,
			Active: true,
		},
	}
	numbers := PaymentSources(cards)
	for _, number := range numbers {
		fmt.Println(number.Number)
	}
	//Output:
	//5058 xxxx xxxx 8882
	//5058 xxxx xxxx 8885
}