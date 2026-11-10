//ff:func feature=cli type=command control=sequence
//ff:what TestExecute — help exits with 0; an unknown flag exits with 1
package cli

import (
	"os"
	"testing"
)

func TestExecute(t *testing.T) {
	old := os.Args
	defer func() { os.Args = old }()
	os.Args = []string{"nlgeul", "--help"}
	if Execute() != 0 {
		t.Fatal("help")
	}
	os.Args = []string{"nlgeul", "--no-such-flag"}
	if Execute() != 1 {
		t.Fatal("unknown flag")
	}
}
