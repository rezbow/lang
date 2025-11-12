package main

import (
	"os"

	"github.com/rezbow/lang"
)

func main() {
	lang.Repl(os.Stdin)
}
