//ff:func feature=cli type=command control=sequence
//ff:what TestRunToken — the boxes of a sentence are printed, with plain its lines; with cuts, the word ends; several sentences are parted by an empty line
package cli

import (
	"bytes"
	"strings"
	"testing"

	"github.com/geul-org/nlgeul/internal/token"
)

func TestRunToken(t *testing.T) {
	r, _ := token.New(token.Options{})
	var out bytes.Buffer
	runToken(&out, r, []string{"아버지가"}, false, false, false, true, 0)
	if !strings.Contains(out.String(), "\n  아버지+가[주격:가]\t성립 0.33") || strings.HasSuffix(out.String(), "\n\n") {
		t.Fatalf("%q", out.String())
	}
	out.Reset()
	runToken(&out, r, []string{"아버지가", "아버지가"}, false, false, false, false, 0)
	if blocks := strings.Split(out.String(), "\n\n"); len(blocks) != 3 || !strings.HasPrefix(blocks[1], "┌ [0-4] 아버지가") || !strings.Contains(blocks[0], "\n1등 길: 아버지가") {
		t.Fatalf("%q", out.String())
	}
	out.Reset()
	runToken(&out, r, []string{"아버지가", "아버지가"}, false, true, false, false, 0)
	if blocks := strings.Split(out.String(), "\n\n"); len(blocks) != 3 || !strings.HasPrefix(blocks[1], "아버지가\t4;") {
		t.Fatalf("%q", out.String())
	}
}
