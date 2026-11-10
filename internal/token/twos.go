//ff:func feature=token type=helper control=iteration dimension=1
//ff:what twos — the places of one sign in a sentence paired two by two in the order they come, each with its partner's (both ways); a last one with no partner is left out, and so are the places to skip and an apostrophe between two letters
package token

func twos(r []rune, sign rune, skip []bool) map[int]int {
	out := map[int]int{}
	open := -1
	for i, c := range r {
		if c != sign || (skip != nil && skip[i]) || (sign == '\'' && lettered(r, i)) {
			continue
		}
		if open < 0 {
			open = i
			continue
		}
		out[open], out[i], open = i, open, -1
	}
	return out
}
