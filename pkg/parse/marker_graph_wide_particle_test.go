//ff:func feature=parse type=reader control=sequence
//ff:what TestMarkerGraphWideParticle — a particle read only with a word table (the shortened 선) is turned down where a word is written right onto it and stands at its unit's end and before a particle; in the stream 여기선계약 has no way with 여기+선 before 계약, and 여기선 계약 written apart keeps it
package parse

import (
	"slices"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestMarkerGraphWideParticle(t *testing.T) {
	e := entry{name: "부사격:선(에서는)", form: "선", class: classParticle, wide: true}
	markerGraph(&e, nil, nil, true)
	at := func(right int) (float64, []string) {
		return verdictOf(e.g, rule.Piece{Form: "선", HasStem: true, Prev: '기', Next: -1, Right: right}, "여기선", true)
	}
	if v, why := at(rule.RightOther); v > 0 || !slices.Contains(why, whyRight) {
		t.Fatal("a word written right onto it", v, why)
	}
	if v, _ := at(rule.RightDone); v <= 0 {
		t.Fatal("at the unit's end", v)
	}
	if v, _ := at(rule.RightParticle); v <= 0 {
		t.Fatal("before a particle the chain judges", v)
	}
	r := streamReader(t, "words-joined.tsv")
	const glued = "여기[열쇠:체언]+선[부사격:선(에서는)] 계약[열쇠:체언]"
	if lines := branchLines(r, "여기선계약"); lines[glued] || !lines["여기[열쇠:체언] 선계약"] {
		t.Fatal("선 before a word written onto it", lines)
	}
	if !branchLines(r, "여기선 계약")["여기[열쇠:체언]+선[부사격:선(에서는)]"] {
		t.Fatal("written apart", branchLines(r, "여기선 계약"))
	}
}
