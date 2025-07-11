package main

import "fmt"

// TaskManager делает ВСЁ: и хранит задачи, и логирует, и отправляет уведомления
type TaskManager struct {
	tasks []string
}

func (tm *TaskManager) AddTask(task string) {
	tm.tasks = append(tm.tasks, task)
	fmt.Printf("Задача '%s' добавлена\n", task)
}

func (tm *TaskManager) CompleteTask(index int) {
	if index < 0 || index >= len(tm.tasks) {
		fmt.Println("Ошибка: неверный индекс")
		return
	}
	task := tm.tasks[index]
	tm.tasks = append(tm.tasks[:index], tm.tasks[index+1:]...)
	fmt.Printf("Задача '%s' выполнена\n", task)
}

func (tm *TaskManager) ProcessTasks() {
	for _, task := range tm.tasks {
		if len(task) > 20 {
			fmt.Printf("Длинная задача: '%s'\n", task)
		} else {
			fmt.Printf("Обычная задача: '%s'\n", task)
		}
	}
}

func main() {
	manager := TaskManager{}
	manager.AddTask("Купить молоко")
	manager.AddTask("Написать код, который нарушает SOLID, а потом исправить его")
	manager.CompleteTask(0)
	manager.ProcessTasks()
}
