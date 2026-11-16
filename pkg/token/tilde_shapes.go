//ff:func feature=token type=helper control=iteration dimension=1
//ff:what tildeShapes — the shapes of the signs between runs of digits with the tilde and each of its variants (tildeSigns) added as a range of two runs (범위:~, 범위:∼, …)
package token

import "github.com/geul-org/nlgeul/pkg/rule"

func tildeShapes(shapes map[rune][]rule.Shape) map[rune][]rule.Shape {
	for _, r := range tildeSigns {
		shapes[r] = append(shapes[r], rule.Shape{Of: "범위:" + string(r), Min: 2, Max: 2})
	}
	return shapes
}
