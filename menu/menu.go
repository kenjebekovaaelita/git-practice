package menu

import "fmt"

func ShowMainMenu() {

	reset := "\033[0m"
	blue := "\033[34m"
	green := "\033[32m"
	yellow := "\033[33m"

	fmt.Println()

	fmt.Printf("%s--- Главное меню ---%s\n", blue, reset)

	fmt.Printf("%s1%s. Добавить запись о настроении\n", green, reset)
	fmt.Printf("%s2%s. Показать все записи\n", green, reset)
	fmt.Printf("%s3%s. Показать статистику за неделю\n", green, reset)
	fmt.Printf("%s4%s. Показать график настроения\n", green, reset)
	fmt.Printf("%s5%s. Ресурсы помощи\n", green, reset)
	fmt.Printf("%s6%s. Выход\n", green, reset)

	fmt.Println()

	fmt.Printf("%sВыберите действие: %s", yellow, reset)
}
