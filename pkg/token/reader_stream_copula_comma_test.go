//ff:func feature=token type=reader control=sequence
//ff:what TestReaderStreamCopulaComma — 이나 right before a comma reads as the copula 이 with the ending 나 first, its preference exactly bonusShortTaken above what the stream with no rules gives it, while the stream with no rules reads it as the auxiliary particle first; the trace says the rule stood; with no comma after it the first candidate stays the particle
package token

import (
	"math"
	"strings"
	"testing"
)

func TestReaderStreamCopulaComma(t *testing.T) {
	const copula, particle = "책[열쇠:체언]+이[서술격:이]+나[맺음:나]", "책[열쇠:체언]+이나[보조사:이나]"
	got, notes := copulaBranches(t, "그건 책이나, 얇다.", "책이나", true)
	plain, _ := copulaBranches(t, "그건 책이나, 얇다.", "책이나", false)
	if len(plain) == 0 || plain[0].String() != particle {
		t.Fatalf("with no rules the particle comes first, got %v", plain)
	}
	if len(got) == 0 || got[0].String() != copula {
		t.Fatalf("the first candidate is the copula, got %v", got)
	}
	after, ok := quotedPref(got, copula)
	before, _ := quotedPref(plain, copula)
	if !ok || math.Abs(after-before-bonusShortTaken) > 1e-9 {
		t.Errorf("%v -> %v, want exactly bonusShortTaken more", before, after)
	}
	if !strings.Contains(strings.Join(notes, "\n"), copula+"\t"+ruleCopulaComma+"\t1.00\t통과") {
		t.Errorf("no 통과 line of %s in\n%s", ruleCopulaComma, strings.Join(notes, "\n"))
	}
	if bare, _ := copulaBranches(t, "책이나 가자.", "책이나", true); len(bare) == 0 || bare[0].String() != particle {
		t.Errorf("with no comma the particle stays first, got %v", bare)
	}
}
