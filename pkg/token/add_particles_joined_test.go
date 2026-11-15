//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestAddParticlesJoined — the shortened particle 서 is not read on a stem it makes a word of the word table with: 요구서는 gives 요구서+는 and no 요구+서+는 of either case (요구서 is a noun of the table), 혼자서 gives the adverb and no 혼자+서 of either case; where the two make no word 서 stands as before (학교서: 학교+서 of both cases)
package token

import "testing"

func TestAddParticlesJoined(t *testing.T) {
	r := streamReader(t, "words-joined.tsv")
	for _, c := range []struct {
		text, line string
		stands     bool
	}{
		{"요구서는", "요구서[열쇠:체언]+는[보조사:는]", true},
		{"요구서는", "요구[열쇠:체언]+서[부사격:서]+는[보조사:는]", false},
		{"요구서는", "요구[열쇠:체언]+서[주격:서]+는[보조사:는]", false},
		{"혼자서", "혼자서[열쇠:부사]", true},
		{"혼자서", "혼자[열쇠:체언]+서[부사격:서]", false},
		{"혼자서", "혼자[열쇠:체언]+서[주격:서]", false},
		{"학교서", "학교[열쇠:체언]+서[부사격:서]", true},
		{"학교서", "학교[열쇠:체언]+서[주격:서]", true},
	} {
		if lines := branchLines(r, c.text); lines[c.line] != c.stands {
			t.Errorf("%s: %s should stand: %v, in %v", c.text, c.line, c.stands, lines)
		}
	}
}
