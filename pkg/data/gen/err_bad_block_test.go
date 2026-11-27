//ff:func feature=data type=helper control=sequence
//ff:what TestErrBadBlock — the error names the line
package main

import (
	"strings"
	"testing"
)

func TestErrBadBlock(t *testing.T) {
	if err := errBadBlock("x"); !strings.Contains(err.Error(), `"x"`) {
		t.Fatal(err)
	}
}
