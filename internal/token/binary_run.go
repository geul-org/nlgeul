//ff:func feature=token type=helper control=sequence
//ff:what binaryRun — is a text made of the digits 0 and 1 only, however short
package token

import "strings"

func binaryRun(text string) bool {
	return text != "" && strings.Trim(text, "01") == ""
}
