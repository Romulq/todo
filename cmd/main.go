package main

func main() {
	server := new(todo.server)
	server.Run("8000")
}
