package main

import "fmt"

// statusLabel превращает код статуса в понятную человеку подпись.
func statusLabel(status string) string {
	switch status {
	case "new":
		return "Новая"
	case "in_progress":
		return "В работе"
	case "done":
		return "Завершена"
	default:
		return "Неизвестный статус"
	}
}

// countStatus считает количество элементов с заданным статусом.
func countStatus(statuses []string, target string) int {
	count := 0
	for _, status := range statuses {
		if status == target {
			count++
		}
	}
	return count
}

func findStatus(statuses []string, target string) int {
	for i, status := range statuses {
		if status == target {
			return i
		}
	}
	return -1
}

func comppletionPercent(statuses []string) float64 {
	length := len(statuses)
	if length == 0 {
		return 0
	}
	return 100 * float64(countStatus(statuses, "done")) / float64(length)
}

func main() {
	statuses := []string{"new", "in_progress", "done", "done"}
	statuses = append(statuses, "new")

	fmt.Printf("Задач всего: %d\n", len(statuses))
	for i, status := range statuses {
		fmt.Printf("%d. %s\n", i+1, statusLabel(status))
	}
	fmt.Printf("Завершено: %d\n", countStatus(statuses, "done"))

	counts := make(map[string]int)
	for _, status := range statuses {
		counts[status]++
	}
	fmt.Printf("Статистика: new=%d, in_progress=%d, done=%d\n", counts["new"], counts["in_progress"], counts["done"])
}
