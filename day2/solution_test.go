package day2

import (
	"bytes"
	"testing"
)

func TestSolve(t *testing.T) {
	spyReader := bytes.NewBufferString(`7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9
`)
	spyWriter := bytes.Buffer{}

	Solve(spyReader, &spyWriter)

	got := spyWriter.String()
	want := `safe
unsafe
unsafe
unsafe
unsafe
safe
2`

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
