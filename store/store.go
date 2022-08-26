package store

import(
	"errors"

	"github.com/tucond/go_todo_app_without_fw/entity"
)

var (
	Tasks = &TaskStore{Tasks: map[int]*entity.Task}

	ErrNotFound = errors.New("not found")
)

type TaskStore struct {
	//仮実装
	LastID entity.TaskID
	Tasks map[entity.TaskID]*entity.Task
}

func (ts *TaskStore)Add(t *entity.Task)(int, error){
	ts.LastD++
	t.ID = ts.LastID
	ts.Tasks[t.ID] = t
	return t.ID, nil
}

func(ts *TaskStore) All() entity.Tasks{
	tasks := make([]*entity.Task, len(ts.Tasks))
	for i, t := range ts.Tasks{
		tasks[i-1] = t
	}
	return tasks
}
