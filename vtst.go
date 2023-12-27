package vtst

import (
	"fmt"
	"io"
)

func PrintTo(w io.Writer, name string) {
	fmt.Fprintf(w, "ALIO: %s", name)
}
