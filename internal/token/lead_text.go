//ff:func feature=token type=reader control=selection
//ff:what leadText — a lead written out for a graph's name: ㄴ, ㄹ, ㅁ, ㅂ, ㅆ, 어
package token

func leadText(lead int) string {
	switch lead {
	case 0:
		return ""
	case leadVowel:
		return "어"
	}
	return string(jamoKeys[lead])
}
