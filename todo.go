package todo

import (
	"encoding/json"
	"errors"
	"fmt"

	// "go/types"
	"os"
	"time"
)

// item struct for item

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

var NoShowDone = false
var Verbose = false

// List represents a list of ToDo items
type List []item

// Add agrega una tarea a List
func (l *List) Add(tasks []string) {
	for _, task := range tasks {
		t := item{
			Task:        task,
			Done:        false,
			CreatedAt:   time.Now(),
			CompletedAt: time.Time{},
		}
		*l = append(*l, t)
	}
}

// Complete mark a ToDo item as completed
func (l *List) Complete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exists", i)
	}
	// Ajustar indice para indices basados en 0
	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()

	return nil
}

// Delete borra un item ToDo de List
func (l *List) Delete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exists", i)
	}

	// Ajustar indice para indices basados en 0
	*l = append(ls[:i-1], ls[i:]...)

	return nil
}

// Save codifica List como JSON y lo escribe a un archivo
func (l *List) Save(filename string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, js, 0644)
}

// Get abre el archivo provisto y lo convierte a JSON
func (l *List) Get(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return nil
	}
	return json.Unmarshal(file, l)
}

// String prints out a formatted list
// implements the fmt.Stringer interface
func (l *List) String() string {
	formatted := ""

	for k, t := range *l {
		prefix := " "
		if t.Done {
			if NoShowDone {
				continue
			}
			prefix = "X"
		}

		moreInfo := ""
		if Verbose {
			duration := time.Since(t.CreatedAt)
			moreInfo = fmt.Sprintf("%.2fhrs ago,", duration.Hours())
		}

		formatted += fmt.Sprintf("[%s] %s %d: %s\n", prefix, moreInfo, k+1, t.Task)
		//ajustar el número del item mediante k para imprimir en base 1, y no base 0
	}

	return formatted
}
