//ff:func feature=cli type=helper control=sequence
//ff:what TestWriteTokenJSON — every sentence is one line ending in a newline, numbered from 0, the Korean letters as they are; an output that cannot be written is an error
package cli

import (
	"bytes"
	"strings"
	"testing"

	"github.com/geul-org/nlgeul/pkg/token"
)

func TestWriteTokenJSON(t *testing.T) {
	r, err := token.New(token.Options{})
	if err != nil {
		t.Fatal(err)
	}
	var out bytes.Buffer
	if err := writeTokenJSON(&out, r, []string{"물을", "아버지가"}); err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(strings.TrimSuffix(out.String(), "\n"), "\n")
	if len(lines) != 2 || !strings.HasPrefix(lines[0], `{"sentence":0,"text":"물을"`) || !strings.HasPrefix(lines[1], `{"sentence":1,"text":"아버지가"`) {
		t.Fatalf("%q", out.String())
	}
	if err := writeTokenJSON(failWriter{}, r, []string{"물을"}); err == nil {
		t.Fatal("a closed output accepted")
	}
}
