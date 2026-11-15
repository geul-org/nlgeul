//ff:func feature=token type=helper control=sequence
//ff:what nominalEnd — does a word end in a nominal: its last part is a nominal stem (nominalStem), an order (제1) or a unit after a number (3차, 1990년)
package token

func nominalEnd(parts []Part) bool {
	if len(parts) == 0 {
		return false
	}
	last := parts[len(parts)-1]
	if last.Stem {
		return nominalStem(last) || last.Name.Kind == nameOf(orderName).Kind
	}
	return last.Name.Kind == nameOf(unitLabel).Kind
}
