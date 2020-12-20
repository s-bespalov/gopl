package bytecounter

import "io"

type ByteCounter struct {
	count int64
	w     io.Writer
}

func (c *ByteCounter) Write(p []byte) (int, error) {
	n, err := c.w.Write(p)
	c.count += int64(n)
	return n, err
}

//CountingWriter returns a new Writer that wraps the original, and a pointer
//to an int64 variable that at any moment contains the number
//of bytes written to the new Writer
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := ByteCounter{0, w}
	return &c, &c.count
}
