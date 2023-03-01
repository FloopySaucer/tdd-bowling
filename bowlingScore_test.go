package tdd_bowling

import "testing"

func TestPerfectScoreIs300(t *testing.T) {

	pin := []Frame{
		Frame{10, 10, 0},
		Frame{10, 10, 0},
		Frame{10, 10, 0},
		Frame{10, 10, 0},
		Frame{10, 10, 0},
		Frame{10, 10, 0},
		Frame{10, 10, 0},
		Frame{10, 10, 0},
		Frame{10, 10, 0},
		Frame{10, 10, 10},
	}
	got := ScoreFrames()
	want := 10

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
