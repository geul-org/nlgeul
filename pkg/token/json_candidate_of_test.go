//ff:func feature=token type=helper control=sequence
//ff:what TestJSONCandidateOf — the parts of the words of a branch in order, its verdict and preference, and its lambda (null for a verdict of 0)
package token

import "testing"

func TestJSONCandidateOf(t *testing.T) {
	br := Branch{Words: []Eojeol{{Parts: []Part{{Text: "물", Stem: true}, {Text: "을"}}}}, Verdict: 1, Pref: -2}
	c := jsonCandidateOf(br)
	if len(c.Parts) != 2 || c.Parts[1].Text != "을" || c.Verdict != 1 || c.Pref != -2 || c.Lambda == nil || *c.Lambda != -2 {
		t.Fatal(c)
	}
	if c = jsonCandidateOf(Branch{}); c.Parts == nil || len(c.Parts) != 0 || c.Lambda != nil {
		t.Fatal(c)
	}
}
