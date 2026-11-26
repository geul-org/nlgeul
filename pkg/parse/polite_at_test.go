//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestPoliteAt — one polite slot before one closing's tail: the slot and the closing with the gate of a predicate's stem and the closing's own ways (하+오+며), also after a slot before it (하+시+오+며) and after the link vowel 으 (먹+으+오+며); a slot of two syllables asks for its first letters (먹+사오+며, not 먹+자오+며); a slot with no stem before it, one after a syllable it does not attach to and a closing it does not stand before give nothing
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestPoliteAt(t *testing.T) {
	ni := &entry{name: "맺음:며", form: "며"}
	of := func(start int) tail {
		return tail{start: start, close: ni, noYo: true, pieces: []piece{{text: "며", names: ni.name, verdict: 1}}}
	}
	slot := func(sc *scan, form string) *entry {
		for _, e := range sc.ix.polite {
			if e.form == form {
				return e
			}
		}
		t.Fatal("no polite slot", form)
		return nil
	}
	sc := wordsScan("하오며", layerCost)
	got := sc.politeAt(of(2), slot(sc, "오"))
	if len(got) != 1 || got[0].start != 1 || len(got[0].pieces) != 2 || got[0].pieces[0].names != "선어말:오" || got[0].pieces[0].tail != rule.TailPred || got[0].pieces[1].names != ni.name || got[0].gate != sc.ix.pred || got[0].close != ni || !got[0].noYo {
		t.Fatal(got)
	}
	sc = wordsScan("하시오며", layerCost)
	got = sc.politeAt(of(3), slot(sc, "오"))
	if len(got) != 2 || got[1].start != 1 || len(got[1].pieces) != 3 || got[1].pieces[0].names != honorName || got[1].gate != sc.ix.pred || got[1].close != ni {
		t.Fatal("after the honorific", got)
	}
	sc = wordsScan("먹으오며", layerCost)
	got = sc.politeAt(of(3), slot(sc, "오"))
	if len(got) != 2 || got[1].start != 1 || got[1].pieces[0].names != euName || got[1].gate != sc.ix.pred {
		t.Fatal("after the link vowel", got)
	}
	sc = wordsScan("먹사오며", layerCost)
	if got = sc.politeAt(of(3), slot(sc, "사오")); len(got) != 1 || got[0].start != 1 || got[0].pieces[0].text != "사오" {
		t.Fatal(got)
	}
	for i, none := range [][]tail{
		sc.politeAt(of(3), slot(sc, "자오")),
		sc.politeAt(of(3), slot(sc, "옵")),
		wordsScan("오며", layerCost).politeAt(of(1), slot(sc, "오")),
		wordsScan("먹오며", layerCost).politeAt(of(2), slot(sc, "오")),
		wordsScan("가사오며", layerCost).politeAt(of(3), slot(sc, "사오")),
	} {
		if len(none) != 0 {
			t.Errorf("%d: %v", i, none)
		}
	}
}
