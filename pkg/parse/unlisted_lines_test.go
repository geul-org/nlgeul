//ff:func feature=parse type=rule control=iteration dimension=1
//ff:what TestUnlistedLines — end to end, the new word's graph as wordOf and whole ask it: with a word table 떨어졌다 loses 떨어졌+다 and 졌+다 that the candidate with no table keeps, and 곰졌다, 곰졌도, 곰졌 each stand as a whole new word only; 거네 and 뭐냐 keep the dropped copula and lose the new word before the closing, while 사과 before a particle and a word no table has before a predicate tail stand, and without a word table 거+네 stands; 먹었을걸 has no new word 으 before (ㄹ)걸, while 사과을걸, where 을걸 does not read as markers, keeps it; 서, a word the table knows, stands after 마당에 though its letters read as a marker there
package parse

import (
	"strings"
	"testing"
)

func TestUnlistedLines(t *testing.T) {
	r, _ := New(Options{Words: "testdata/words.tsv"})
	bare, _ := New(Options{})
	if got := streamText(r, "떨어졌다"); strings.Contains(got, "떨어졌+") || strings.Contains(got, " 졌+") || !strings.Contains(got, "\n떨어졌다\n") {
		t.Fatal("떨어졌다 with a word table\n" + got)
	}
	if got := streamText(bare, "떨어졌다"); !strings.Contains(got, "\n떨어졌+다[맺음:다]\n") || !strings.Contains(got, " 졌+다[맺음:다]\n") {
		t.Fatal("떨어졌다 without a word table\n" + got)
	}
	segs := streamFirsts(r.Stream("곰졌다 곰졌도 곰졌"))
	if len(segs) != 3 {
		t.Fatalf("%+v", segs)
	}
	for _, sg := range segs {
		if sg.Ways.Int64() != 1 || len(sg.Nodes) != 1 || len(sg.Nodes[0].Contest.Branches) != 1 || len(sg.Nodes[0].Contest.Branches[0].Words) != 1 {
			t.Fatalf("%s: a whole new word only, %+v", segmentLine(sg), sg.Nodes)
		}
		if b := sg.Nodes[0].Contest.Branches[0]; b.String() != sg.Nodes[0].Contest.Text || b.Words[0].Parts[0].Name != (Name{}) || !b.Words[0].Parts[0].Stem {
			t.Fatalf("%s: a whole new word only, %+v", segmentLine(sg), b)
		}
	}

	r.ix.lex["뭐"] = word{kinds: kindNominal, count: 50, by: [6]int{50}}
	for in, want := range map[string][2]string{
		"거네": {"거[의존명사:거]+(이)[서술격:이(빠짐)]+네[맺음:네]", "거+네["},
		"뭐냐": {"뭐[열쇠:체언]+(이)[서술격:이(빠짐)]+냐[맺음:냐]", "뭐+냐["},
	} {
		if got := streamText(r, in); !strings.Contains(got, want[0]) || strings.Contains(got, want[1]) {
			t.Errorf("%s: want %s and no %s\n%s", in, want[0], want[1], got)
		}
	}
	if got := streamText(r, "사과는 싫다"); !strings.Contains(got, "사과[열쇠:체언]+는[보조사:는]") || !strings.Contains(got, "싫+다[맺음:다]") {
		t.Fatal("a nominal before particles, a word no table has before a predicate tail\n" + got)
	}
	if got := streamText(bare, "거네"); !strings.Contains(got, "거+네[맺음:네]") {
		t.Fatal("without a word table the new word stands\n" + got)
	}

	if got := streamText(r, "먹었을걸"); strings.Contains(got, " 으+(ㄹ)걸") || !strings.Contains(got, "먹[열쇠:동사]+었[선어말:었]+을걸[맺음:을걸]") {
		t.Fatal("먹었을걸\n" + got)
	}
	if got := streamText(r, "사과을걸"); !strings.Contains(got, " 으+(ㄹ)걸") {
		t.Fatal("사과을걸: where 을걸 does not read as markers the new word 으 stands\n" + got)
	}
	if got := streamText(r, "마당에서"); !strings.Contains(got, "\n마당[열쇠:체언]+에[부사격:에] 서[열쇠:체언]\n") {
		t.Fatal("서, a noun of the table, is not asked though its letters read as a marker there\n" + got)
	}
}
