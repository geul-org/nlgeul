//ff:func feature=token type=helper control=iteration dimension=1
//ff:what holdsMid — whether any candidate of the contests still prints the place name of a sign between letters (기호), which is only held while a unit is read
package token

import (
	"strings"
)

func holdsMid(contests []Contest) bool {
	for _, b := range contests {
		if strings.Contains(strings.Join(render(b, 0), "\n"), "["+midName) {
			return true
		}
	}
	return false
}
