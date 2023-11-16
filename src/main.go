package main

import (
	"flag"
	"fmt"
	"github.com/kouhei-github/todo-list-cli-tool/todo"
	"os"
)

const (
	todoFile = "./storages/todos.json"
)

func main() {
	add := flag.Bool("add", false, "タスクを追加する")

	flag.Parse()

	todos := &todo.Todos{}

	switch {
	case *add:
		todos.Add("Sample todo")
		if err := todos.Store(todoFile); err != nil {
			fmt.Fprintln(os.Stdout, err.Error())
			os.Exit(0)
		}
	default:
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(0)
	}

}
