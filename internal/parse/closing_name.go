//ff:func feature=parse type=reader control=selection
//ff:what closingName — a closing's name with its lead written out: ㄴ다, ㄹ까, ㅂ니다, 어서
package parse

func closingName(e ending) string {
	switch e.lead {
	case jongN:
		return "ㄴ" + e.form
	case jongL:
		return "ㄹ" + e.form
	case jongM:
		return "ㅁ" + e.form
	case jongB:
		return "ㅂ" + e.form
	case leadA:
		return "어" + e.form
	}
	return e.form
}
