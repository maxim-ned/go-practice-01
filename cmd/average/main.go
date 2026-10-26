package main

import "fmt"

func main() {
	time1 := 1
	time2 := 1
	time3 := 2
	sum := time1 + time2 + time3
	var average float64 = float64(sum) / 3
	fmt.Printf("Сумма: %d, среднее: %.2f", sum, average)
}
