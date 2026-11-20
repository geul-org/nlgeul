//ff:type feature=parse type=model
//ff:what gap — a boundary of a candidate seen by the spacing rules: the two words that meet there (left, right) and how many letters each takes up
package parse

type gap struct {
	left     Eojeol
	right    Eojeol
	leftLen  int
	rightLen int
}
