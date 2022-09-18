package store

import (
	"context"
	"errors"

	"github.com/jmorion/sqlx"
	"github.com/tucond/go_todo_app_without_fw/entity"
)

var (
	Tasks = &TaskStore{Tasks: map[int]*entity.Task{}}

	ErrNotFound = errors.New("not found")
)

type TaskStore struct {
	//仮実装
	LastID entity.TaskID
	Tasks  map[entity.TaskID]*entity.Task
}

func (ts *TaskStore) Add(t *entity.Task) (int, error) {
	ts.LastID++
	t.ID = ts.LastID
	ts.Tasks[t.ID] = t
	return t.ID, nil
}

func (ts *TaskStore) All() entity.Tasks {
	tasks := make([]*entity.Task, len(ts.Tasks))
	for i, t := range ts.Tasks {
		tasks[i-1] = t
	}
	return tasks
}

func (r *Repository) ListTasks(
	ctx context.Context, db *sqlx.DB,
) (entity.Tasks, error) {
	tasks := entity.Tasks{}
	sql := `SELECT
	    id, user_id, title,
	    status, created, modified
		FROM task;`
	if err := db.SelectionContext(ctx, &tasks, sql); err != nil {
		return nil, err
	}
	return tasks, nil
}
