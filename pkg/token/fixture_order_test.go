//ff:func feature=token type=helper control=iteration dimension=1
//ff:what missingLine — the first of some line starts that is not found, in the order given, among printed lines (each start is looked for after the line the one before was found in); "" when all are found
package token

import "strings"

func missingLine(lines, starts []string) string {
	at := 0
	for _, start := range starts {
		for at < len(lines) && !strings.HasPrefix(lines[at], start) {
			at++
		}
		if at == len(lines) {
			return start
		}
		at++
	}
	return ""
}
