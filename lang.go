package lang

import (
	"strconv"
	"strings"
	"unicode"
)

func sum(nums []int) (result int) {
	for _, n := range nums {
		result += n
	}
	return
}

func sub(nums []int) (result int) {
	result = nums[0]
	for idx, n := range nums {
		if idx == 0 {
			continue
		}
		result -= n
	}
	return
}

func number(index int, runes []rune) (int, int) {
	var n int
	for ; index < len(runes) && unicode.IsDigit(runes[index]); index++ {
		num, _ := strconv.Atoi(string(runes[index]))
		n = n*10 + num
	}
	return n, index
}

func Run(exp string) int {
	exp = strings.TrimSpace(exp)
	nums := make([]int, 0)
	runes := []rune(exp)

	var arithmeticFunc func([]int) int

	for i := 0; i < len(runes); {
		if unicode.IsDigit(runes[i]) {
			var num int
			num, i = number(i, runes)
			nums = append(nums, num)
			continue
		}
		switch runes[i] {
		case '+':
			arithmeticFunc = sum
		case '-':
			arithmeticFunc = sub
		}
		i++
	}
	return arithmeticFunc(nums)
}
