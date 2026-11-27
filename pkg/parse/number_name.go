//ff:func feature=parse type=helper control=sequence
//ff:what numberName — is a name the name of a number: the kind 수 whatever its form, and the names a shape gives (원문자:수, 괄호문자:수, 특수:로마자)
package parse

import "slices"

func numberName(n Name) bool {
	return n.Kind == numName || slices.Contains(shapeNumberNames, n.String())
}
