//ff:func feature=token type=reader control=iteration dimension=1
//ff:what signBundle — one stretch of one sign in a row as a bundle: a reading for each graph of the sign that stands (a heading is named with how many signs stand in a row, and holds that number as its depth); the letters with no name when none stands, and beside the readings when all of them read the sign by what it does for the words after it (roleNames)
package token

import (
	"slices"
	"strconv"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func (sc *scan) signBundle(text string, m rule.Mark, end int) Bundle {
	b := Bundle{Text: text}
	marks := sc.judgeMark(sc.ix.marks[[]rune(text)[0]], m, text)
	role := true
	for _, mk := range marks {
		name := mk.name
		p := Part{Text: text, Name: Name{Kind: name}}
		if name == headingName {
			p.Name.Kind, p.Depth = name+strconv.Itoa(m.Run), m.Run
		}
		role = role && slices.Contains(roleNames, name)
		b.Branches = append(b.Branches, signBranch(p, end))
	}
	if role {
		b.Branches = append(b.Branches, signBranch(Part{Text: text, Stem: true}, end))
	}
	return b
}
