package wordscounter

import (
	"bufio"
	"bytes"
)

type WordsCounter int

func (c *WordsCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewBuffer(p))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		return count, err
	}
	*c += WordsCounter(count)
	return count, nil
}
