//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestTopSplit — the readings with no letters written wrong are kept as top keeps them (the k best, in the same order), the readings with some are kept apart (the k best of them), and both come out in the order of before; a typo reading never takes a plain reading's place
package token

import (
	"fmt"
	"testing"
)

func TestTopSplit(t *testing.T) {
	ix := scanOf("x").ix
	var plain, typo, all []*path
	for i := 0; i < 20; i++ {
		plain = append(plain, &path{sig: fmt.Sprint("p", i), verdict: 1, pref: -float64(i) - 10})
	}
	for i := 0; i < 20; i++ {
		typo = append(typo, &path{sig: fmt.Sprint("t", i), verdict: 1, pref: -float64(i), typo: true})
	}
	all = append(append(all, typo...), plain...)
	got := ix.topSplit(all, 16)
	want := ix.top(append([]*path(nil), plain...), 16)
	var onlyPlain []*path
	for _, p := range got {
		if !p.typo {
			onlyPlain = append(onlyPlain, p)
		}
	}
	if len(got) != 32 || len(onlyPlain) != 16 {
		t.Fatal(len(got), len(onlyPlain))
	}
	for i := range want {
		if onlyPlain[i] != want[i] {
			t.Fatal("the plain readings are those top keeps, in its order", i, onlyPlain[i].sig, want[i].sig)
		}
	}
	for i := 1; i < len(got); i++ {
		if ix.before(got[i], got[i-1]) {
			t.Fatal("not in the order of before", i)
		}
	}
	if got[0].sig != "t0" || got[len(got)-1].sig != "p15" {
		t.Fatal(got[0].sig, got[len(got)-1].sig)
	}
	if one := ix.topSplit([]*path{plain[3], plain[3]}, 16); len(one) != 1 {
		t.Fatal("the same reading once", len(one))
	}
	if none := ix.topSplit(nil, 16); len(none) != 0 {
		t.Fatal(none)
	}
}
