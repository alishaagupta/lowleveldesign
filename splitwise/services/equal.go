package services

type Equal struct{}

func (expenseType *Equal) SettleExpense(expense *Expense) {
	amountPerUser := expense.amount / float64(expense.numberOfUsers)

	for _, user := range expense.users {
		if user == expense.createdBy.GetUserID() {
			continue
		}
		expense.createdBy.AddBalance(user, amountPerUser)

		// userObj := get(user)
		// userObj.AddBalance(expense.createdBy.userId, -amountPerUser)
	}

}
