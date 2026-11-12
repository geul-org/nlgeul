//ff:func feature=rule type=helper control=selection
//ff:what listedFact — the fact a list is asked for (List…), read from the piece, the sign, the word or the tail under judgment
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func listedFact(ctx toulmin.Context, of int) string {
	switch of {
	case ListForm:
		return PieceOf(ctx).Form
	case ListBefore:
		return PieceOf(ctx).Before
	case ListAfter:
		return PieceOf(ctx).After
	case ListBase:
		return PieceOf(ctx).Base
	case ListExt:
		return MarkOf(ctx).Ext
	case ListLast:
		return MarkOf(ctx).Last
	case ListWord:
		return WordOf(ctx).Text
	}
	return TailOf(ctx).Closing
}
