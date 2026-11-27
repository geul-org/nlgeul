//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what loneForms — the candidates of a stem that is one letter of the tables of special letters read as a letter of a run of Latin letters and digits (a full-width letter Ａ or digit １; a Roman numeral letter is read by romanForms, before its reading as a number), with nothing written onto it: one for each name the tables give the letter (Ａ[전각:로마자], １[전각:수]), with the verdict of the graph loneFormGraph, a new word's, and the kind that stands behind the other readings of those letters (kindBehind: the run's own reading, a number or a word of the word table, comes first); a stem with a piece on it and a stem of more letters are not read so
package parse

import "github.com/geul-org/nlgeul/pkg/rule"

func (ix *index) loneForms(stem string, t tail) []kindCandidate {
	r := []rune(stem)
	if len(r) != 1 || len(t.pieces) > 0 || !rule.RunRune(r[0]) || formShaped(r[0], romanShape) {
		return nil
	}
	var out []kindCandidate
	for _, name := range formNames(r[0]) {
		out = append(out, kindCandidate{kindBehind, name, ix.loneForm})
	}
	return out
}
