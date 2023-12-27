package vtst_test

import (
	"bytes"
	"testing"

	"github.com/jjngx/vtst"
)

func TestPrintToWriter(t *testing.T) {
	t.Parallel()

	buf := bytes.Buffer{}

	vtst.PrintTo(&buf, "Bolek")

	want := "Alio alio: Bolek"
	got := buf.String()

	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}
