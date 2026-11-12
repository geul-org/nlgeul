//ff:func feature=token type=helper control=sequence
//ff:what TestBundleContest — a bundle of two branches or more is a contest; one branch or no branch is not
package token

import "testing"

func TestBundleContest(t *testing.T) {
	br := Branch{Verdict: 1}
	low := Branch{Verdict: 0.3}
	if !(Bundle{Text: "물을", Branches: []Branch{br, br}}).Contest() {
		t.Fatal("two branches are not a contest")
	}
	if !(Bundle{Text: "물을", Branches: []Branch{br, low, low}}).Contest() {
		t.Fatal("three branches, the last two with a low verdict, are not a contest")
	}
	if (Bundle{Text: "물을", Branches: []Branch{br}}).Contest() {
		t.Fatal("one branch is a contest")
	}
	if (Bundle{Text: "물을"}).Contest() {
		t.Fatal("no branch is a contest")
	}
}
