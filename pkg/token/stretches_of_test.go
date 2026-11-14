//ff:func feature=token type=helper control=sequence
//ff:what TestStretchesOf — contests read the old way become stretches with their branches and no cells, in order; none stay none
package token

import "testing"

func TestStretchesOf(t *testing.T) {
	got := stretchesOf([]Contest{{Text: "가"}, {Text: "나다"}})
	if len(got) != 2 || got[0].contest.Text != "가" || got[1].contest.Text != "나다" || got[0].cells != nil || got[1].cells != nil {
		t.Fatal(got)
	}
	if len(stretchesOf(nil)) != 0 {
		t.Fatal("none")
	}
}
