package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"github.com/kouhei-github/todo-list-cli-tool/todo"
	"io"
	"os"
	"strings"
)

const (
	todoFile = "./storages/todos.json"
)

func main() {
	add := flag.Bool("add", false, "タスクを追加する")
	complete := flag.Int("complete", 0, "タスクを完了させる")
	del := flag.Int("del", 0, "タスクを削除する")

	flag.Parse()

	todos := &todo.Todos{}

	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stdout, err.Error())
		os.Exit(0)
	}

	switch {
	case *add:
		task, err := getInput(os.Stdin, flag.Args()...)
		if err = todos.Store(todoFile); err != nil {
			fmt.Fprintln(os.Stdout, err.Error())
			os.Exit(0)
		}
		todos.Add(task)
		if err = todos.Store(todoFile); err != nil {
			fmt.Fprintln(os.Stdout, err.Error())
			os.Exit(0)
		}
	case *complete > 0 && *complete < len(*todos)+1:
		if err := todos.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stdout, err.Error())
			os.Exit(0)
		}
		if err := todos.Store(todoFile); err != nil {
			fmt.Fprintln(os.Stdout, err.Error())
			os.Exit(0)
		}
	case *del > 0 && *del < len(*todos)+1:
		if err := todos.Delete(*del); err != nil {
			fmt.Fprintln(os.Stdout, err.Error())
			os.Exit(0)
		}
		if err := todos.Store(todoFile); err != nil {
			fmt.Fprintln(os.Stdout, err.Error())
			os.Exit(0)
		}
	default:
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(0)
	}

}

func getInput(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	text := scanner.Text()

	if len(text) == 0 {
		return "", errors.New("todoリストがからです")
	}

	return text, nil
}
