//ff:func feature=parse type=reader control=iteration dimension=3
//ff:what TestModifierWrittenName — end to end, over the small word table with a noun added: a name no table lists that begins with the letters of the numeral determiner 한 and goes on with those of a known noun (한송별은, 두송별이라고) is still a candidate as one new word with its tail after it (한송별+은, 두송별+이+라고), with the verdict of a new word rebutted once more (2/(1+forceNew+forceModifier)-1), while the first way stays the one that reads the known noun apart (송별+은, 송별+이라고); a name with no such modifier at its head (세송별은) has a new word's own verdict
package parse

import (
	"math"
	"testing"
)

func TestModifierWrittenName(t *testing.T) {
	r := streamReader(t, "words.tsv")
	r.ix.lex["송별"] = word{kinds: kindNominal, count: 400, by: [6]int{400}}
	lowered, plain := 2/(1+forceNew+forceModifier)-1, 2/(1+forceNew)-1
	for _, c := range []struct {
		in, first, whole string
		verdict          float64
	}{
		{"어제 한송별은 집에 갔다.", "송별[열쇠:체언]+은[보조사:은]", "한송별+은[보조사:은]", lowered},
		{"그는 두송별이라고 했다.", "송별[열쇠:체언]+이라고[부사격:이라고]", "두송별+이[서술격:이]+라고[맺음:라고]", lowered},
		{"어제 세송별은 집에 갔다.", "", "세송별+은[보조사:은]", plain},
	} {
		verdicts := map[string]float64{}
		for _, rg := range netRegions(r, c.in) {
			for _, n := range rg.nodes {
				for _, b := range n.Contest.Branches {
					verdicts[b.String()] = b.Verdict
				}
			}
		}
		if v, ok := verdicts[c.whole]; !ok || math.Abs(v-c.verdict) > 1e-9 {
			t.Errorf("%s: the whole stem as a new word (%s) is a candidate: %v, with the verdict %v, want %v", c.in, c.whole, ok, v, c.verdict)
		}
		if c.first == "" {
			continue
		}
		words := map[string]bool{}
		for _, b := range bestNodes(r, c.in) {
			words[b.String()] = true
		}
		if !words[c.first] || words[c.whole] {
			t.Errorf("%s: the first way does not read the known noun apart: %v", c.in, words)
		}
	}
}
