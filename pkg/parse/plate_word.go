//ff:func feature=parse type=reader control=sequence
//ff:what plateWord — a word whose stem is a motor vehicle's registration number (its first part, a stem named plateKind with no form) with that part as its pieces: the kind of vehicle, a number; the letter of its use, named plateKind with the letter as its form; the serial number, a number (206[수:십진]+너[차등록:너]+9499[수:십진]); the name of a region before them is a nominal and the word's stem (서울[열쇠:체언]+34[수:십진]+가[차등록:가]+5678[수:십진]), else the kind is the stem; the parts after the stem follow as they are; any other word is given back as it is
package parse

func plateWord(w Eojeol) Eojeol {
	if len(w.Parts) == 0 || !w.Parts[0].Stem || w.Parts[0].Name != (Name{Kind: plateKind}) {
		return w
	}
	s := []rune(w.Parts[0].Text)
	region := syllablesFrom(s, 0)
	letter := region + digitsFrom(s, region)
	parts := []Part{
		{Text: string(s[:region]), Name: regionName, Stem: true},
		{Text: string(s[region:letter]), Name: nameOf(numberNameOf(string(s[region:letter]))), Stem: region == 0},
		{Text: string(s[letter : letter+1]), Name: Name{Kind: plateKind, Form: string(s[letter : letter+1])}},
		{Text: string(s[letter+1:]), Name: nameOf(numberNameOf(string(s[letter+1:])))},
	}
	if region == 0 {
		parts = parts[1:]
	}
	w.Parts = append(parts, w.Parts[1:]...)
	return w
}
