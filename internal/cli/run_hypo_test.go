//ff:func feature=cli type=command control=iteration dimension=1
//ff:what TestRunHypo — the count first, then a line for every hypothesis ranked from 1, the first being the first reading of every bundle; an empty sentence has none; a sentence past hypo.Limit is only counted and refused
package cli

import (
	"bytes"
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/geul-org/nlgeul/internal/hypo"
	"github.com/geul-org/nlgeul/internal/token"
)

func TestRunHypo(t *testing.T) {
	r, err := token.New(token.Options{Words: "testdata/words.tsv"})
	if err != nil {
		t.Fatal(err)
	}
	sentence := "우리는 마당에서 사과를 먹었다."
	choices := hypo.Choices(r.Bundles(sentence))
	count := hypo.Count(choices)
	var firsts []string
	for _, b := range choices {
		firsts = append(firsts, b.Branches[0].String())
	}

	var out bytes.Buffer
	if err := runHypo(&out, r, sentence); err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(strings.TrimSuffix(out.String(), "\n"), "\n")
	if count.Int64() < 2 || lines[0] != "가설 "+count.String() || int64(len(lines)) != count.Int64()+1 {
		t.Fatalf("count %s, %d lines: %q", count, len(lines), out.String())
	}
	if top := strings.SplitN(lines[1], "\t", 3); len(top) != 3 || top[0] != "1" || !strings.HasPrefix(top[1], "성립 ") || top[2] != strings.Join(firsts, " | ") {
		t.Fatalf("first %q, want the first readings %q", lines[1], strings.Join(firsts, " | "))
	}
	for i, l := range lines[1:] {
		if !strings.HasPrefix(l, fmt.Sprintf("%d\t성립 ", i+1)) {
			t.Fatalf("line %d: %q", i+1, l)
		}
	}

	out.Reset()
	if err := runHypo(&out, r, ""); err != nil || out.String() != "가설 0\n" {
		t.Fatalf("empty sentence: %v %q", err, out.String())
	}

	bare, err := token.New(token.Options{})
	if err != nil {
		t.Fatal(err)
	}
	many := strings.TrimSpace(strings.Repeat("작은 ", 13))
	over := hypo.Count(hypo.Choices(bare.Bundles(many)))
	if over.Cmp(big.NewInt(hypo.Limit)) <= 0 {
		t.Fatalf("%s hypotheses are not past the limit", over)
	}
	out.Reset()
	if err := runHypo(&out, bare, many); err == nil || out.String() != "가설 "+over.String()+"\n" {
		t.Fatalf("past the limit: %v %q", err, out.String())
	}
}
