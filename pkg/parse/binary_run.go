//ff:func feature=parse type=helper control=sequence
//ff:what binaryRun — is a text made of the digits 0 and 1 only, however short
package parse

import "strings"

func binaryRun(text string) bool {
	return text != "" && strings.Trim(text, "01") == ""
}
