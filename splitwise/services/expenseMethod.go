package services

type ExpenseMethod interface {
	SettleExpense(expense *Expense)
}

type Context struct {
	expenseSettler ExpenseMethod
}

func (c *Context) SetExpenseSettler(settler ExpenseMethod) {
	c.expenseSettler = settler
}

func (c *Context) ExecuteExpenseSettler(expense *Expense) {
	c.expenseSettler.SettleExpense(expense)
}
