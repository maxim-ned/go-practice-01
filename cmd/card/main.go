package main

import "fmt"

func main() {
	const course string = "Бэкенд разработка на Go"
	var tasks int = 500
	var group string = "ТЭП-03-26"
	var interest bool = true
	name := "Дональд"
	fmt.Println("Участник:", name)
	fmt.Println("Группа:", group)
	fmt.Printf("Курс: %s\n", course)
	fmt.Printf("Завершено работ: %d\n", tasks)
	fmt.Println("Интерес к бэкенду:", interest)
}
