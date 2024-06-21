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
		fmt.Fprintf(flag.CommandLine.Output(), "\nPuede exportar la variable TODO_FILENAME para usar ")
		fmt.Fprintf(flag.CommandLine.Output(), "\nun archivo de tareas predeterminado. Ejemplo:")
		fmt.Fprintf(flag.CommandLine.Output(), "\necho export TODO_FILENAME='$HOME/.todo.json' >> ~/.zshrc\n")
		flag.PrintDefaults()
		fmt.Fprintln(flag.CommandLine.Output(), "")
	}
	// parsing command line flags
	add := flag.Bool("a", false, "Add task to the ToDo list")
	list := flag.Bool("l", false, "List all tasks")
	complete := flag.Int("c", 0, "Item to be completed")
	delete := flag.Int("d", 0, "Item to be deleted")
	noShowComplete := flag.Bool("x", false, "No show completed tasks")

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
		todo.NoShowDone = *noShowComplete
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
		tasks, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		l.Add(tasks)
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
func getTask(r io.Reader, args ...string) ([]string, error) {
	// tomar la tarea desde args

	myList := []string{}
	if len(args) > 0 {
		myList = append(myList, strings.Join(args, " "))
		return myList, nil
	}

	// tomar la(s) tarea(s) desde STDIN
	s := bufio.NewScanner(r)
	for s.Scan() {
		myList = append(myList, s.Text())
	}

	if err := s.Err(); err != nil {
		return myList, err
	}

	return myList, nil
}
