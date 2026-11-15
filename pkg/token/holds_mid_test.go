//ff:func feature=token type=helper control=sequence
//ff:what holdsMid — whether any candidate of the contests still prints the place name of a sign between letters (기호), which is only held while a unit is read
package token

import (
	"strings"
)

func holdsMid(contests []Contest) bool {
	return strings.Contains(strings.Join(contestLines(contests), "\n"), "["+midName)
}
