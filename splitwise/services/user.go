package services

import "fmt"

type User struct {
	userId       string
	name         string
	email        string
	mobileNumber string
	Balance      float64
	expenses     map[string]float64 // userid to amount
}

func NewUser(userId, name, email, mobileNumber string) *User {

	user := &User{
		userId, //TODO fix
		name,
		email,
		mobileNumber,
		0,
		make(map[string]float64),
	}

	add(user)

	return user
}

func (user *User) getBalance() string {

	return "No Balance"
}

func (user *User) getExpenses() {
	if len(user.expenses) == 0 {
		fmt.Println("No balances")
	}

	for key, val := range user.expenses {
		fmt.Printf("%s owes %s : %g", key, user.userId, val)
		fmt.Println("")
	}
}

func (user *User) CreateExpense(expType string, amount float64, numberOfUsers int, userIds []string, dividends []int) *Expense {

	expense := &Expense{
		GetExpenseSettlerObject(expType),
		user,
		amount,
		numberOfUsers,
		false,
		userIds,
		dividends,
	}

	expense.Settle()

	return expense

}

func (user *User) GetUserID() string {
	return user.userId
}

func (user *User) AddBalance(userId string, amount float64) {

	_, ok := user.expenses[userId]

	// check if user owes something to userId

	owedUser := get(userId)
	owedUserExpense, exsts := owedUser.expenses[user.GetUserID()]

	if exsts {

		netAmount := owedUserExpense - amount

		if netAmount > 0 {
			user.expenses[userId] = netAmount
		} else {
			owedUser.expenses[user.GetUserID()] = -netAmount
		}

	} else {

		if ok {
			user.expenses[userId] += amount
		} else {
			user.expenses[userId] = amount
		}
	}

}

func ShowExpenseFor(userId string) {

	user := get(userId)

	user.getExpenses()

}

func ShowExpenses() {

	for _, user := range userMap {

		user.getExpenses()

	}

}
