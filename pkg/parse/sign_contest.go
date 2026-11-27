//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what signContest — one stretch of one sign in a row as a contest: a candidate for each graph of the sign that stands, those of a higher verdict first and each with its graph's verdict (a heading is named with how many signs stand in a row, and holds that number as its depth); after them a candidate for each name the tables of special letters give the sign by its shape (ix.forms: ⓐ[원:로마자], ㈎[괄호:한글]); the letters with no name when none stands, and beside the candidates when all of them read the sign by what it does for the words after it (roleNames)
package parse

import (
	"cmp"
	"slices"
	"strconv"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func (sc *scan) signContest(text string, m rule.Mark, end int) Contest {
	b := Contest{Text: text}
	marks := sc.judgeMark(sc.ix.marks[[]rune(text)[0]], m, text)
	role := true
	slices.SortStableFunc(marks, func(a, b markCandidate) int { return cmp.Compare(b.v, a.v) })
	for _, mk := range marks {
		name := mk.name
		p := Part{Text: text, Name: signPartName(name, text)}
		if name == headingName {
			p.Name.Kind, p.Depth = name+strconv.Itoa(m.Run), m.Run
		}
		role = role && slices.Contains(roleNames, name)
		b.Branches = append(b.Branches, lowered(signBranch(p, end), mk.v))
	}
	for _, mk := range sc.judgeMark(sc.ix.forms[[]rune(text)[0]], m, text) {
		role = false
		b.Branches = append(b.Branches, lowered(signBranch(Part{Text: text, Name: nameOf(mk.name)}, end), mk.v))
	}
	if role {
		b.Branches = append(b.Branches, signBranch(Part{Text: text, Stem: true}, end))
	}
	return b
}
