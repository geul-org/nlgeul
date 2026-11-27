//ff:func feature=data type=helper control=selection
//ff:what runeKind — what one character of a value is: 수, 한글, 가나, 한자, 로마자, 그리스 글자, 그 밖의 글자 or 기호
package main

import "strings"

func runeKind(t *ucdTable, r rune) string {
	cat := categoryOf(t, r)
	name := t.Chars[r].Name
	switch {
	case cat == "Nd":
		return "수"
	case r >= 0x1100 && r <= 0x11FF, r >= 0x3130 && r <= 0x318F, r >= 0xAC00 && r <= 0xD7A3:
		return "한글"
	case r >= 0x3040 && r <= 0x30FF:
		return "가나"
	case r >= 0x4E00 && r <= 0x9FFF, r >= 0x3400 && r <= 0x4DBF, r >= 0x20000 && cat == "Lo":
		return "한자"
	case cat[0] == 'L' && strings.Contains(name, "LATIN"):
		return "로마자"
	case cat[0] == 'L' && strings.Contains(name, "GREEK"):
		return "그리스 글자"
	case cat[0] == 'L':
		return "그 밖의 글자"
	}
	return "기호"
}
