package main

import (
	"lld/todo/services"
	"time"
)

func main() {

	user1 := services.CreateUser("Alisha", 1)

	task1 := services.CreateTask(1, "FirstTask", "some description", time.Date(2025, 8, 15, 14, 30, 45, 100, time.Local), user1.Id)
	task2 := services.CreateTask(1, "Second task", "vaguee description", time.Date(2025, 1, 15, 14, 30, 45, 100, time.Local), user1.Id)

	user1.AddTask(task1)
	user1.AddTask(task2)

}
