//ff:func feature=token type=reader control=selection
//ff:what closingEntries — the index entries of one closing of the table: a whole-syllable form (and its 으 form, which like 은, 을 and 음 attaches after a consonant other than ㄹ), the three whole forms and the fused form of an 어/아 closing, or a form under a final consonant; each with the observed count this layer keeps for its form (the fused form of an 어/아 closing has none of its own)
package token

import (
	"github.com/geul-org/nlgeul/internal/parse"
	"github.com/geul-org/nlgeul/pkg/rule"
)

func closingEntries(m parse.Marker) []*entry {
	name := closingLabel(m.Kind) + ":"
	base := entry{class: classClosing, ends: m.Kind & (rule.KindFinal | rule.KindLink), noYo: m.NoYo, bars: m.NoIda, needs: m.NeedIda, holds: m.SelfIda}
	make1 := func(form string, lead, after int, shown string) *entry {
		e := base
		e.name, e.form, e.lead, e.after = name+shown, form, lead, after
		e.count = closingCounts[e.name]
		return &e
	}
	switch {
	case m.Lead == 0 && m.Eu:
		base.eu = true
		return []*entry{make1(m.Form, 0, afterOpen, m.Form), make1("으"+m.Form, 0, rule.AfterHard, "으"+m.Form)}
	case m.Lead == 0:
		return []*entry{make1(m.Form, 0, euAfter(m.Form, m.After), m.Form)}
	case m.Lead < 0:
		return []*entry{
			make1("어"+m.Form, 0, afterAny, "어"+m.Form), make1("아"+m.Form, 0, afterAny, "아"+m.Form),
			make1("여"+m.Form, 0, afterAny, "여"+m.Form), make1(m.Form, leadVowel, afterAny, "어"+m.Form+"(합침)"),
		}
	}
	return []*entry{make1(m.Form, m.Lead, afterAny, rule.LeadText(m.Lead)+m.Form)}
}
