package lang

import (
	"bufio"
	"fmt"
	"io"
)

func Repl(r io.Reader) {
	scanner := bufio.NewScanner(r)
	scan := func() bool {
		fmt.Print("> ")
		res := scanner.Scan()
		return res
	}
	for scan() {
		output, err := Run(scanner.Text())
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(output)
		}
	}
}
