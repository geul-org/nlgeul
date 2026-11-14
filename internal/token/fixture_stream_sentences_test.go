//ff:func feature=token type=helper control=sequence
//ff:what streamSentences — the sentences the stream tests read, made up for them or the well known one: words written apart, a unit with no space, a number, signs of a pair, a closing dot
package token

func streamSentences() []string {
	return []string{
		"아버지가방에들어가신다",
		"아버지가방에들어가신다.",
		"아버지 가방에 들어가신다",
		"아버지가 방에 들어가신다",
		"돌아가셨다",
		"작은 고양이가 물을 마신다.",
		"우리는 사과 3개를 먹어보다.",
		"고양이(물)을 ~/a/b를 했다.",
		"어떤 학교에 갔다...",
	}
}
