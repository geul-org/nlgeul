//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what keptName — is a name one that the hand-written table of special letters or the table of kinds gives a sign (special-etc.tsv, special-groups.tsv: ○ 가림:원, * 기호:별표, → 화살표:오른쪽), and not one of the table of shapes (ⓐ 원:로마자): a sign so named is read on as a sign with no name was, a stem that stands where a nominal stands (○○ for a name kept back)
package parse

func keptName(sign rune, name string) bool {
	for _, f := range forms[sign] {
		if f.origin != madeTable+".tsv" && f.name() == name {
			return true
		}
	}
	return false
}
