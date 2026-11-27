//ff:type feature=data type=model
//ff:what failAfterWriter — a writer that takes one write and then fails, for the tests of a write error after the first
package main

import "errors"

type failAfterWriter struct{ wrote bool }

func (w *failAfterWriter) Write(p []byte) (int, error) {
	if w.wrote {
		return 0, errors.New("full")
	}
	w.wrote = true
	return len(p), nil
}
