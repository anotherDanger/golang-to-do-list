package todo

import "fmt"

type Task struct {
	id       int
	todo     string
	complete bool
}

var lists []*Task
var next = 1

func AddTodos(task string) {
	list := &Task{id: next, todo: task, complete: false}
	lists = append(lists, list)
	next++
	fmt.Println("Sukses ditambahkan")
}

func ShowTodos() {
	if len(lists) == 0 {
		println("Kosong")
		return
	}
	fmt.Println("Daftar Todos")
	for _, todo := range lists {
		complete := "X"
		if todo.complete {
			complete = "Yes"
		}
		fmt.Println(todo.id, todo.todo, complete)
	}
}

func UpdateTodos(id int) {
	if len(lists) == 0 {
		println("Kosong")
		return
	}
	for i := range lists {
		if lists[i].id == id {
			lists[i].complete = true
		}
	}
}
