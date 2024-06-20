package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/justafewwords4/todo_go"
	"io"
	"os"
	"strings"
)

// hardcoding el nombre de archivo a usar
var todoFileName = ".todo.json"

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s tool. Developed for The Pragmatic Bookshelf\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Copyright 2020\n")
		fmt.Fprintf(flag.CommandLine.Output(), "\nUsage information\n")
		flag.PrintDefaults()
		fmt.Fprintln(flag.CommandLine.Output(), "")
	}
	// parsing command line flags
	add := flag.Bool("a", false, "Add task to the ToDo list")
	list := flag.Bool("l", false, "List all tasks")
	complete := flag.Int("c", 0, "Item to be completed")
	delete := flag.Int("d", 0, "Item to be deleted")

	flag.Parse()

	// verificar si el usuario ha definido alguna variable de ambiente para el archivo .todo.json
	if os.Getenv("TODO_FILENAME") != "" {
		todoFileName = os.Getenv("TODO_FILENAME")
	}

	// define an item list
	l := &todo.List{}

	// use the Get command to read ToDo items from file
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// decidir qué hacer basado en las banderas
	switch {
	case *list:
		// listar los items en el archivo todo
		fmt.Print(l)
	case *complete > 0:
		// completar el item dado
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		// salvar la nueva lista
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *add:
		// agregar la nueva tarea
		// si se pasan argumentos, estos deben ser
		// usados como la nueva tarea
		t, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		l.Add(t)
		// salvar la nueva lista
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *delete > 0:
		if err := l.Delete(*delete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		// salvar la nueva lista
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		// invalid flag provided
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}
}

// la función getTask decide de donde tomar la descripción
// para una nueva tarea ya sea desde STDIN o desde args
func getTask(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	s := bufio.NewScanner(r)
	s.Scan()
	if err := s.Err(); err != nil {
		return "", err
	}

	if len(s.Text()) == 0 {
		return "", fmt.Errorf("Task cannot be blank")
	}

	return s.Text(), nil
}
