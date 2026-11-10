//ff:func feature=token type=reader control=iteration dimension=1
//ff:what runsOf — the runs of Latin letters, digits and underscores in a unit, each as where it starts and ends
package token

import "github.com/geul-org/koparse/pkg/rule"

func runsOf(s []rune) [][2]int {
	var out [][2]int
	start := -1
	for i := 0; i <= len(s); i++ {
		in := i < len(s) && rule.RunRune(s[i])
		if in && start < 0 {
			start = i
		}
		if !in && start >= 0 {
			out, start = append(out, [2]int{start, i}), -1
		}
	}
	return out
}
