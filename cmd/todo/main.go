package main

import (
	"flag"
	"fmt"
	"github.com/justafewwords4/todo_go"
	"os"
	// "strings"
)

// hardcoding el nombre de archivo a usar
const todoFileName = ".todo.json"

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s tool. Developed for The Pragmatic Bookshelf\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Copyright 2020\n")
		fmt.Fprintf(flag.CommandLine.Output(), "\nUsage information\n")
		flag.PrintDefaults()
		fmt.Fprintln(flag.CommandLine.Output(), "")
	}
	// parsing command line flags
	task := flag.String("t", "", "Task to be included in the ToDo list")
	list := flag.Bool("l", false, "List all tasks")
	complete := flag.Int("c", 0, "Item to be completed")

	flag.Parse()

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
		for _, item := range *l {
			if !item.Done {
				fmt.Println(item.Task)
			}
		}
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
	case *task != "":
		// agregar la nueva tarea
		l.Add(*task)
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
