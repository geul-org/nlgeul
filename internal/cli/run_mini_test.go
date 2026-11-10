//ff:func feature=cli type=command control=sequence
//ff:what TestRunMini — the user's example prints one line per word; a failing judgment is returned
package cli

import (
	"bytes"
	"errors"
	"testing"
)

func TestRunMini(t *testing.T) {
	var out bytes.Buffer
	if err := runMini(&out, "이것은 문장을 분해한다"); err != nil || out.String() != "이것+은:관형어|주어|목적어\n문장+을:목적어|관형어\n분해+한+다:서술어\n" {
		t.Fatalf("%v %q", err, out.String())
	}
	old := lines
	defer func() { lines = old }()
	lines = func(string) ([]string, error) { return nil, errors.New("x") }
	if runMini(&out, "가") == nil {
		t.Fatal("failure not returned")
	}
}
