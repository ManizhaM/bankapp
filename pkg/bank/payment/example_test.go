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