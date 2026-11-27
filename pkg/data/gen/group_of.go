//ff:func feature=data type=helper control=selection
//ff:what groupOf — the group and name of a symbol, by the first rule that fits: money, emoji, box drawing, arrow with its direction, math, shape, or the rest
package main

import (
	"slices"
	"strings"
)

func groupOf(t *ucdTable, r rune, block string) (string, string) {
	cat := categoryOf(t, r)
	name := t.Chars[r].Name
	switch {
	case cat == "Sc":
		return "기호", "돈"
	case slices.Contains(emojiBlocks, block):
		return "기호", "이모지"
	case slices.Contains(boxBlocks, block):
		return "기호", "상자 그림"
	case slices.Contains(arrowBlocks, block):
		if d := arrowDirection(name); d != "" && strings.Contains(name, "ARROW") {
			return "화살표", d
		}
		return "기호", "화살표"
	case slices.Contains(mathBlocks, block), cat == "Sm" && slices.Contains(commonBlocks, block):
		return "기호", "수학"
	case slices.Contains(shapeBlocks, block):
		return "기호", "도형"
	}
	return "기호", "그 밖"
}
