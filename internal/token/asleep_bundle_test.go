//ff:func feature=token type=reader control=sequence
//ff:what TestAsleepBundle — no readings asleep give no bundle; readings asleep are one bundle with the unit's letters and each reading whole as a branch marked asleep, with the reading's own verdict and preference
package token

import (
	"reflect"
	"testing"
)

func TestAsleepBundle(t *testing.T) {
	sc := &scan{s: []rune("고양이")}
	if got := sc.asleepBundle(nil); got != nil {
		t.Fatalf("no readings asleep: %+v", got)
	}
	gw := Eojeol{Parts: []Part{{Text: "고양", Stem: true}}, End: 2}
	iw := Eojeol{Parts: []Part{{Text: "이", Stem: true}}, End: 3}
	whole := Eojeol{Parts: []Part{{Text: "고양이", Stem: true}}, End: 3}
	root := &path{verdict: 1}
	first := &path{prev: root, last: step{end: 2, word: gw, verdict: 0.9, pref: -1}, verdict: 0.9, pref: -1}
	split := &path{prev: first, last: step{end: 3, word: iw, verdict: 0.8, pref: -2}, verdict: 0.25, pref: -7.5}
	one := &path{prev: root, last: step{end: 3, word: whole, verdict: 0.7, pref: -4}, verdict: 0.5, pref: -6}
	got := sc.asleepBundle([]*path{split, one})
	want := []Bundle{{
		Text: "고양이",
		Branches: []Branch{
			{Words: []Eojeol{gw, iw}, Verdict: 0.25, Pref: -7.5, Asleep: true},
			{Words: []Eojeol{whole}, Verdict: 0.5, Pref: -6, Asleep: true},
		},
	}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %+v\nwant %+v", got, want)
	}
}
