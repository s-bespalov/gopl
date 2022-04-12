package countingwriter

import (
	"bytes"
	"fmt"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	w, i := CountingWriter(&bytes.Buffer{})
	fmt.Fprint(w, "Test 1")
	if *i != 6 {
		t.Fail()
	}
	fmt.Fprint(w, "ab")
	if *i != 8 {
		t.Fail()
	}
	fmt.Fprint(w, "ab")
	if *i != 10 {
		t.Fail()
	}
}
