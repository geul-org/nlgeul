//ff:func feature=hypo type=helper control=iteration dimension=1
//ff:what TestCount — the product of the numbers of readings; none when there is no bundle; a product past int64 is exact
package hypo

import (
	"math/big"
	"testing"

	"github.com/geul-org/nlgeul/internal/token"
)

func TestCount(t *testing.T) {
	small := []token.Bundle{
		{Branches: make([]token.Branch, 2)},
		{Branches: make([]token.Branch, 3)},
		{Branches: make([]token.Branch, 1)},
		{Branches: make([]token.Branch, 4)},
	}
	if got := Count(small); got.Cmp(big.NewInt(24)) != 0 {
		t.Fatalf("count %s, want 24", got)
	}
	if got := Count(nil); got.Sign() != 0 {
		t.Fatalf("count of no bundle %s, want 0", got)
	}

	var large []token.Bundle
	for i := 0; i < 40; i++ {
		large = append(large, token.Bundle{Branches: make([]token.Branch, 7)})
	}
	want := new(big.Int).Exp(big.NewInt(7), big.NewInt(40), nil)
	got := Count(large)
	if got.Cmp(want) != 0 || got.IsInt64() || got.String() != "6366805760909027985741435139224001" {
		t.Fatalf("count %s, want %s (past int64)", got, want)
	}
}
