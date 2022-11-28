package answer

import (
	"fmt"
	"sort"
	"strings"
	"unicode/utf8"
)

func SortByCharacter(input []string, char string) []string {
	sort.Slice(input, func(i, j int) bool {
		s1, s2 := input[i], input[j]
		count1, count2 := strings.Count(s1, char), strings.Count(s2, char)
		if count1 != count2 {
			return count1 > count2
		}
		return utf8.RuneCountInString(s1) > utf8.RuneCountInString(s2)
	})

	fmt.Println(input)
	return input
}
