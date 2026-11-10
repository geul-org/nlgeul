//ff:func feature=rule type=helper control=iteration dimension=1
//ff:what TestLinks — 에+는, 에서+부터, 만+으로, 까지+가, 와+의, 들+이, 는+요 are chains; 과+가, 에+서(부사격·주격), 이+는, 는+도, 야 뒤, 요 뒤 are not
package rule

import "testing"

func TestLinks(t *testing.T) {
	cases := []struct {
		left, right int
		want        bool
	}{
		{LinkAdv, LinkTopic, true}, {LinkAdv, LinkRange, true}, {LinkRange, LinkAdv, true}, {LinkRange, LinkCase, true},
		{LinkWith, LinkAdn, true}, {LinkSuffix, LinkCase, true}, {LinkTopic, LinkYo, true}, {LinkAdv, LinkComp, true},
		{LinkAdv, LinkDir, true}, {LinkHonor, LinkTopic, true}, {LinkDir, LinkTopic, true},
		{LinkDir, LinkAdv, false}, {LinkHonor, LinkCase, false},
		{LinkWith, LinkCase, false}, {LinkAdv, LinkAdv, false}, {LinkAdv, LinkCase, false}, {LinkCase, LinkTopic, false},
		{LinkTopic, LinkTopic, false}, {LinkVoc, LinkYo, false}, {LinkYo, LinkTopic, false},
	}
	for _, c := range cases {
		if Links(c.left, c.right) != c.want {
			t.Errorf("Links(%d, %d)", c.left, c.right)
		}
	}
}
