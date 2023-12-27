package alfa

import (
	"os"

	"github.com/jjngx/vtst/v2/internal/bingo"
)

func Print() {
	bingo.PrintInternalBingo(os.Stdout, "in ALFA")
}
