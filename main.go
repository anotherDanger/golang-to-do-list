package main

import (
	"fmt"
	list "golang-to-do-list/todo"
)

func main() {
	for {
		println("1. Menambahkan todo")
		println("2. Lihat daftar todo")
		println("3. Update todo")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Masukkan todo: ")
			var todo string
			fmt.Scan(&todo)
			list.AddTodos(todo)
			fmt.Println()

		case 2:
			fmt.Println("Daftar todos")
			list.ShowTodos()
			fmt.Println()
		case 3:
			fmt.Println("Update")
			fmt.Print("Masukkan id: ")
			var id int
			fmt.Scan(&id)
			list.UpdateTodos(id)
			fmt.Println()
		}

	}
}
