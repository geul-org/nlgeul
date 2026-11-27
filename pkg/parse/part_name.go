//ff:func feature=parse type=helper control=sequence
//ff:what partName — the name of a part from the name a piece or stem is given and the letters of the part: the plain number name is made by numberNameOf from the letters, any other name is as nameOf reads it
package parse

func partName(name, text string) Name {
	if name == numName {
		return signPartName(name, text)
	}
	return nameOf(name)
}
