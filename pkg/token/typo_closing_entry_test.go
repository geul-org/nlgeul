//ff:func feature=token type=reader control=sequence
//ff:what TestTypoClosingEntry — the entry of a closing written wrong copies the right closing's name, class, count, head and ways with the copula, takes its own form and lead, has the right closing's place to attach after for a form of whole syllables and any place for a form under a final consonant, and keeps the letters and the right closing as it is printed (in brackets for a final consonant); the right closing is not changed
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestTypoClosingEntry(t *testing.T) {
	gel := &entry{name: "맺음:ㄹ게", form: "게", class: classClosing, lead: rule.JongL, after: rule.AfterAny, count: 77, head: headL, noYo: true}
	e := typoClosingEntry(gel, typoClosing{"께", rule.JongL, "맺음:ㄹ게"})
	if e == gel || e.name != gel.name || e.class != classClosing || e.count != 77 || e.head != headL || !e.noYo || e.g != gel.g {
		t.Fatal("the right closing's own", e)
	}
	if e.form != "께" || e.lead != rule.JongL || e.after != rule.AfterAny || e.wrote != "ㄹ께" || e.right != "(ㄹ)게" {
		t.Fatal(e)
	}
	if gel.form != "게" || gel.wrote != "" || gel.right != "" {
		t.Fatal("the right closing is changed", gel)
	}
	seup := &entry{name: "맺음:습니다", form: "습니다", class: classClosing, after: rule.AfterBatchim, count: 5}
	e = typoClosingEntry(seup, typoClosing{"읍니다", 0, "맺음:습니다"})
	if e.form != "읍니다" || e.lead != 0 || e.after != rule.AfterBatchim || e.wrote != "읍니다" || e.right != "습니다" || e.count != 5 {
		t.Fatal(e)
	}
	ryeo := &entry{name: "맺음:려고", form: "려고", class: classClosing, after: rule.AfterHangul, eu: true}
	e = typoClosingEntry(ryeo, typoClosing{"려고", rule.JongL, "맺음:려고"})
	if e.after != rule.AfterAny || e.wrote != "ㄹ려고" || e.right != "려고" || !e.eu {
		t.Fatal("a final consonant written for a whole closing", e)
	}
}
