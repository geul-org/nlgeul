//ff:func feature=parse type=reader control=sequence
//ff:what matchParticle — where p starts if s[:end] ends with it, something comes before it, and that syllable meets p's condition
package parse

import "github.com/geul-org/nlgeul/pkg/rule"

func matchParticle(s []rune, end int, p particle) (int, bool) {
	start, ok := suffixAt(s, end, p.form)
	return start, ok && start > 0 && rule.Fits(s[start-1], p.after)
}
