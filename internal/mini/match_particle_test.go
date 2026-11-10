//ff:func feature=mini type=reader control=sequence
//ff:what TestMatchParticle — 을 after a consonant is read, after a vowel or with nothing before it is not
package mini

import "testing"

func TestMatchParticle(t *testing.T) {
	p := particle{"을", afterBatchim, 3, caseObj, 1}
	if start, ok := matchParticle([]rune("문장을"), 3, p); !ok || start != 2 {
		t.Fatal(start, ok)
	}
	if _, ok := matchParticle([]rune("나무을"), 3, p); ok {
		t.Fatal("after a vowel")
	}
	if _, ok := matchParticle([]rune("을"), 1, p); ok {
		t.Fatal("nothing before")
	}
	if _, ok := matchParticle([]rune("문장이"), 3, p); ok {
		t.Fatal("another syllable")
	}
}
