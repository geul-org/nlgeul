//ff:func feature=cli type=command control=sequence
//ff:what TestRunToken — the lines of a sentence are printed; with cuts, the word ends
package cli

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunToken(t *testing.T) {
	var out bytes.Buffer
	if err := runToken(&out, "아버지가", false, false); err != nil || !strings.Contains(out.String(), "\n  아버지+가[주격:가|보격:가]\t성립 0.33") {
		t.Fatalf("%v %q", err, out.String())
	}
	out.Reset()
	if err := runToken(&out, "아버지가", false, true); err != nil || !strings.HasPrefix(out.String(), "아버지가\t4;") {
		t.Fatalf("%v %q", err, out.String())
	}
}
