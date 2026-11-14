//ff:func feature=token type=helper control=iteration dimension=1
//ff:what numeralReach — how many letters back from a place a numeral of the number table needs to be looked for: the letters of the longest native numeral (열 and one of the ones, or the longest of the others), or one; a Sino-Korean numeral run ends in a digit or a greater number, and one letter of it, the last, is already one (sinoTokens of the last letter is not 0 whenever it is of a longer run)
package token

func numeralReach() int {
	reach := 1
	for _, text := range nativeAlone {
		reach = max(reach, len([]rune(text)))
	}
	for _, text := range nativeOnes {
		reach = max(reach, 1+len([]rune(text)))
	}
	return reach
}
