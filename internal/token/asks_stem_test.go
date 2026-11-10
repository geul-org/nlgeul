//ff:func feature=token type=reader control=sequence
//ff:what TestAsksStem — a fused 어/아 or ㅆ asks for a predicate's stem, whether it shows or not, but not in 해 and 했; a lone ㄴ under 하 does not ask, under another syllable it does; a consonant with letters after it does not
package token

import "testing"

func TestAsksStem(t *testing.T) {
	sc := scanOf("가봐간한갔해했")
	if !sc.asksStem(&entry{lead: leadVowel}, 1, '가') || !sc.asksStem(&entry{lead: leadVowel}, 2, '보') || sc.asksStem(&entry{lead: leadVowel}, 6, '하') {
		t.Fatal("a fused 어/아")
	}
	if !sc.asksStem(&entry{lead: jongN}, 3, '가') || sc.asksStem(&entry{lead: jongN}, 4, '하') {
		t.Fatal("a lone final consonant")
	}
	if sc.asksStem(&entry{lead: jongN, form: "다"}, 3, '가') || !sc.asksStem(&entry{lead: jongSS}, 5, '가') || sc.asksStem(&entry{lead: jongSS}, 7, '하') {
		t.Fatal("a consonant with letters after it, or ㅆ")
	}
}
