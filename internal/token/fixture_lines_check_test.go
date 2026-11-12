//ff:func feature=token type=helper control=iteration dimension=1
//ff:what checkLines — report as test errors the lines a reading's output (joined with newlines) lacks of those it must hold and the ones it holds that it must not
package token

import (
	"strings"
	"testing"
)

func checkLines(t *testing.T, sentence, got string, has, lacks []string) {
	t.Helper()
	for _, want := range has {
		if !strings.Contains(got, want) {
			t.Errorf("%s: missing %q in\n%s", sentence, want, got)
		}
	}
	for _, no := range lacks {
		if strings.Contains(got, no) {
			t.Errorf("%s: %q in\n%s", sentence, no, got)
		}
	}
}
