package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleWithdraw_positive() {
	card := types.Card{Balance: 20_000_00, Active: true}
	Withdraw(&card, 10_000_00)
	fmt.Println(card.Balance)
	//Output: 1000000
}

func ExampleWithdraw_noMoney() {
	card := types.Card{Balance: 0, Active: true}
	Withdraw(&card, 10_000_00)
	fmt.Println(card.Balance)
	//Output: 0
}

func ExampleWithdraw_inactive() {
	card := types.Card{Balance: 20_000_00, Active: false}
	Withdraw(&card, 10_000_00)
	fmt.Println(card.Balance)
	//Output: 2000000
}

func ExampleWithdraw_limit() {
	card := types.Card{Balance: 20_000_00, Active: true}
	Withdraw(&card, 24_000_00)
	fmt.Println(card.Balance)
	//Output: 2000000
}

func ExampleDeposit_positive() {
	card := types.Card{Balance: 2000000, Active: true}
	Deposit(&card, 1000000)
	fmt.Println(card.Balance)
	//Output: 3000000
}

func ExampleDeposit_inactive() {
	card := types.Card{Balance: 2000000, Active: false}
	Deposit(&card, 20000)
	fmt.Println(card.Balance)
	//Output: 2000000
}

func ExampleDeposit_limit() {
	card := types.Card{Balance: 2000000, Active: true}
	Deposit(&card, 5000001)
	fmt.Println(card.Balance)
	//Output: 2000000
}

func ExampleAddBonus_positive() {
	card := types.Card{Balance: 2000, MinBalance: 1050, Active: true}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	//Output: 2002
}

func ExampleAddBonus_inactive() {
	card := types.Card{Balance: 5000, MinBalance: 1000, Active: false}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	//Output: 5000
}

func ExampleAddBonus_negativeBalance() {
	card := types.Card{Balance: -1000, MinBalance: -5000, Active: false}
	AddBonus(&card, 3, 30, 365)
	fmt.Println(card.Balance)
	//Output: -1000
}

func ExampleAddBonus_limit() {
	card := types.Card{Balance: 2000000, MinBalance: 1525000, Active: true}
	AddBonus(&card, 4, 30, 366)
	fmt.Println(card.Balance)
	//Output: 2005000
}

// Autotest for Function of counting sum on card balances - Total
func ExampleTotal(){
	fmt.Println(Total([]types.Card{
		{
			Balance: 10_000_00,
			Active: true,
		},
		{
			Balance: 10_000_00,
			Active: false,
		},
		{
			Balance: 5000,
			Active: true,
		},
	}))
	fmt.Println(Total([]types.Card{
		{
			Balance: 10_000_00,
			Active: true,
		},
	}))
	fmt.Println(Total([]types.Card{
		{
			Balance: 10_000_00,
			Active: false,
		},
	}))
	fmt.Println(Total([]types.Card{
		{
			Balance: -10_000_00,
			Active: true,
		},
	}))
	//Output: 
	//1005000	
	//1000000	
	//0
	//0
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