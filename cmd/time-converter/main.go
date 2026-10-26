package main

import "fmt"

func main() {
	const minutes_in_hour = 60
	var time int
	fmt.Print("Введите количество минут: ")
	_, err := fmt.Scan(&time)
	if err != nil {
		fmt.Println("Ошибка: нужно ввести целое число")
		return
	}
	if time < 0 {
		fmt.Println("Ошибка: число не должно быть отрицательным")
		return
	}
	var hours int = time / minutes_in_hour
	var minutes int = time % minutes_in_hour
	fmt.Printf("Начальное число минут: %d. Итоговое число часов: %d, число минут: %d.", time, hours, minutes)
}
