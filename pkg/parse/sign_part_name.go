//ff:func feature=parse type=helper control=sequence
//ff:what signPartName — the name of a part that stands for a sign's letters: the whole name as its kind, except the plain number name, which numberNameOf makes from the letters
package parse

func signPartName(name, text string) Name {
	if name == numName {
		return nameOf(numberNameOf(text))
	}
	return Name{Kind: name}
}
