//ff:func feature=parse type=reader control=sequence
//ff:what TestFullStems — only the stems under the longest tails: 된다 gives 되, not 된
package parse

import "testing"

func TestFullStems(t *testing.T) {
	got := fullStems(readWord("된다", true, &sentence{}))
	if len(got) != 1 || got[0] != "되" {
		t.Fatal(got)
	}
}
