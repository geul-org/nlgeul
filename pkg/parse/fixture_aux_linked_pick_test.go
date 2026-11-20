//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what auxLinkedPick — the candidate of a list that prints as the name given (the test stops when there is none)
package parse

import "testing"

func auxLinkedPick(t *testing.T, brs []Branch, name string) Branch {
	t.Helper()
	for _, br := range brs {
		if br.String() == name {
			return br
		}
	}
	t.Fatalf("no candidate %q", name)
	return Branch{}
}
