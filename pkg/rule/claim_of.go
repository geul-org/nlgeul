//ff:func feature=rule type=helper control=sequence
//ff:what ClaimOf — the claim name and the conclusion written in an evaluation result's name by the node's Claim criterion ("" and "" when it carries none)
package rule

import "regexp"

var claimIn = regexp.MustCompile(`claim\{"Name":"([^"]*)","Concludes":"([^"]*)"\}`)

func ClaimOf(result string) (string, string) {
	m := claimIn.FindStringSubmatch(result)
	if m == nil {
		return "", ""
	}
	return m[1], m[2]
}
