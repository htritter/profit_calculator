package main

import (
	"fmt"
)

// To Do
// Ask for input for Revenue
// Ask for input for Expense
// Ask for input for  Tax Rate

// Calculate earnings before tax(EBT)
// Calculate earning after tax(profit)

// Calculate ratio (EBT/profit)

// Output EBT, profit and ratio

func main() {
	var revenue float64
	var expense float64
	var taxRate float64

	fmt.Print("Please enter in your revenue in decimal format: ")
	fmt.Scan(&revenue)
	fmt.Print("Please enter in your expenses in decimal format: ")
	fmt.Scan(&expense)
	fmt.Print("Please enter in your expected taxes in decimal format: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expense
	afterTax := revenue * taxRate
	ratio := ebt / revenue

	fmt.Println("Your income before taxes is and after expenses is ", ebt)
	fmt.Println("Your income after taxes is ", afterTax)
	fmt.Println("Your ratio is ", ratio)

}
