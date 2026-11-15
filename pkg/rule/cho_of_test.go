//ff:func feature=rule type=helper control=sequence
//ff:what TestChoOf — 가 begins with ㄱ(0), 낮 with ㄴ(2), 잎 with ㅇ(11), 힣 with ㅎ(18), a Latin letter and the code point after 힣 are not Hangul
package rule

import "testing"

func TestChoOf(t *testing.T) {
	if ChoOf('가') != 0 || ChoOf('낮') != 2 || ChoOf('잎') != 11 || ChoOf('힣') != 18 || ChoOf('A') != -1 || ChoOf(0xD7A4) != -1 {
		t.Fatal("ChoOf")
	}
}
