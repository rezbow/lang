package lang

import (
	"bufio"
	"fmt"
	"io"
)

func Repl(r io.Reader) {
	scanner := bufio.NewScanner(r)
	eachLine := func() bool {
		fmt.Print("> ")
		res := scanner.Scan()
		return res
	}
	for eachLine() {
		result, err := Run(scanner.Text())
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Println(result)
		}
	}
}
