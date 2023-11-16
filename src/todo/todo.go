package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/alexeyco/simpletable"
	"os"
	"time"
)

type item struct {
	Task        string    `json:"task"`
	Done        bool      `json:"done"`
	CreatedAt   time.Time `json:"createdAt"`
	CompletedAt time.Time `json:"completedAt"`
}

type Todos []item

func (t *Todos) Add(task string) {
	todo := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*t = append(*t, todo)
}

func (t *Todos) Complete(index int) error {
	todos := *t
	if index <= 0 || index > len(todos) {
		return errors.New("invalid index")
	}

	todos[index-1].Done = true
	todos[index-1].CompletedAt = time.Now()
	return nil
}

func (t *Todos) Delete(index int) error {
	todos := *t
	if index <= 0 || index > len(todos) {
		return errors.New("invalid index")
	}

	*t = append(todos[:index-1], todos[index:]...)
	return nil
}

func (t *Todos) Print() {

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			&simpletable.Cell{Align: simpletable.AlignCenter, Text: "#"},
			&simpletable.Cell{Align: simpletable.AlignCenter, Text: "Task"},
			&simpletable.Cell{Align: simpletable.AlignCenter, Text: "Done?"},
			&simpletable.Cell{Align: simpletable.AlignCenter, Text: "CreatedAt"},
			&simpletable.Cell{Align: simpletable.AlignCenter, Text: "CompletedAt"},
		},
	}

	var cells [][]*simpletable.Cell

	for i, task := range *t {
		cells = append(cells, *&[]*simpletable.Cell{
			&simpletable.Cell{Text: fmt.Sprintf("%d", i+1)},
			&simpletable.Cell{Text: task.Task},
			&simpletable.Cell{Text: fmt.Sprintf("%t", task.Done)},
			&simpletable.Cell{Text: task.CreatedAt.String()},
			&simpletable.Cell{Text: task.CompletedAt.String()},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		&simpletable.Cell{Align: simpletable.AlignCenter, Text: "your todos are here", Span: 5},
	}}

	table.SetStyle(simpletable.StyleUnicode)
	table.Println()
}

func (t *Todos) Store(jsonPath string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return os.WriteFile(jsonPath, data, 0644)
}

func (t *Todos) Load(jsonPath string) error {
	file, err := os.OpenFile(jsonPath, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	if err = decoder.Decode(t); err != nil {
		return err
	}
	return nil
}
