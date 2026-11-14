//ff:func feature=token type=helper control=iteration dimension=1
//ff:what bodyLen — how many letters a word takes up without its particles: the letters it takes up (taken) less the letters of its particle parts
package token

func bodyLen(e Eojeol, taken int) int {
	for _, p := range tailParts(e) {
		if particlePart(p) {
			taken -= len([]rune(p.Text))
		}
	}
	return taken
}
