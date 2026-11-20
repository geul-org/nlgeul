//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what TestLastOfAny — 지, 냐 and 가, the last letters of the closings that ask (askClosings), are such; the other letters of their names and the last letters of other closings are not
package parse

import "testing"

func TestLastOfAny(t *testing.T) {
	for _, r := range []rune{'지', '냐', '가'} {
		if !lastOfAny(askClosings, r) {
			t.Errorf("%c", r)
		}
	}
	for _, r := range []rune{'는', '은', '을', 'ㄴ', 'ㄹ', ':', '기', '고', '다', 'x', '를'} {
		if lastOfAny(askClosings, r) {
			t.Errorf("%c", r)
		}
	}
}
