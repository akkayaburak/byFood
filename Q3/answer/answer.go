package answer

import "fmt"

func FindMostRepeatedData[T comparable](arr []T) T {
	m := map[T]int{}
	var maxCnt int
	var most T
	for _, value := range arr {
		m[value]++
		if m[value] > maxCnt {
			maxCnt = m[value]
			most = value
		}
	}
	fmt.Println(most)
	return most
}
