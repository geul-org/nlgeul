//ff:func feature=token type=reader control=sequence
//ff:what TestLinked — 으 before a closing that takes it as a piece begins a tail after the stem and (when asked) after a slot; not after a vowel or ㄹ, not before another closing, not where no 으 stands; after a changed stem it is judged by the stem's own last syllable and the tail asks for that change
package token

import "testing"

func TestLinked(t *testing.T) {
	ni := []piece{{text: "니", names: "맺음:니", verdict: 1}}
	got := scanOf("읽으니").linked(2, ni, "맺음:니", true)
	if len(got) != 1 || got[0].start != 1 || got[0].pieces[0].names != euName || got[0].head != headEuForm || got[0].irr != 0 || len(got[0].pieces) != 2 {
		t.Fatal(got)
	}
	got = scanOf("먹었으니").linked(3, ni, "맺음:니", true)
	if len(got) < 3 || got[1].start != 1 || got[1].pieces[0].names != "선어말:었" || len(got[1].pieces) != 3 || got[2].base == 0 {
		t.Fatal("after a slot", got)
	}
	if got = scanOf("먹었으니").linked(3, ni, "맺음:니", false); len(got) != 1 {
		t.Fatal("slots not asked for", got)
	}
	if scanOf("가으니").linked(2, ni, "맺음:니", true) != nil || scanOf("살으니").linked(2, ni, "맺음:니", true) != nil {
		t.Fatal("after a vowel or ㄹ")
	}
	if scanOf("읽으고").linked(2, []piece{{text: "고"}}, "맺음:고", true) != nil {
		t.Fatal("before a closing that takes no 으")
	}
	if scanOf("읽니").linked(1, ni, "맺음:니", true) != nil || scanOf("으니").linked(1, ni, "맺음:니", true) != nil {
		t.Fatal("no 으, or nothing before it")
	}
	got = wordsScan("들으니", layerCost).linked(2, ni, "맺음:니", true)
	if len(got) != 1 || got[0].irr != irrD || got[0].start != 1 {
		t.Fatal("after a stem that changed ㄷ to ㄹ", got)
	}
	got = wordsScan("지으니", layerCost).linked(2, ni, "맺음:니", true)
	if len(got) != 1 || got[0].irr != irrS {
		t.Fatal("after a stem that dropped ㅅ", got)
	}
}
