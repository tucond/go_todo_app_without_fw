package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

)

func TestAddTask(t *testing.T){
	t.Parallel()
	type want struct {
		status int
		rspFile string
	}
	tests := map[string]struct{
		reqFile string
		want want
	}//途中
}