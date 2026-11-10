//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestSearch — 다음으로써 comes back with 다음+으로써 first and 다음+으로 써 second; no reading takes a marker for a new word's stem after a boundary; an empty unit has no readings; a unit that is a particle chain (로) comes back with its reading as markers attaching to the unit before, next to the new word, and a unit that is also a word (이) keeps the word's reading; the scan is not left attaching
package token

import "testing"

func TestSearch(t *testing.T) {
	got := scanOf("다음으로써").search()
	if len(got) < 2 || got[0].sig != "다음+으로써[부사격:으로써]" || got[1].sig != "다음+으로[부사격:으로] 써" {
		t.Fatal(got[0].sig, got[1].sig)
	}
	for _, p := range wordsScan("곰이다", layerCost).search() {
		if p.sig == "곰 이+다[맺음:다]" {
			t.Fatal("a marker read as a new word's stem after a boundary", p.sig)
		}
	}
	if scanOf("").search() != nil {
		t.Fatal("empty")
	}
	lone := scanOf("로")
	var attach, bare bool
	for _, p := range lone.search() {
		attach = attach || (p.sig == "+로[부사격:로]" && p.last.word.Attach && p.verdict <= lone.ix.unknown)
		bare = bare || p.sig == "로"
	}
	if !attach || !bare || lone.attach {
		t.Fatal("로: the attached reading next to the new word", attach, bare)
	}
	sigs := map[string]bool{}
	for _, p := range scanOf("이").search() {
		sigs[p.sig] = true
	}
	if !sigs["+이[주격:이|보격:이]"] || !sigs["이[관형사:이]"] {
		t.Fatal("이: the attached reading next to the word's", sigs)
	}
}
