package services

import "fmt"

type Percentage struct{}

func (expenseType *Percentage) SettleExpense(expense *Expense) {

	for index, user := range expense.users {

		if user == expense.createdBy.GetUserID() {
			continue
		}
		fmt.Println("Segregation on percent", user, getAmountForPercent(expense.dividends[index], expense.amount))
		expense.createdBy.AddBalance(user, getAmountForPercent(expense.dividends[index], expense.amount))
		// get(user).AddBalance(expense.createdBy.userId, -getAmountForPercent(expense.dividends[index], expense.amount))
	}
}

func getAmountForPercent(percent int, totalAmount float64) float64 {

	return (float64(percent) / 100) * totalAmount
}
