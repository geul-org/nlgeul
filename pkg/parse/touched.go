//ff:func feature=parse type=helper control=sequence
//ff:what touched — whether a sign other than itself that is no punctuation mark of a sentence (textMarks) stands right before or right after a stretch of one sign, from start to stop in some letters (=>, |-, ★☆); a sign of the unit table read with the number before it is none (5%↓), nor is a punctuation mark of a sentence (○○·△△, ○○,)
package parse

import "strings"

func touched(all []rune, start, stop int) bool {
	before := start > 0 && signRune(all[start-1]) && !unitSign(all, start-1) && !strings.ContainsRune(textMarks, all[start-1])
	after := stop < len(all) && signRune(all[stop]) && !strings.ContainsRune(textMarks, all[stop])
	return before || after
}
