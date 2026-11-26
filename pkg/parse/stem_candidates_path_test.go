//ff:func feature=parse type=reader control=sequence
//ff:what TestStemCandidatesPath — two names joined by one slash are a path and the new word of the same letters; the whole unit of a path with Hangul after its last slash is that path and the new word; the letters of a path with no head and one slash are that path before the new word, and before the words bound by its slash when they are known (입력/출력-결과)
package parse

import "testing"

func TestStemCandidatesPath(t *testing.T) {
	sc := scanOf("ab/cd를")
	if got := sc.stemCandidates(0, tail{start: 5}, false, false); len(got) != 2 || got[0].name != pathName || got[0].v != 1 || got[1].name != "" {
		t.Fatal(got)
	}
	sc = scanOf("~/a/b/회의록")
	sc.search()
	if got := sc.stemCandidates(0, tail{start: 9}, true, false); len(got) != 2 || got[0].name != pathName || got[0].v != sc.ix.pathNamed || got[1].name != "" || got[1].stem != "~/a/b/회의록" {
		t.Fatal(got)
	}
	if got := sc.stemCandidates(0, tail{start: 6}, false, false); len(got) != 1 || got[0].name != pathName {
		t.Fatal("the path up to its slash", got)
	}
	sc = scanOf("notes/회의-결과")
	sc.search()
	if got := sc.stemCandidates(0, tail{start: 11}, true, false); len(got) != 2 || got[0].name != pathName || got[0].v != sc.ix.pathWeak || got[1].name != "" {
		t.Fatal(got)
	}
	ws := wordsScan("io/입력-출력", layerCost)
	ws.search()
	got := ws.stemCandidates(0, tail{start: 8}, true, false)
	if len(got) < 2 || got[0].name != pathName || got[0].v != ws.ix.pathWeak {
		t.Fatal(got)
	}
}
