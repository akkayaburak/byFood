package answer

import (
	answer "byFood/Q1/answer"
	"testing"
)

func TestAnswerWithValidArgument(t *testing.T) {
	got := answer.SortByCharacter([]string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"}, "a")
	want := []string{"aaaasd", "aaabcd", "aab", "a", "lklklklklklklklkl", "cssssssd", "fdz", "ef", "kf", "zc", "l"}

	for i, v := range got {
		if v != want[i] {
			t.Errorf("got %q, wanted %q", got, want)
			t.FailNow()
		}
	}
}

func TestAnswerWithOneArgument(t *testing.T) {
	got := answer.SortByCharacter([]string{"aaaasd"}, "a")
	want := []string{"aaaasd"}

	for i, v := range got {
		if v != want[i] {
			t.Errorf("got %q, wanted %q", got, want)
			t.FailNow()
		}
	}
}

func TestAnswerWithNilArray(t *testing.T) {
	got := answer.SortByCharacter(nil, "a")
	if got != nil {
		t.Errorf("got %q, wanted nil", got)
		t.FailNow()
	}
}

func TestAnswerWithNumericChar(t *testing.T) {
	got := answer.SortByCharacter([]string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"}, "1")
	want := []string{"lklklklklklklklkl", "cssssssd", "aaaasd", "aaabcd", "aab", "fdz", "ef", "kf", "zc", "a", "l"}

	for i, v := range got {
		if v != want[i] {
			t.Errorf("got %q, wanted %q", got, want)
			t.FailNow()
		}
	}
}
