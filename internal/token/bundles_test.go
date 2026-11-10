//ff:func feature=token type=reader control=sequence
//ff:what TestBundles — 아버지가방에들어가신다 is cut where its readings agree, and a stretch they read differently holds each way once
package token

import "testing"

func TestBundles(t *testing.T) {
	sc := scanOf("아버지가방에들어가신다")
	got := sc.bundles(sc.search())
	if len(got) != 3 || got[0].text != "아버지가" || len(got[0].branches) != 1 || got[1].text != "방에들어가" || len(got[1].branches) < 2 {
		t.Fatalf("%+v", got)
	}
}
