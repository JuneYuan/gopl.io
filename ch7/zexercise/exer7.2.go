package zexercise

import "io"

// CountingWriter receives an io.Writer, then returns a new Writer that
// wraps the original, and a pointer to an int64 variable that at any moment
// contains the number of bytes written to the new writer
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	// TODO
	return nil, nil
}
