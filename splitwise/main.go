package main

import (
	"fmt"
	"lld/splitwise/services"
)

func main() {

	user1 := services.NewUser("u1", "Alisha", "alishagupta262@gmail.com", "6239304604")

	user2 := services.NewUser("u2", "Noor", "noorkarkara@gmail.com", "9815552490")

	user3 := services.NewUser("u3", "Diksha", "diksha@gmail.com", "9837421734")

	user4 := services.NewUser("u4", "Arshiya", "arshiya@gmail.com", "9837421734")

	// equal := &services.Equal{}

	// expType ExpenseType, amount float64, numberOfUsers int, userIds []string, dividends []int
	services.ShowExpenseFor("u1")
	user1.CreateExpense("EQUAL", 1000, 4, []string{user1.GetUserID(), user2.GetUserID(), user3.GetUserID(), user4.GetUserID()}, nil)
	fmt.Println("------")
	services.ShowExpenseFor("u4")
	fmt.Println("------")
	services.ShowExpenseFor("u1")
	fmt.Println("------")
	user1.CreateExpense("EXACT", 1250, 2, []string{user2.GetUserID(), user3.GetUserID()}, []int{370, 880})

	services.ShowExpenses()
	fmt.Println("------")
	user4.CreateExpense("PERCENTAGE", 1200, 4, []string{user1.GetUserID(), user2.GetUserID(), user3.GetUserID(), user4.GetUserID()}, []int{40, 20, 20, 20})

	services.ShowExpenseFor("u1")
	fmt.Println("------")
	services.ShowExpenses()
	fmt.Println("------")

}
