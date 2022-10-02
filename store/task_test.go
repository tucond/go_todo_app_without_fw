package store

import (
	"context"
	"testing"
)

func TestRepository_ListTasks(t *testing.T) {
	ctx := context.Background()
	tx, err := OpenDBForTest(t).BeginTxx(ctx, nil)
}
