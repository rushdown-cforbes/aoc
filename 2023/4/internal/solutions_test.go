package internal

import "testing"

const (
	correct1 = 22488
	correct2 = 7013204
)

func TestSolution(t *testing.T) {
	answer1, answer2 := Sample.DoWork()
	if answer1 != correct1 {
		t.Errorf("The first answer was not correct\nExpected:%d\nGot:%d", correct1, answer1)
	}
	if answer2 != correct2 {
		t.Errorf("The second answer was not correct\nExpected:%d\nGot:%d", correct2, answer2)
	}
}
