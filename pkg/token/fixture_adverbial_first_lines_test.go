//ff:func feature=token type=helper control=iteration dimension=1
//ff:what adverbialFirstLines — the lines of a trace that the rule 에서·서는 부사격 먼저 wrote (those with its name), in order
package token

import "strings"

func adverbialFirstLines(notes []string) []string {
	var out []string
	for _, n := range notes {
		if strings.Contains(n, "\t"+ruleAdverbialFirst+"\t") {
			out = append(out, n)
		}
	}
	return out
}
