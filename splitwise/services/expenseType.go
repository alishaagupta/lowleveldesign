package services

func GetExpenseSettlerObject(expenseType string) ExpenseMethod {

	switch expenseType {
	case "EQUAL":
		return &Equal{}
	case "EXACT":
		return &Exact{}
	case "PERCENTAGE":
		return &Percentage{}
	}

	//TODO : throw error
	return nil
}
