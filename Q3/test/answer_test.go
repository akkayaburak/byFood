package answer

import (
	answer "byFood/Q3/answer"
	"testing"
)

func TestAnswerWithValidArgumentString(t *testing.T) {
	got := answer.FindMostRepeatedData([]string{"apple", "pie", "apple", "red", "red", "red"})
	want := "red"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
		t.FailNow()
	}
}

func TestAnswerWithValidArgumentInteger(t *testing.T) {
	got := answer.FindMostRepeatedData([]int{1, 2, 3, 4, 4, 5, 6, 6, 6})
	want := 6

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
		t.FailNow()
	}
}

func TestAnswerWithOneArgument(t *testing.T) {
	got := answer.FindMostRepeatedData([]int{1})
	want := 1

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
		t.FailNow()
	}
}
