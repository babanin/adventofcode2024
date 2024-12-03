package day3

import (
	"bytes"
	"testing"
)

func TestSolve(t *testing.T) {
	spyReader := bytes.NewBufferString(`mul(1,5)xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))mul(8,5)mul(8,mul(8,5))`)
	spyWriter := bytes.Buffer{}

	Solve(spyReader, &spyWriter)

	got := spyWriter.String()
	want := `246`

	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}
