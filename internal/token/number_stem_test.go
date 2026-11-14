//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestNumberStem — a number: a stem named 수, Sino-Korean letters, a native numeral; not another stem, nor a part that is no stem
package token

import "testing"

func TestNumberStem(t *testing.T) {
	for _, p := range []Part{{Stem: true, Text: "3", Name: Name{Kind: "수"}}, {Stem: true, Text: "삼십"}, {Stem: true, Text: "세"}} {
		if !numberStem(p) {
			t.Fatal(p)
		}
	}
	for _, p := range []Part{{Stem: true, Text: "씨"}, {Text: "삼십"}} {
		if numberStem(p) {
			t.Fatal(p)
		}
	}
}
