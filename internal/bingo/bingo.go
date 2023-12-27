package bingo

import (
	"fmt"
	"io"
)

func PrintInternalBingo(w io.Writer, msg string) {
	fmt.Fprintf(w, "BINGO: %s", msg)
}
