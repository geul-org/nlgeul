//ff:func feature=token type=helper control=iteration dimension=1
//ff:what lastOneWord — the last of some branches that holds one word, printed (Branch.String); "" when none does
package token

func lastOneWord(bs []Branch) string {
	last := ""
	for _, b := range bs {
		if len(b.Words) == 1 {
			last = b.String()
		}
	}
	return last
}
