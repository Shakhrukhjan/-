package main

import (
	"bank/pkg/bank/card"
	"fmt"
)

func main() {
	temp := card.IssueCard("TJS", "White", "HumoCard")
	fmt.Println(temp)
}
