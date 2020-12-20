package wordcounter

import (
	"bufio"
	"bytes"
)

//WordCounter stores counts of words that was written through buffer
type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	scaner := bufio.NewScanner(bytes.NewReader(p))
	scaner.Split(bufio.ScanWords)
	for scaner.Scan() {
		*c++
	}

	if error := scaner.Err(); error != nil {
		return len(p), error
	}

	return len(p), nil
}
