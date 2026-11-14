//ff:func feature=token type=helper control=sequence
//ff:what netSentences — made-up sentences (and 아버지가방에들어가신다) for the net against the old reading: written apart, written together, with signs, numbers, signs of a pair, signs between letters, and a sign at an edge (ontoSentences are those with a sign of a pair with letters written onto it)
package token

func netSentences() []string {
	return []string{
		"아버지가방에들어가신다",
		"아버지 가방에 들어가신다",
		"고양이가 마당에서 물을 마신다",
		"고양이가마당에서물을마신다",
		"작은 고양이가 사과를 먹는다.",
		"작은고양이가사과를먹는다.",
		"고양이/마당 물·사과",
		"고양이/마당에서물·사과를먹는다",
		"1,000개를 먹었다",
		"10:30에 마당에서 3\" 물을 마셨다",
		"3\"",
		"그렇다. 그러나 모른다?",
		"서로서로 하나둘 어디서든",
		"foo-bar of foo/of",
		"-3 --xx ~/a/b",
		"입력과출력의문제",
		"예쁘다지우개켜서듣고돕고살고",
		"갑돌씨가 반장 겸 총무 및 운영",
		"물물물물물물물물물물",
	}
}
