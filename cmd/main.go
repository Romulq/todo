package main

import "github.com/Romulq/todo"

func main() {
	server := new(todo.Server)
	server.Run("8000")
}
