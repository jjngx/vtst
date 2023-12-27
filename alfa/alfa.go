package alfa

import (
	"os"

	"github.com/jjngx/vtst/internal/bingo"
)

func Print() {
	bingo.PrintInternalBingo(os.Stdout, "in ALFA")
}
