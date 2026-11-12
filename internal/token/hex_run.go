//ff:func feature=token type=helper control=sequence
//ff:what hexRun — is a run of Latin letters and digits 0x and then one hexadecimal digit or more (0x1234, 0xFF); the prefix is the lower case 0x only
package token

import "strings"

func hexRun(label string) bool {
	rest, ok := strings.CutPrefix(label, hexLead)
	return ok && rest != "" && strings.Trim(rest, hexDigit) == ""
}
