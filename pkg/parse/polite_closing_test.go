//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestPoliteClosing — the closing a polite slot is read before: a whole closing right after the slot's last syllable where its first sound fits (오+며, 옵+고; not 오+고, not at the unit's end, not 오+니 where 오니 is a closing of the tables), a closing fused under 오 with ㄴ, ㄹ or ㅁ for a slot in 오 (not ㅂ), ㅂ니다 fused under 오 as the closing 니다 for a slot in 옵 (no other ㅂ form); a tail of more than its closing, one with no closing, one written wrong, one of a changed stem, one of one stem only, one fused under another syllable give none
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestPoliteClosing(t *testing.T) {
	ni := &entry{name: "맺음:며", form: "며"}
	n := &entry{name: "꾸밈:ㄴ", lead: rule.JongN}
	l := &entry{name: "꾸밈:ㄹ", lead: rule.JongL}
	m := &entry{name: "이름:ㅁ", lead: rule.JongM}
	b := &entry{name: "맺음:ㅂ니다", form: "니다", lead: rule.JongB}
	bq := &entry{name: "맺음:ㅂ니까", form: "니까", lead: rule.JongB}
	of := func(e *entry, start int, base rune) tail {
		return tail{start: start, base: base, close: e, pieces: []piece{{text: e.form, names: e.name, verdict: 1}}}
	}
	sc := scanOf("하오며")
	for i, c := range []struct {
		t    tail
		last rune
		ok   bool
	}{
		{of(ni, 2, 0), '오', true},
		{of(ni, 2, 0), '옵', false},
		{of(ni, 3, 0), '며', false},
		{of(ni, 0, 0), '오', false},
		{of(n, 2, '오'), '오', true},
		{of(l, 2, '오'), '오', true},
		{of(m, 2, '오'), '오', true},
		{of(b, 2, '오'), '오', false},
		{of(n, 2, '오'), '옵', false},
		{of(bq, 2, '오'), '옵', false},
		{of(n, 2, '시'), '오', false},
		{tail{start: 2, close: ni, pieces: []piece{{names: "선어말:시"}, {names: ni.name}}}, '오', false},
		{tail{start: 2, close: ni, pieces: []piece{{names: "맺음:으니"}}}, '오', false},
		{tail{start: 2, pieces: []piece{{names: ni.name}}}, '오', false},
		{tail{start: 2, close: &entry{name: ni.name, wrote: "늬"}, pieces: []piece{{names: ni.name}}}, '오', false},
		{tail{start: 2, close: ni, irr: 1, pieces: []piece{{names: ni.name}}}, '오', false},
		{tail{start: 2, close: ni, only: "아니", pieces: []piece{{names: ni.name}}}, '오', false},
		{tail{start: 2, close: ni, shown: "하", pieces: []piece{{names: ni.name}}}, '오', false},
		{tail{start: 2, close: ni, as: "하", pieces: []piece{{names: ni.name}}}, '오', false},
	} {
		if pc, ok := sc.politeClosing(c.t, c.last); ok != c.ok || (ok && pc.names != c.t.close.name) {
			t.Errorf("%d: %v %v", i, pc, ok)
		}
	}
	pc, ok := scanOf("하옵니다").politeClosing(of(b, 2, '오'), '옵')
	if !ok || pc.text != "니다" || pc.names != "맺음:니다" || pc.verdict != 1 {
		t.Fatal("ㅂ니다 under 오 is 옵 and 니다", pc, ok)
	}
	if _, ok := scanOf("하옵고").politeClosing(of(&entry{name: "맺음:고", form: "고"}, 2, 0), '옵'); !ok {
		t.Fatal("옵 before a consonant")
	}
	if _, ok := scanOf("하오니").politeClosing(of(&entry{name: "맺음:니", form: "니"}, 2, 0), '오'); ok {
		t.Fatal("오니 is a closing of the tables")
	}
	if _, ok := scanOf("하오고").politeClosing(of(&entry{name: "맺음:고", form: "고"}, 2, 0), '오'); ok {
		t.Fatal("오 before a consonant other than ㄴ, ㄹ, ㅁ")
	}
}
