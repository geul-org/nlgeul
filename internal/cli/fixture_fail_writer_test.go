//ff:type feature=cli type=model
//ff:what failWriter — an output that cannot be written (every write is an error), for the tests of what is written
package cli

import "errors"

type failWriter struct{}

func (failWriter) Write([]byte) (int, error) { return 0, errors.New("closed") }
