//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestFindMid — a middle dot between letters is found at each place (가·나·다 has two), also ㆍ and ・; none for dots in a run, at an edge, or right after the head of a unit; the readings are one 가운뎃점 part printing the dot; none where an atom is noted to begin
package token

import "testing"

func TestFindMid(t *testing.T) {
	sc := scanOf("가·나ㆍ다・라")
	if len(sc.mid) != 3 || sc.mid[1].reads[0].Text != "·" || sc.mid[3].reads[0].Text != "ㆍ" || sc.mid[5].reads[0].Text != "・" {
		t.Fatal(sc.mid)
	}
	if r := sc.mid[1].reads; len(r) != 1 || len(r[0].Names) != 1 || r[0].Names[0].Kind != "가운뎃점" || r[0].Stem {
		t.Fatal("the reading", r)
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
	if len(head.mid) != 1 || head.mid[3].reads[0].Text != "·" {
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
