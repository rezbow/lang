package lang

import (
	"bufio"
	"fmt"
	"io"
)

func Repl(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Println(Run(scanner.Text()))
	}
}
