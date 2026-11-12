//ff:func feature=token type=helper control=sequence
//ff:what noteOf — what a trace says of a graph's verdict: 통과 when it stands, else 기각 with the reasons of the exceptions that applied
package token

import "strings"

func noteOf(v float64, why []string) string {
	if v > 0 {
		return "통과"
	}
	return "기각: " + strings.Join(why, ", ")
}
