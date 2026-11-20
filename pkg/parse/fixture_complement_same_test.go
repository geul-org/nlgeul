//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what complementSame — checks that the ruled node has every candidate of the plain node under the same name, and no other
package parse

import "testing"

func complementSame(t *testing.T, s string, plain, ruled Node) {
	t.Helper()
	if len(plain.Contest.Branches) != len(ruled.Contest.Branches) {
		t.Fatal(s, plain.Contest.Branches, ruled.Contest.Branches)
	}
	for _, br := range plain.Contest.Branches {
		if _, at := branchNamed(ruled, br.String()); at < 0 {
			t.Fatal(s, "lost", br.String())
		}
	}
}
