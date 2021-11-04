package payment

import "bank/pkg/bank/types"

// Max return max payment
func Max(payments []types.Payment)types.Payment{
	
	max := types.Payment{Amount: payments[0].Amount, ID: payments[0].ID}
	for _, payment := range payments {
		if(max.Amount < payment.Amount){
			max.Amount = payment.Amount
			max.ID = payment.ID
		}
	}
	return max
}
