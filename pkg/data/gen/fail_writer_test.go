//ff:type feature=data type=model
//ff:what failWriter — a writer that always fails, for the tests of write errors
package main

import "errors"

type failWriter struct{}

func (failWriter) Write([]byte) (int, error) { return 0, errors.New("full") }
