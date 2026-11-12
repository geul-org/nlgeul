//ff:func feature=token type=helper control=sequence
//ff:what TestContestPlural — a contest of two branches or more is plural; one branch or no branch is not
package token

import "testing"

func TestContestPlural(t *testing.T) {
	br := Branch{Verdict: 1}
	low := Branch{Verdict: 0.3}
	if !(Contest{Text: "물을", Branches: []Branch{br, br}}).Plural() {
		t.Fatal("two branches are not plural")
	}
	if !(Contest{Text: "물을", Branches: []Branch{br, low, low}}).Plural() {
		t.Fatal("three branches, the last two with a low verdict, are not plural")
	}
	if (Contest{Text: "물을", Branches: []Branch{br}}).Plural() {
		t.Fatal("one branch is plural")
	}
	if (Contest{Text: "물을"}).Plural() {
		t.Fatal("no branch is plural")
	}
}
