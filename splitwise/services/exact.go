package services

type Exact struct {
}

func (expenseType *Exact) SettleExpense(expense *Expense) {

	for index, user := range expense.users {

		if user == expense.createdBy.GetUserID() {
			continue
		}
		expense.createdBy.AddBalance(user, float64(expense.dividends[index]))
		// get(user).AddBalance(expense.createdBy.userId, -float64(expense.dividends[index]))
	}

}
