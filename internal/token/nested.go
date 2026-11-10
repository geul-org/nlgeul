//ff:func feature=token type=helper control=iteration dimension=1
//ff:what nested — the round brackets of a sentence that open and close a pair, by nesting, each place with its partner's (both ways); a bracket with no partner, and one at a place to skip, are left out
package token

func nested(r []rune, skip []bool) map[int]int {
	out := map[int]int{}
	var stack []int
	for i, c := range r {
		if skip[i] || (c != '(' && c != ')') {
			continue
		}
		if c == '(' {
			stack = append(stack, i)
			continue
		}
		if n := len(stack); n > 0 {
			out[stack[n-1]], out[i] = i, stack[n-1]
			stack = stack[:n-1]
		}
	}
	return out
}
