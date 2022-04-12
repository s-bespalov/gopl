package countingwriter

import "io"

type counter struct {
	io.Writer
	int64
}

func (c *counter) Write(p []byte) (int, error) {
	i, err := c.Writer.Write(p)
	c.int64 += int64(i)
	return i, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := counter{w, 0}
	return &c, &c.int64
}
