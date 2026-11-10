//ff:func feature=token type=reader control=selection
//ff:what nextUnit — the unit that begins at place i of a sentence and where it ends: a span of code with what is written onto it; a sign that opens a pair, alone; a sign that closes a pair with what is written onto it; else the letters up to the next space or sign of a pair
package token

func nextUnit(r []rune, i int, pairs map[int]int) (unit, int) {
	j, paired := pairs[i]
	switch {
	case paired && r[i] == '`':
		end := stretchEnd(r, j+1, pairs)
		return unit{string(r[i:end]), j + 1 - i, codeName}, end
	case paired && j > i:
		return unit{string(r[i : i+1]), 1, "여는 " + pairNames[r[i]]}, i + 1
	case paired:
		end := stretchEnd(r, i+1, pairs)
		return unit{string(r[i:end]), 1, "닫는 " + pairNames[r[i]]}, end
	}
	end := stretchEnd(r, i+1, pairs)
	return unit{text: string(r[i:end])}, end
}
