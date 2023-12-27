package bingo_test

import (
	"bytes"
	"testing"

	"github.com/jjngx/vtst/v2/internal/bingo"
)

func TestPrintBingo(t *testing.T) {
	t.Parallel()

	buf := bytes.Buffer{}

	bingo.PrintInternalBingo(&buf, "server start")

	want := "BINGO: server start"
	got := buf.String()

	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}
