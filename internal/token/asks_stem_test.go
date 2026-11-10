//ff:func feature=token type=reader control=sequence
//ff:what TestAsksStem — an 어/아 that leaves no trace asks for a predicate's stem, one that shows does not; a lone ㄴ under 하 does not, under another syllable it does; a consonant with letters after it, and a fused ㅆ, never
package token

import "testing"

func TestAsksStem(t *testing.T) {
	sc := scanOf("가봐간한갔")
	if !sc.asksStem(&entry{lead: leadVowel}, 1, '가') || sc.asksStem(&entry{lead: leadVowel}, 2, '보') {
		t.Fatal("a fused 어/아")
	}
	if !sc.asksStem(&entry{lead: jongN}, 3, '가') || sc.asksStem(&entry{lead: jongN}, 4, '하') {
		t.Fatal("a lone final consonant")
	}
	if sc.asksStem(&entry{lead: jongN, form: "다"}, 3, '가') || sc.asksStem(&entry{lead: jongSS}, 5, '가') {
		t.Fatal("a consonant with letters after it, or ㅆ")
	}
}
