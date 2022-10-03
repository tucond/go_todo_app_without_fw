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
	}
}
