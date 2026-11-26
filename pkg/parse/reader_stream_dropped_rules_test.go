//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestReaderStreamDroppedRules — over the small word table for streams, a closing right after a run of Latin letters and digits (xyz라면) or right after a closing sign written onto it ('가'다) is read first with the copula's dropped stem, the closing read right on the run behind it, and with no rules of a sentence (st.rules nil) the closing read on the run or the sign alone comes first; a closing 고 right onto a closing sign is read with no stem, and 면 right after a number is the noun, not the copula's
package parse

import "testing"

func TestReaderStreamDroppedRules(t *testing.T) {
	for _, c := range []struct{ sentence, text, first string }{
		{"이것은 xyz라면 된다.", "xyz라면", "xyz[열:영숫자]+(이)[서술격:이(빠짐)]+라면[맺음:라면]"},
		{"이것은 '가'다.", "다", "+(이)[서술격:이(빠짐)]+다[맺음:다]"},
		{"이것은 '가'고 된다.", "고", "+고[맺음:고]"},
	} {
		got, _ := copulaBranches(t, c.sentence, c.text, true)
		if len(got) == 0 || got[0].String() != c.first {
			t.Errorf("%s: the first candidate is %v, want %s", c.sentence, got, c.first)
		}
	}
	plain, _ := copulaBranches(t, "이것은 xyz라면 된다.", "xyz라면", false)
	if len(plain) == 0 || plain[0].String() != "xyz[열:영숫자]+라면[맺음:라면]" {
		t.Errorf("with no rules the closing on the run comes first, got %v", plain)
	}
}
