//ff:func feature=token type=rule control=iteration dimension=1
//ff:what TestAddNouns — 테 gets one graph (줄임:테) with three exceptions, counted as the copula and kept out of the total; the noun with the copula's stem in it stands right after a word that ends in a modifying closing before a closing the copula takes (니까, ㄴ데 under it), and not where no such word stands before it, before a closing the copula does not take, or before 어/아 or ㅆ fused into its own vowel, each turned down with its reason
package token

import (
	"slices"
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestAddNouns(t *testing.T) {
	ix := &index{byKey: map[rune][]*entry{}, words: map[string]*entry{}, closed: map[string]int{}}
	ix.addClosings()
	ix.addSlots()
	graphs, edges, total := ix.graphs, ix.edges, ix.total
	ix.addNouns()
	if ix.graphs != graphs+1 || ix.edges != edges+3 || ix.total != total || len(ix.nouns) != len(nounCopulas) {
		t.Fatal(ix.graphs-graphs, ix.edges-edges, ix.total-total, len(ix.nouns))
	}
	te := ix.nouns[0]
	if te.name != "줄임:테" || te.form != "테" || te.count != copulaCount || te.g == nil {
		t.Fatal(te)
	}
	judged := func(p rule.Piece) (float64, []string) { return verdictOf(te.g, p, "x", true) }
	if v, _ := judged(rule.Piece{Form: "테", HasStem: true, Adn: true, After: "맺음:니까"}); v <= 0 {
		t.Fatal("테 before 니까 right after a modifying closing", v)
	}
	if v, _ := judged(rule.Piece{Form: "테", HasStem: true, Adn: true, After: "맺음:ㄴ데"}); v <= 0 {
		t.Fatal("테 with ㄴ데 fused under it right after a modifying closing", v)
	}
	if v, why := judged(rule.Piece{Form: "테", HasStem: true, After: "맺음:니까"}); v > 0 || !slices.Contains(why, whyNoAdn) {
		t.Fatal("테 where no word ending in a modifying closing stands before it", v, why)
	}
	fused := ix.fusedNames()
	for _, name := range []string{"맺음:어(합침)", "선어말:ㅆ"} {
		if v, why := judged(rule.Piece{Form: "테", HasStem: true, Adn: true, After: name}); !slices.Contains(fused, name) || v > 0 || !slices.Contains(why, whyFusedIn) {
			t.Error(name, "fused into the vowel of 테", v, why)
		}
	}
	if v, why := judged(rule.Piece{Form: "테", HasStem: true, Adn: true, After: "맺음:자"}); !slices.Contains(ix.barred, "맺음:자") || v > 0 || !slices.Contains(why, whyBarred) {
		t.Fatal("테 before a closing the copula does not take", v, why)
	}
}
