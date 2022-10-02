package store

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
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
