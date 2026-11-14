//ff:func feature=token type=helper control=iteration dimension=1
//ff:what lineKey — a way of cutting written as one string, its spans in order (0-3,3-6), to tell the ways apart
package token

import "strconv"

func lineKey(spans []Span) string {
	key := ""
	for i, s := range spans {
		if i > 0 {
			key += ","
		}
		key += strconv.Itoa(s.Start) + "-" + strconv.Itoa(s.End)
	}
	return key
}
