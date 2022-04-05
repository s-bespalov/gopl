package wordscounter

import (
	"fmt"
	"testing"
)

func TestWordsCounter(t *testing.T) {
	var c WordsCounter
	fmt.Fprint(&c, "w1 w213 word3 w, w5, w6-7\n")
	if c != 6 {
		t.Fail()
	}
}
