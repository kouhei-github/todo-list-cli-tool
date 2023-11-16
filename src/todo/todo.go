package todo

import (
	"encoding/json"
	"errors"
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
