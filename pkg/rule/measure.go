//ff:func feature=rule type=helper control=selection
//ff:what measure — the length a side names: the particle chain's, the tail reading's under judgment, or the word's longest tail's
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func measure(ctx toulmin.Context, side string) int {
	switch side {
	case SideChain:
		return WordOf(ctx).Chain.Size
	case SideTail:
		return TailOf(ctx).Size
	}
	return WordOf(ctx).Longest
}
