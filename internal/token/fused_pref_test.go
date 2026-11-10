//ff:func feature=token type=reader control=sequence
//ff:what TestFusedPref — a fused 어/아 and a fused ㅆ take the preference of the whole form they are written out as; a piece under another consonant, and a form the index gives no count, keep the preference given
package token

import "testing"

func TestFusedPref(t *testing.T) {
	ix, _ := newIndex(Options{})
	vowel := ix.named(keyVowel, "맺음:어서(합침)")
	if got := ix.fusedPref(vowel, "아", 0); got != ix.pref(closingCounts["맺음:아서"], "아서") || got >= 0 {
		t.Fatal(got)
	}
	past := ix.named('ㅆ', "선어말:ㅆ")
	if got := ix.fusedPref(past, "았", 0); got != ix.pref(ix.named('았', "선어말:았").count, "았") || got >= 0 {
		t.Fatal(got)
	}
	if got := ix.fusedPref(ix.named('ㄴ', "꾸밈:ㄴ"), "ㄴ", -7); got != -7 {
		t.Fatal(got)
	}
	if got := ix.fusedPref(&entry{name: "맺음:없는꼴", lead: leadVowel, form: "꼴"}, "어", -3); got != -3 {
		t.Fatal(got)
	}
}
