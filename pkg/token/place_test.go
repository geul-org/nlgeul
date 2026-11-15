//ff:func feature=token type=reader control=sequence
//ff:what TestPlace — a whole form reports the end before it; a form at the unit's start has no stem; a piece under a consonant or fused into a vowel sits in the stem's last syllable; a fused ㅆ only on the vowels it fuses into, an unseen 어/아 only with more of the tail or at the unit's end; after 되 an 어 the writing leaves out, with more of the tail, before 요 or at the unit's end but not inside it, and a fused ㅆ under it (됬); a form that is not there is not placed; a form at the unit's start has a stem in the unit before while the unit is read as attaching; a whole form right after a middle dot (·, ㆍ, ‧) has no stem, right after another sign (/) it has
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestPlace(t *testing.T) {
	sc := scanOf("다음으로간다해서")
	if start, p, ok := sc.place(&entry{form: "으로"}, 4, false); !ok || start != 2 || !p.HasStem || p.Prev != '음' {
		t.Fatal(start, p, ok)
	}
	if _, p, ok := sc.place(&entry{form: "다"}, 1, false); !ok || p.HasStem {
		t.Fatal(p, ok)
	}
	if start, p, ok := sc.place(&entry{form: "다", lead: rule.JongN}, 6, false); !ok || start != 5 || p.Base != "가" {
		t.Fatal(start, p, ok)
	}
	if _, _, ok := sc.place(&entry{form: "다", lead: rule.JongL}, 6, false); ok {
		t.Fatal("wrong consonant")
	}
	if start, _, ok := sc.place(&entry{form: "서", lead: rule.LeadVowel}, 8, false); !ok || start != 7 {
		t.Fatal(start, ok)
	}
	if _, _, ok := scanOf("있다").place(&entry{lead: rule.JongSS}, 1, false); ok {
		t.Fatal("ㅆ under a vowel 었/았 does not fuse into")
	}
	if _, _, ok := scanOf("갔다").place(&entry{lead: rule.JongSS}, 1, false); !ok {
		t.Fatal("갔")
	}
	if _, _, ok := scanOf("가다").place(&entry{lead: rule.LeadVowel}, 1, false); ok {
		t.Fatal("an unseen fusion inside a unit")
	}
	if _, _, ok := scanOf("가보다").place(&entry{lead: rule.LeadVowel}, 1, false); !ok {
		t.Fatal("an unseen fusion right before an auxiliary verb's stem")
	}
	if _, _, ok := scanOf("가").place(&entry{lead: rule.LeadVowel}, 1, false); !ok {
		t.Fatal("an unseen fusion at the unit's end")
	}
	if start, _, ok := scanOf("되도").place(&entry{form: "도", lead: rule.LeadVowel}, 2, false); !ok || start != 1 {
		t.Fatal("an 어 left out after 되 with more of the tail", start, ok)
	}
	if _, _, ok := scanOf("보도").place(&entry{form: "도", lead: rule.LeadVowel}, 2, false); ok {
		t.Fatal("no 어 left out after a syllable not of unspelt")
	}
	if _, _, ok := scanOf("되요").place(&entry{lead: rule.LeadVowel}, 1, true); !ok {
		t.Fatal("an 어 left out after 되 before 요")
	}
	if _, _, ok := scanOf("되").place(&entry{lead: rule.LeadVowel}, 1, false); !ok {
		t.Fatal("an 어 left out after 되 at the unit's end")
	}
	if _, _, ok := scanOf("되다").place(&entry{lead: rule.LeadVowel}, 1, false); ok {
		t.Fatal("no 어 left out after 되 inside a unit")
	}
	if start, _, ok := scanOf("됬다").place(&entry{lead: rule.JongSS}, 1, false); !ok || start != 1 {
		t.Fatal("a fused ㅆ under 되 with the 어 left out", start, ok)
	}
	if _, _, ok := sc.place(&entry{form: "로써"}, 4, false); ok {
		t.Fatal("form not there")
	}
	if _, _, ok := sc.place(&entry{form: "다음으로간"}, 3, false); ok {
		t.Fatal("longer than the text")
	}
	if _, _, ok := sc.place(&entry{lead: rule.JongN}, 0, false); ok {
		t.Fatal("a fused piece at the start")
	}
	if start, p, ok := scanOf("갑·을").place(&entry{form: "을"}, 3, false); !ok || start != 2 || p.HasStem || p.Prev != '·' {
		t.Fatal("no stem right after a middle dot", start, p, ok)
	}
	if _, p, ok := scanOf("읍ㆍ면").place(&entry{form: "면"}, 3, false); !ok || p.HasStem {
		t.Fatal("no stem right after ㆍ", p)
	}
	if _, p, ok := scanOf("벽보‧를").place(&entry{form: "를"}, 4, false); !ok || p.HasStem {
		t.Fatal("no stem right after a middle dot variant", p)
	}
	if _, p, ok := scanOf("갑/을").place(&entry{form: "을"}, 3, false); !ok || !p.HasStem || p.Prev != '/' {
		t.Fatal("a stem before a form right after another sign", p)
	}
	lone := scanOf("로")
	if _, p, ok := lone.place(&entry{form: "로"}, 1, false); !ok || p.HasStem {
		t.Fatal("no stem at the unit's start", p)
	}
	lone.attach = true
	if start, p, ok := lone.place(&entry{form: "로"}, 1, false); !ok || start != 0 || !p.HasStem {
		t.Fatal("a stem in the unit before", p)
	}
}
