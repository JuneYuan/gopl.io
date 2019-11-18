package zexercise

import "io"

// LimitReader accepts an io.Reader r and a number of bytes n,
// and returns another Reader that reads from r but reports an end-of-file condition after n bytes.
func LimitReader(r io.Reader, n int64) io.Reader {
	return nil
}
