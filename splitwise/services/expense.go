package services

type Expense struct {
	expenseType   ExpenseMethod
	createdBy     *User
	amount        float64
	numberOfUsers int
	isSettled     bool
	users         []string
	dividends     []int //optional param
}

func (expense *Expense) Settle() {
	expense.expenseType.SettleExpense(expense)
}
