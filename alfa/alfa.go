package alfa

import (
	"os"

	"github.com/jjngx/vtst/internal/bingo"
)

func Print() {
	bingo.PrintBingo(os.Stdout, "in ALFA")
}
