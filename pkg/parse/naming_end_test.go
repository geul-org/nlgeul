//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestNamingEnd — 기, 음 and ㅁ, the last letters of the naming closings a particle chain is read right after (chainHeads), are ones, and so is a syllable with ㅁ fused under it (됨, 함, 임, 름) or ㅁ under ㄻ (삶); the other letters of the closings' names, letters of other closings and syllables under another final consonant are not
package parse

import "testing"

func TestNamingEnd(t *testing.T) {
	for _, r := range []rune{'기', '음', 'ㅁ', '됨', '함', '임', '름', '삶'} {
		if !namingEnd(r) {
			t.Errorf("%c", r)
		}
	}
	for _, r := range []rune{':', '이', '다', '지', '을', '먹', '읍', 'x'} {
		if namingEnd(r) {
			t.Errorf("%c", r)
		}
	}
}
