//ff:func feature=token type=helper control=sequence
//ff:what ontoSentences — made-up sentences with a sign of a pair at the head of a unit with letters written onto it (a closing sign and a particle, a closing sign and letters, signs of a code span or of a quotation, a title mark, a sign after the letters), for the net read beside the old reading of such a unit
package token

func ontoSentences() []string {
	return []string{
		"물(고양이)을 마신다",
		"(고양이)가 마신다",
		")를 마신다",
		")를.",
		"(주)고양이가마신다.",
		"「고양이」를 마신다",
		"작은 (고양이)를, 마신다",
		"(고양이)고양이가마당에서물을마신다",
		"(고양이)·사과를 먹는다",
		"(고양이)",
		"\"고양이\"가 마당에 있다",
	}
}
