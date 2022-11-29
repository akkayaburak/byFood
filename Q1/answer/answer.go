package answer

import (
	"fmt"
	"sort"
	"strings"
	"unicode/utf8"
)

/*
Sorts the given string array by the given string.
@param : sortedBy - string to be sorted by.
@param : input - []string to be sorted.
*/
func SortByCharacter(input []string, sortedBy string) []string {
	sort.Slice(input, func(i, j int) bool {
		s1, s2 := input[i], input[j]
		count1, count2 := strings.Count(s1, sortedBy), strings.Count(s2, sortedBy)
		if count1 != count2 {
			return count1 > count2
		}
		//len() returns the UTF-8 encoded byte length, not the number of runes. For the letter, use utf8.RuneCountInString()
		return utf8.RuneCountInString(s1) > utf8.RuneCountInString(s2)
	})

	fmt.Println(input)
	return input
}
