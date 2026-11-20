//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what auxLinkedHas — the test stops unless every line given is among the trace lines
package parse

import (
	"slices"
	"testing"
)

func auxLinkedHas(t *testing.T, notes []string, lines ...string) {
	t.Helper()
	for _, line := range lines {
		if !slices.Contains(notes, line) {
			t.Fatalf("no line %q in %q", line, notes)
		}
	}
}
