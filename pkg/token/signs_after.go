//ff:func feature=token type=reader control=sequence
//ff:what signsAfter — the letters of a unit that stand after the letters that are read (the signs trimmed from its end); the core is looked for after the unit's head (a sign of a pair or a span of code), as a sign dropped between the head and the core leaves the two apart in the unit's letters; "" when nothing is read
package token

import "strings"

func (sc *scan) signsAfter(field string) string {
	if len(sc.s) == 0 {
		return ""
	}
	head, body := string(sc.s[:sc.head]), string(sc.s[sc.head:])
	return field[len(head)+strings.Index(field[len(head):], body)+len(body):]
}
