//ff:func feature=token type=reader control=sequence
//ff:what TestClosedEnding — the closed-set words that end with a form, in order; not the form itself, nothing for an empty form
package token

import "testing"

func TestClosedEnding(t *testing.T) {
	words := map[string]*entry{"그리고": {}, "고": {}, "하지만": {}, "먹고": {}}
	got := closedEnding(words, "고")
	if len(got) != 2 || got[0] != "그리고" || got[1] != "먹고" || closedEnding(words, "") != nil {
		t.Fatal(got)
	}
}
