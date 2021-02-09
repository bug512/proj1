package tasks

import (
	"fmt"
)

func GetTasks() {
	fmt.Println("Tasks list")
}

func ProbeGolang(sum float64) (x, y float64) {
	x = sum * 2
	y = sum / 5

	return
}
