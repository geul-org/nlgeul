//ff:type feature=parse type=model
//ff:what numeral — a number before a unit that is not a run of digits: its letters, its preference and what it is named (the plain number name, which the letters then tell the form of, or the name its reading gives it)
package parse

type numeral struct {
	text string
	pref float64
	name string
}
