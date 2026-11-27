//ff:func feature=rule type=helper control=selection
//ff:what listedFact — the fact a list is asked for (List…), read from the piece, the sign, the word or the tail under judgment, or from the candidate judged in its sentence and the words around it
package rule

import (
	"strings"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

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
	case ListSelfStem:
		return AroundOf(ctx).Self.Stem
	case ListSelfKind:
		return AroundOf(ctx).Self.Kind
	case ListSelfLast:
		return AroundOf(ctx).Self.Last
	case ListPrevLast:
		return nearAt(AroundOf(ctx).Before, 0).Last
	case ListNextStem:
		return nearAt(AroundOf(ctx).After, 0).Stem
	case ListNextLast:
		return nearAt(AroundOf(ctx).After, 0).Last
	case ListNext2Stem:
		return nearAt(AroundOf(ctx).After, 1).Stem
	case ListPrevKind:
		kind, _, _ := strings.Cut(nearAt(AroundOf(ctx).Before, 0).Kind, ":")
		return kind
	case ListPrevStem:
		return nearAt(AroundOf(ctx).Before, 0).Stem
	case ListNextKind:
		kind, _, _ := strings.Cut(nearAt(AroundOf(ctx).After, 0).Kind, ":")
		return kind
	case ListSelfText:
		return AroundOf(ctx).Self.Text
	case ListPrevName:
		return nearAt(AroundOf(ctx).Before, 0).Kind
	case ListNextName:
		return nearAt(AroundOf(ctx).After, 0).Kind
	case ListNext2Last:
		return nearAt(AroundOf(ctx).After, 1).Last
	}
	return TailOf(ctx).Closing
}
