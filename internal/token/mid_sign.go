//ff:type feature=token type=model
//ff:what midSign — a sign between letters of a unit that is also read as a word of its own: the parts that print it, one for each reading that stands (a middle dot is one reading)
package token

type midSign struct {
	reads []Part
}
