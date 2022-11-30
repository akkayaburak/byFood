package answer

import (
	answer "byFood/Q2/answer"
	"testing"
)

func TestAnswerWithValidArgument(t *testing.T) {
	got := answer.Recursive(9, 1)
	want := 2

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
		t.FailNow()
	}
}
