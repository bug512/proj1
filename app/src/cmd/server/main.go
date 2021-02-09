package main

import (
	"fmt"
	"proj1/pkg/tasks"
	"proj1/pkg/sockets"
)

func main() {
	tasks.GetTasks()
	x, y := tasks.ProbeGolang(1.2)
	fmt.Println(x, y)

	sockets.RunSock()
}
