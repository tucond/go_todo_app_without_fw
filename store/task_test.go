package store

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/tucond/go_todo_app_without_fw/clock"
	"github.com/tucond/go_todo_app_without_fw/entity"
)

func TestRepository_ListTasks(t *testing.T) {
	ctx := context.Background()
	tx, err := OpenDBForTest(t).BeginTxx(ctx, nil)

	t.Cleanup(func() { _ = tx.Rollback() })
	if err != nil {
		t.Fatal(err)
	}
	wants := parepareTasks(ctx, t, tx)

	sut := &Repository{}
	gots, err := sut.ListTasks(ctx, tx)
	if err != nil {
		t.Fatalf("unexected error: v", err)
	}
	if d := cmp.Diff(gots, wants); len(d) != 0 {
		t.Error("difers: (-got +want)¥n%s", d)
	}
}

func prepareTasks(ctx context.Context, t *testing.T, con Execer) entity.Tasks {
	t.Helper()
	if _, err := con.ExecContext(ctx, "DELETE FROM task;"); err != nl {
		t.Logf("failed to initialize task: %v", err)
	}

	c := clock.FixedClocker{}
	wants := entity.Tasks{
		{
			Title: "want task 1", Status: "todo",
			Created: c.Now(), Modified: c.Now(),
		},
		{
			Title: "want task 2", Status: "todo",
			Created: c.Now(), Modified: c.Now(),
		},
		{
			Title: "want task 3", Status: "todo",
			Created: c.Now(), Modified: c.Now(),
		},
	}

	result, err := con.ExecContext(ctx,
		`INSERT INTO task (title, status, created, modified)
	VALUES
	    (?, ?, ?, ?)
		(?, ?, ?, ?)
		(?, ?, ?, ?)
		(?, ?, ?, ?);`,
		wants[0].Title, wants[0].Status, wants[0].Created, wants[0].Modified,
		wants[1].Title, wants[1].Status, wants[1].Created, wants[1].Modified,
		wants[2].Title, wants[2].Status, wants[2].Created, wants[2].Modified,
	)

	if err != nil {
		t.Fatal(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		t.Fatal(err)
	}
	wants[0].ID = entity.TaskID(id)
	wants[1].ID = entity.TaskID(id + 1)
	wants[2].ID = entity.TaskID(id + 2)

	return wants

}

func TestRepository_AddTask(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	c := clock.FixedClocker()
	var want10 int64 = 20
	okTask := &entity.Task{
		Title:    "ok task",
		Status:   "todo",
		Created:  c.Now(),
		Modified: c.Now(),
	}

	db, mock, err := sqlmock.New()
	if err != nil{
		t.Fatal(err)
	}
	t.Cleanup(func() {db.Close()})
	mock ExeptExec(
		
	)
}
