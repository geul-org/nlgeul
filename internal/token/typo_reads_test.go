//ff:func feature=token type=reader control=sequence
//ff:what TestTypoReads — a tail after the letters of a misspelt form is read once more showing the right form and keeping the letters, with the kinds of the misspelt line; the tail already read is kept in front; no read when the right form does not take the tail, or the letters are no misspelt form
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestTypoReads(t *testing.T) {
	sc := typoScan("맞겨", layerCost)
	conn := []piece{{text: "(어)", tail: rule.TailPred}}
	base := []tail{{start: 1, pieces: conn}}
	got := sc.typoReads(base, "맞기", tail{start: 1, pieces: conn})
	if len(got) != 2 || got[0].wrote != "" || got[1].shown != "맡기" || got[1].wrote != "맞기" || got[1].kinds != kindVerb || got[1].irr != 0 {
		t.Fatal(got)
	}
	if got = sc.typoReads(nil, "맞기", tail{start: 1, pieces: []piece{{text: "이", tail: rule.TailCase}}}); len(got) != 0 {
		t.Fatal("a verb does not take a particle", got)
	}
	if got = sc.typoReads(nil, "맞이", tail{start: 1, pieces: conn}); len(got) != 0 {
		t.Fatal("letters that are no misspelt form", got)
	}
}
