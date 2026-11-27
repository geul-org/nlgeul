//ff:func feature=parse type=reader control=sequence
//ff:what stemCandidates — the candidates of the stem of a word as it is printed and looked up, each with its kind, name and verdict (one with no name and a verdict of 0 for a new word): the letters before the tail, with the last syllable as it was before a piece was fused into it, the word a fused form stands for, or known words bound by a slash or a hyphen (boundTypo: a word of them written wrong makes the right form of the whole); a stem no table lists may be two nominals of the word table bound by 사이시옷 (saisiotCandidates), and is a new word as well only up to newStemMax tokens (farNew; the stem alone, and not when it is the whole of its unit), a stem written wrong is the right form, judged as a misspelt form (misspeltCandidates) and by nothing else, with the letters it is written with; a stem of the tables is read once for each kind and each set that lets it stand (and a stem of letters that are a unit of the unit table, no Hangul syllables, once more for each line of it behind the words of the table: loneUnits, 700 m), an atom once for each name that joins it (and two runs joined by a slash alone once more as a new word: plainToo), a bound stem once for each candidate of its sign, a stem of 사이시옷 once before its new word (alone when the new word is too long), a stem of the tables and a stem no table lists once more after these as a number that is no run of digits with nothing in its word that tells it (loneNumbers: 사십), and a stem of the tables as a unit with no number written onto it (loneUnits); before these a greater number of Hangul right before the digits of its number (bigFront: 천 of 천500억); and before all of these a stem that is a path by a reading the letters alone do not settle (namedTail: the Hangul after a path's last slash as its name, a path with no head and one slash) once as that path; sole tells that the word is the whole of its unit, which the stem's graph is told, as it is told whether the stem is an auxiliary verb's written right after the 어/아 form of its main verb (auxAfterConn), whether it begins right after a number with the letters of a unit and a word written after a unit (unitWordAt), whether it is the dependent noun 것 or 거 written right after a modifying closing (boundAfterAdn), and, as main tells, whether it is a main verb's stem that begins its unit with an auxiliary verb written onto its 어/아 form (mainBeforeAux)
package parse

// A fused form's word is given by the closed set, so it is known under
// the form's name when no table has it. A reading of 사이시옷 is a guess
// made without the origin of the words (a loanword or a name may look the
// same), so the new word of the same letters stays beside it, unlike a
// word of the tables or a bound stem.
func (sc *scan) stemCandidates(b int, t tail, sole, main bool) []stemCandidate {
	stem := sc.letters(b, t.start)
	if t.base != 0 {
		stem = string(sc.s[b:t.start-1]) + string(t.base)
	}
	if t.shown != "" {
		stem = t.shown
	}
	if t.as != "" {
		return []stemCandidate{{kindCandidate{0, t.as, sc.ix.number}, stem, ""}}
	}
	if t.wrote != "" {
		if candidates := stemWith(stem, t.wrote, sc.ix.misspeltCandidates(stem, t.kinds, tailClass(t.pieces), sole)); len(candidates) > 0 {
			return candidates
		}
		if sc.farNew(b, t, sole) {
			return nil
		}
		return []stemCandidate{{kindCandidate{}, stem, t.wrote}}
	}
	front := stemWith(stem, "", sc.bigFront(b, stem, t))
	if candidates := stemWith(stem, "", sc.ix.knownCandidates(stem, tailClass(t.pieces), sole, sc.auxAfterConn(b, stem, t), sc.unitWordAt(b, stem), sc.boundAfterAdn(b, stem), main)); len(candidates) > 0 {
		return append(append(front, candidates...), stemWith(stem, "", append(sc.ix.loneUnits(stem), sc.ix.loneNumbers(stem, t)...))...)
	}
	if candidates := stemWith(stem, "", kindCandidatesOf(sc.atomStem(b, t.start))); len(candidates) > 0 {
		return plainToo(stem, candidates)
	}
	if t.label != "" {
		return []stemCandidate{{kindCandidate{0, t.label, 1}, stem, ""}}
	}
	named := stemWith(stem, "", kindCandidatesOf(sc.namedTail(b, t.start)))
	right, marks := sc.boundOf(stem, tailClass(t.pieces), t.boundTypo)
	wrote := ""
	if right != stem {
		wrote = stem
	}
	if candidates := stemWith(right, wrote, kindCandidatesOf(marks)); len(candidates) > 0 {
		return append(named, candidates...)
	}
	saisiot := append(append(front, named...), stemWith(stem, "", sc.ix.saisiotCandidates(stem, tailClass(t.pieces)))...)
	saisiot = append(saisiot, stemWith(stem, "", sc.ix.loneNumbers(stem, t))...)
	if sc.farNew(b, t, sole) {
		return saisiot
	}
	return append(saisiot, stemCandidate{kindCandidate{}, stem, ""})
}
