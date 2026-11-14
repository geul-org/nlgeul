//ff:func feature=token type=helper control=sequence
//ff:what TestBestOf — the lambda of the first candidate of each node
package token

import "testing"

func TestBestOf(t *testing.T) {
	nodes := []Node{
		{Contest: Contest{Branches: []Branch{{Verdict: 1, Pref: -2}, {Verdict: 1, Pref: -9}}}},
		{Contest: Contest{Branches: []Branch{{Verdict: 1, Pref: -4}}}},
	}
	if got := bestOf(nodes); len(got) != 2 || got[0] != -2 || got[1] != -4 {
		t.Fatal(got)
	}
}
