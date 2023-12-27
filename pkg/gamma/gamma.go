package gamma

import (
	"fmt"
	"os"

	"github.com/jjngx/vtst/internal/bingo"
	"github.com/jjngx/vtst/pkg/internal/beta"
)

func PrintGamma() {
	fmt.Println("in pkg/gamma/")
}

func PrintBeta() {
	beta.PrintBeta()
}

func PrintBingo() {
	bingo.PrintBingo(os.Stdout, "alio")
}
