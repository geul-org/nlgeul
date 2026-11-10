//ff:func feature=rule type=helper control=sequence
//ff:what TestBareBases — ㅁ under three syllables, the other consonants under five
package rule

import "testing"

func TestBareBases(t *testing.T) {
	if len(BareBases(JongM)) != 3 || len(BareBases(JongN)) != 5 || BareBases(JongB)[2] != "시" {
		t.Fatal("BareBases")
	}
}
