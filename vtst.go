package vtst

import (
	"fmt"
	"io"
	"os"

	"github.com/jjngx/vtst/internal/bingo"
	"github.com/jjngx/vtst/pkg/gamma"
)

func PrintTo(w io.Writer, name string) {
	fmt.Fprintf(w, "Alio alio: %s", name)
}

func PrintBingo() {
	bingo.PrintBingo(os.Stdout, "alio")
}

func PrintGamma() {
	gamma.PrintGamma()
}
