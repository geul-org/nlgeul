//ff:func feature=parse type=helper control=sequence
//ff:what TestAddFlagged — a word is put at the end of the list of every name whose bit is set, and of no other; lists already there keep their words
package parse

import (
	"strings"
	"testing"
)

func TestAddFlagged(t *testing.T) {
	out := map[string][]string{"가": {"앞"}}
	addFlagged(out, "말", 5, map[string]int{"가": 1, "나": 2, "다": 4})
	if len(out) != 2 || strings.Join(out["가"], ",") != "앞,말" || strings.Join(out["다"], ",") != "말" || out["나"] != nil {
		t.Fatal(out)
	}
}
