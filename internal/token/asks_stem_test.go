//ff:func feature=token type=reader control=sequence
//ff:what TestAsksStem — a fused 어/아 or ㅆ asks for a predicate's stem, whether it shows or not, with a trace (봐, 줬, 켜) or none (가, 갔), but not in 해 and 했; a lone ㄴ under 하 does not ask, under another syllable it does; a consonant with letters after it does not
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAsksStem(t *testing.T) {
	sc := scanOf("가봐간한갔해했")
	if !sc.asksStem(&entry{lead: rule.LeadVowel}, 1, '가') || !sc.asksStem(&entry{lead: rule.LeadVowel}, 2, '보') || sc.asksStem(&entry{lead: rule.LeadVowel}, 6, '하') {
		t.Fatal("a fused 어/아")
	}
	if !sc.asksStem(&entry{lead: rule.JongN}, 3, '가') || sc.asksStem(&entry{lead: rule.JongN}, 4, '하') {
		t.Fatal("a lone final consonant")
	}
	if sc.asksStem(&entry{lead: rule.JongN, form: "다"}, 3, '가') || !sc.asksStem(&entry{lead: rule.JongSS}, 5, '가') || sc.asksStem(&entry{lead: rule.JongSS}, 7, '하') {
		t.Fatal("a consonant with letters after it, or ㅆ")
	}
	traced := scanOf("줬켜켰")
	if !traced.asksStem(&entry{lead: rule.JongSS}, 1, '주') || !traced.asksStem(&entry{lead: rule.LeadVowel}, 2, '켜') || !traced.asksStem(&entry{lead: rule.JongSS}, 3, '켜') {
		t.Fatal("a fused piece that shows in the letters")
	}
}
