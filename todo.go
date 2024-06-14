package todo

import (
	"encoding/json"
	"errors"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
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

func (t *Todos) Print() {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl := table.New("#", "Task Title", "Done", "Created At", "Completed At")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for idx, item := range *t {
		tbl.AddRow(idx+1, item.Task, item.Done, item.CreatedAt, item.CompletedAt)
	}

	tbl.Print()
}

func (t *Todos) Complete(idx int) error {
	list := *t
	if idx < 0 || idx >= len(list) {
		return errors.New("invalid index")
	}

	list[idx].CompletedAt = time.Now()
	list[idx].Done = true
	return nil
}

func (t *Todos) Delete(idx int) error {
	list := *t
	if idx < 0 || idx >= len(list) {
		return errors.New("invalid index")
	}

	*t = append(list[:idx], list[idx+1:]...)

	return nil
}

func (t *Todos) Load(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return errors.New("file lenght is zero")
	}

	err = json.Unmarshal(file, t)

	if err != nil {
		return err
	}

	return nil
}

func (t *Todos) Store(filename string) error {
	data, err := json.Marshal(t)

	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}
