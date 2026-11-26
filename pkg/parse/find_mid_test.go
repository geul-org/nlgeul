//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestFindMid — a middle dot between letters is found at each place (가·나·다 has two), also ㆍ and ・ and the variants ‧ ․ • (가운뎃점); a slash (빗금) and a tilde or its variants (물결표) only with Hangul on both sides; a dash ― or — or the box-drawing line ─ (줄표), a run of them one part; a run of two hyphens or more (줄표) but not one hyphen, nor a run with a sign that draws next to it (-->); none for dots in a run, at an edge, or right after the head of a unit; the candidates are one 가운뎃점 part printing the dot; none where an atom is noted to begin
package parse

import "testing"

func TestFindMid(t *testing.T) {
	sc := scanOf("가·나ㆍ다・라")
	if len(sc.mid) != 3 || sc.mid[1].candidates[0].Text != "·" || sc.mid[3].candidates[0].Text != "ㆍ" || sc.mid[5].candidates[0].Text != "・" {
		t.Fatal(sc.mid)
	}
	if r := sc.mid[1].candidates; len(r) != 1 || r[0].Name.Kind != "가운뎃점" || r[0].Stem {
		t.Fatal("the candidate", r)
	}
	variants := scanOf("가‧나․다•라")
	if len(variants.mid) != 3 || variants.mid[1].candidates[0].Text != "‧" || variants.mid[3].candidates[0].Text != "․" || variants.mid[5].candidates[0].Text != "•" || variants.mid[5].candidates[0].Name.Kind != "가운뎃점" {
		t.Fatal("the middle dot variants", variants.mid)
	}
	for text, kind := range map[string]string{"여름/가을": "빗금", "여름~가을": "물결표", "서울∼부산": "물결표", "서울～부산": "물결표", "서울〜부산": "물결표", "답이다―누구": "줄표", "답이다—누구": "줄표", "답이다─누구": "줄표"} {
		got := scanOf(text)
		at := len([]rune(text)) - 3
		if len(got.mid) != 1 || len(got.mid[at].candidates) != 1 || got.mid[at].candidates[0].Name.Kind != kind {
			t.Errorf("%q: %v, want %s at %d", text, got.mid, kind, at)
		}
	}
	if run := scanOf("가―――나"); len(run.mid) != 1 || run.mid[1].candidates[0].Text != "―――" {
		t.Fatal("a run of dashes is one part", run.mid)
	}
	for _, text := range []string{"가/b", "a/가", "1~가", "가~3", "가……나", "가-나", "가-->나", "가--|나"} {
		if got := scanOf(text); len(got.mid) != 0 {
			t.Errorf("%q: %v", text, got.mid)
		}
	}
	if run := scanOf("갔다---그리고"); len(run.mid) != 1 || len(run.mid[2].candidates) != 1 || run.mid[2].candidates[0].Text != "---" || run.mid[2].candidates[0].Name.Kind != "줄표" {
		t.Fatal("a run of two hyphens or more is a dash", run.mid)
	}
	for _, text := range []string{"가··나", "·가", "가·", "가나", "가.나"} {
		if got := scanOf(text); len(got.mid) != 0 {
			t.Errorf("%q: %v", text, got.mid)
		}
	}
	if len(scanOf("가·(나").mid) != 1 || len(scanOf("가)·나").mid) != 1 {
		t.Fatal("a sign next to a middle dot does not hold it back")
	}
	head := scanOf(")·가·나")
	head.head = 1
	head.prepare()
	if len(head.mid) != 1 || head.mid[3].candidates[0].Text != "·" {
		t.Fatal("not right after the head", head.mid)
	}
	if len(scanOf("A·B").mid) != 1 || len(scanOf("3·5").mid) != 1 {
		t.Fatal("between Latin runs")
	}
	held := scanOf("가·나")
	held.atoms[1] = atom{end: 2}
	held.findMid()
	if len(held.mid) != 0 {
		t.Fatal("not where an atom begins", held.mid)
	}
}
