//ff:func feature=parse type=reader control=sequence
//ff:what TestSpokenEntry — the entry of a marker written as it is spoken copies the standard marker's name, class, lead, place, kind in a chain and head, takes the written form and the gate handed in, has no count, and is read only after the copula stem when its row says so; the standard marker is not changed
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestSpokenEntry(t *testing.T) {
	gate := &entry{name: predGate}
	go1 := &entry{name: "맺음:고", form: "고", class: classClosing, after: rule.AfterAny, count: 99, head: headFused}
	e := spokenEntry(go1, spoken{"구", 0, "맺음:고", false}, gate)
	if e == go1 || e.name != "맺음:고" || e.form != "구" || e.class != classClosing || e.after != rule.AfterAny || e.head != headFused || e.count != 0 || e.gate != gate || e.needs || e.wrote != "" {
		t.Fatal(e)
	}
	if go1.form != "고" || go1.count != 99 || go1.gate != nil {
		t.Fatal("the standard marker is changed", go1)
	}
	o := &entry{name: "맺음:오", form: "오", class: classClosing, after: rule.AfterVowel, noYo: true}
	e = spokenEntry(o, spoken{"요", 0, "맺음:오", true}, gate)
	if e.form != "요" || !e.needs || !e.noYo || e.after != rule.AfterVowel || o.needs {
		t.Fatal("a form read only after the copula stem", e, o)
	}
	do := &entry{name: "보조사:도", form: "도", class: classParticle, slot: rule.LinkTopic, count: 7}
	e = spokenEntry(do, spoken{"두", 0, "보조사:도", false}, nil)
	if e.form != "두" || e.class != classParticle || e.slot != rule.LinkTopic || e.gate != nil || e.count != 0 {
		t.Fatal(e)
	}
}
