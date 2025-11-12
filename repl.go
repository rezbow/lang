package lang

import (
	"bufio"
	"fmt"
	"io"
)

func Repl(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		output, err := Run(scanner.Text())
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(output)
		}
	}
}
