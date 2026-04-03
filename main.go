package main

import "fmt"

// intToWord переводит числа от 1 до 20 в текст
func intToWord(n int) string {
	words := []string{
		"", "one", "two", "three", "four", "five",
		"six", "seven", "eight", "nine", "ten",
		"eleven", "twelve", "thirteen", "fourteen", "fifteen",
		"sixteen", "seventeen", "eighteen", "nineteen", "twenty",
	}

	if n >= 1 && n <= 20 {
		return words[n]
	} else {
		return "big number"
	}
}

// myCalculation выполняет операцию op (+ или *) и возвращает строку с результатом
func myCalculation(a int, b int, op string) string {
	var result int

	if op == "+" {
		result = a + b
	} else if op == "*" {
		result = a * b
	} else {
		panic("unsupported operation")
	}

	return fmt.Sprintf("%s %s %s = %s",
		intToWord(a),
		op,
		intToWord(b),
		intToWord(result),
	)
}

func main() {
	for i := 1; i <= 20; i++ {
		fmt.Println(myCalculation(i, 3, "+"))
		fmt.Println(myCalculation(i, 3, "*"))
	}
}
