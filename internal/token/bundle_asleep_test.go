//ff:func feature=token type=helper control=sequence
//ff:what TestBundleAsleep — a bundle is asleep when its first branch is; a bundle with no branches or an awake first branch is not
package token

import "testing"

func TestBundleAsleep(t *testing.T) {
	awake := Branch{Verdict: 1}
	asleep := Branch{Verdict: 0.3, Asleep: true}
	if (Bundle{Text: "물을"}).Asleep() {
		t.Fatal("a bundle with no branches is asleep")
	}
	if (Bundle{Text: "물을", Branches: []Branch{awake, awake}}).Asleep() {
		t.Fatal("a contest of awake branches is asleep")
	}
	if !(Bundle{Text: "물을", Branches: []Branch{asleep}}).Asleep() {
		t.Fatal("a bundle of one branch asleep is not asleep")
	}
	if !(Bundle{Text: "물을", Branches: []Branch{asleep, asleep}}).Asleep() {
		t.Fatal("a bundle of branches asleep is not asleep")
	}
}
