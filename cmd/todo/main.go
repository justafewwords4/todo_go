package main

import (
	"fmt"
	"github.com/justafewwords4/todo_go"
	"os"
	"strings"
)

// hardcoding el nombre de archivo a usar
const todoFileName = ".todo.json"

func main() {
	l := &todo.List{}

	// usar el método Get para leer los items del archivo
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// switch para seleccionar qué hacer de acuerdo a los parámetros
	switch {
	case len(os.Args) == 1:
		// se usó nombre del comando, listar las tareas
		for _, item := range *l {
			fmt.Println(item.Task)
		}
	// concatenar los argumentos y agregarlos a la lista de todo's
	default:
		item := strings.Join(os.Args[1:], " ")
		// Agregar la tarea
		l.Add(item)
		// save the new list
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
