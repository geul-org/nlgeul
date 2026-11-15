//ff:func feature=token type=reader control=sequence
//ff:what surnameCandidate — article 48: when its rule is on, a stem that begins the unit, has no marker after it (no piece fused into it, not written wrong, no word shown for it, not read as a number) and is a family name of the list, right before a title or an office of the list that ends the unit or is followed by a particle (titleAfter), is read once more as that family name: the name 성:<letters>, the kind of a nominal and the verdict of the family name's graph; nil otherwise
package token

import "slices"

// The candidate is one more way the stem stands; the candidates the
// stem has otherwise are kept as they are.
func (sc *scan) surnameCandidate(b int, t tail) []stemCandidate {
	stem := sc.letters(b, t.start)
	if !sc.ix.spacing["48"] || b != 0 || len(t.pieces) != 0 || t.as != "" || t.wrote != "" || t.shown != "" || t.base != 0 {
		return nil
	}
	if !slices.Contains(surnameWords, stem) || !sc.titleAfter(t.start) {
		return nil
	}
	return []stemCandidate{{kindCandidate{kindNominal, setSurname + ":" + stem, sc.ix.surname}, stem, ""}}
}
