//ff:func feature=token type=helper control=sequence
//ff:what TestBundleContest — a bundle of two awake branches is a contest; one branch, no branch, or branches asleep are not
package token

import "testing"

func TestBundleContest(t *testing.T) {
	awake := Branch{Verdict: 1}
	asleep := Branch{Verdict: 0.3, Asleep: true}
	if !(Bundle{Text: "물을", Branches: []Branch{awake, awake}}).Contest() {
		t.Fatal("two awake branches are not a contest")
	}
	if (Bundle{Text: "물을", Branches: []Branch{awake}}).Contest() {
		t.Fatal("one branch is a contest")
	}
	if (Bundle{Text: "물을"}).Contest() {
		t.Fatal("no branch is a contest")
	}
	if (Bundle{Text: "물을", Branches: []Branch{asleep, asleep}}).Contest() {
		t.Fatal("branches asleep are a contest")
	}
}
