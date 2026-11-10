//ff:func feature=cli type=command control=sequence
//ff:what TestMainRun — help exits with 0; an unknown flag exits with 1
package main

import (
	"os"
	"testing"
)

func TestMainRun(t *testing.T) {
	code := -1
	exit = func(c int) { code = c }
	old := os.Args
	defer func() { exit, os.Args = os.Exit, old }()
	os.Args = []string{"nlgeul", "--help"}
	main()
	if code != 0 {
		t.Fatalf("exit %d", code)
	}
	os.Args = []string{"nlgeul", "--no-such-flag"}
	main()
	if code != 1 {
		t.Fatalf("exit %d", code)
	}
}
