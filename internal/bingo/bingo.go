package bingo

import (
	"fmt"
	"io"
)

func PrintBingo(w io.Writer, msg string) {
	fmt.Fprintf(w, "BINGO: %s", msg)
}
