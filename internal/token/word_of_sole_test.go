//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestWordOfSole — with the word table of testdata a one-syllable verb stem and its closing that make up the whole unit (하자, 가는) stand fully and are printed so (성립 1.00); the same word after another word in its unit (입력하자 read as 입력 하+자, 학교가는 as 학교 가+는), or before another (하고먹자 read from 하+고), stands as the stem does when it is not the whole of its unit, below 1, and the reading of 입력 하+자 or of 학교 가+는 is printed below 1
package token

import (
	"strings"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestWordOfSole(t *testing.T) {
	ja := tail{start: 1, pieces: []piece{{"자", "맺음:자", 1, -1, rule.TailVerb, false, ""}}}
	neun := tail{start: 1, pieces: []piece{{"는", "꾸밈:는", 1, -1, rule.TailVerb, false, ""}}}
	if e := wordsScan("하자", layerCost).wordOf(0, 2, ja); e.sig != "하[열쇠:동사]+자[맺음:자]" || e.verdict != 1 {
		t.Fatal("a one-syllable stem and its closing as the whole unit", e)
	}
	if e := wordsScan("가는", layerCost).wordOf(0, 2, neun); e.sig != "가[열쇠:동사]+는[꾸밈:는]" || e.verdict != 1 {
		t.Fatal("a one-syllable stem and its modifying closing as the whole unit", e)
	}
	ws := wordsScan("입력하자", layerCost)
	_, ha := ws.ix.known("하", rule.TailVerb, false)
	if e := ws.wordOf(2, 4, tail{start: 3, pieces: ja.pieces}); e.sig != "하[열쇠:동사]+자[맺음:자]" || e.verdict != ha || ha >= 1 || ha <= 0 {
		t.Fatal("the same word after another word in its unit", e, ha)
	}
	ws = wordsScan("학교가는", layerCost)
	_, ga := ws.ix.known("가", rule.TailVerb, false)
	if e := ws.wordOf(2, 4, tail{start: 3, pieces: neun.pieces}); e.sig != "가[열쇠:동사]+는[꾸밈:는]" || e.verdict != ga || ga >= 1 || ga <= 0 {
		t.Fatal("the same word after another word in its unit", e, ga)
	}
	ws = wordsScan("하고먹자", layerCost)
	_, link := ws.ix.known("하", rule.TailPred, false)
	if e := ws.wordOf(0, 2, tail{start: 1, pieces: []piece{{"고", "맺음:고", 1, -1, rule.TailPred, false, ""}}}); e.sig != "하[열쇠:동사]+고[맺음:고]" || e.verdict != link || link >= 1 {
		t.Fatal("the same stem at the unit's head with another word after it", e, link)
	}
	r, _ := New(Options{Words: "testdata/words.tsv"})
	got := strings.Join(r.Lines("하자 가는 입력하자 학교가는", false), "\n")
	for _, want := range []string{
		"  하[열쇠:동사]+자[맺음:자]\t성립 1.00 ",
		"  가[열쇠:동사]+는[꾸밈:는]\t성립 1.00 ",
		"  입력[열쇠:체언] 하[열쇠:동사]+자[맺음:자]\t성립 0.60 ",
		"  학교[열쇠:체언] 가[열쇠:동사]+는[꾸밈:는]\t성립 0.60 ",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("missing %q in\n%s", want, got)
		}
	}
}
