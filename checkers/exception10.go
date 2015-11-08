package checkers

import (
	m "github.com/AntoineAugusti/moduluschecking/models"
)

// Perform the check for the exception 10
func PerformException10Check(b m.BankAccount, scData m.SortCodeData) bool {
	if !scData.IsException(10) {
		panic("Should be exception of type 10")
	}

	if isException10(b) {
		zeros := []int{0, 0, 0, 0, 0, 0, 0, 0}
		scData.Weights = append(zeros, scData.Weights[8:]...)
		return PerformRegularCheck(b, scData)
	}

	return PerformRegularCheck(b, scData)
}

// Check if a bank account matches the criteria of the exception 10
func isException10(account m.BankAccount) bool {
	// if ab = 09 or ab = 99 and g = 9
	a := account.NumberAtPosition("a")
	b := account.NumberAtPosition("b")
	g := account.NumberAtPosition("g")

	return (a == 0 || a == 9) && b == 9 && g == 9
}