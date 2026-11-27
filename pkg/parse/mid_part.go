//ff:func feature=parse type=helper control=sequence
//ff:what midPart — the part that prints one candidate of a sign read between letters: named as the candidate is (signPartName), a stem when it has no name; a candidate with no name of a sign the tables of special letters name (bareName: =) is a stem with that name (=[기호:같음])
package parse

func midPart(text, name string, sign rune) Part {
	if kept := bareName(name, sign); keptName(sign, kept) {
		return Part{Text: text, Name: nameOf(kept), Stem: true}
	}
	return Part{Text: text, Name: signPartName(name, text), Stem: name == ""}
}
