package answer

import "fmt"

func Recursive(n int, at int) int {
	if n == 2 {
		fmt.Println(n)
		return n
	}
	fmt.Println(n)

	return Recursive((n-at)/2, at-1)
}
