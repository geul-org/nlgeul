//ff:func feature=parse type=reader control=sequence
//ff:what caseNumberWord — a word whose stem is a court case number (its first part, a stem named caseKind with no form) with that part as three: the year, a number and the word's stem; the mark, named caseKind with the mark as its form; the serial number, a number (2005[수]+두[사건:두]+1329[수]); the parts after the stem follow as they are; any other word is given back as it is
package parse

func caseNumberWord(w Eojeol) Eojeol {
	if len(w.Parts) == 0 || !w.Parts[0].Stem || w.Parts[0].Name != (Name{Kind: caseKind}) {
		return w
	}
	s := []rune(w.Parts[0].Text)
	year := digitsFrom(s, 0)
	mark := year + syllablesFrom(s, year)
	parts := []Part{
		{Text: string(s[:year]), Name: Name{Kind: numName}, Stem: true},
		{Text: string(s[year:mark]), Name: Name{Kind: caseKind, Form: string(s[year:mark])}},
		{Text: string(s[mark:]), Name: Name{Kind: numName}},
	}
	w.Parts = append(parts, w.Parts[1:]...)
	return w
}
