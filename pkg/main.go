// Проект калькулятор для задачи SCRUM-12
package main

import "fmt"

func intToWord(n int) string {
	switch n {
	case 1:
		return "one"
	case 2:
		return "two"
	case 3:
		return "three"
	case 4:
		return "four"
	case 5:
		return "five"
	case 6:
		return "six"
	case 7:
		return "seven"
	case 8:
		return "eight"
	case 9:
		return "nine"
	case 10:
		return "ten"
	case 11:
		return "eleven"
	case 12:
		return "twelve"
	case 13:
		return "thirteen"
	case 14:
		return "fourteen"
	case 15:
		return "fifteen"
	case 16:
		return "sixteen"
	case 17:
		return "seventeen"
	case 18:
		return "eighteen"
	case 19:
		return "nineteen"
	case 20:
		return "twenty"
	default:
		return "big number"
	}
}

func myCalculation(a int, b int, op string) string {
	if op != "+" && op != "*" {
		return "invalid operation"
	}

	var result int
	if op == "+" {
		result = a + b
	} else {
		result = a * b
	}
	return intToWord(a) + " " + op + " " + intToWord(b) + " = " + intToWord(result)
}

func main() {
	for i := 1; i <= 20; i++ {
		fmt.Println(myCalculation(i, 3, "+"))
		fmt.Println(myCalculation(i, 3, "*"))
	}
}
