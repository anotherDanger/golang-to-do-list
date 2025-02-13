package main

import (
	"fmt"
	list "golang-to-do-list/todo"
)

func main() {
	for {
		println("1. Menambahkan todo")
		println("2. Lihat daftar todo")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Masukkan todo: ")
			var todo string
			fmt.Scanln(&todo)
			list.AddTodos(todo)

		case 2:
			fmt.Println("Daftar todos")
			list.ShowTodos()
		}
	}
}
