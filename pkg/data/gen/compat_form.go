//ff:func feature=data type=helper control=selection
//ff:what compatForm — the form of a <compat> decomposition and its value with the brackets or the dot or comma taken off
package main

func compatForm(v []rune, block string) (string, []rune) {
	n := len(v)
	switch {
	case n > 2 && v[0] == '(' && v[n-1] == ')':
		return "괄호", v[1 : n-1]
	case n > 2 && v[0] == '〔' && v[n-1] == '〕':
		return "거북등 괄호", v[1 : n-1]
	case n > 1 && v[n-1] == '.' && allDigits(string(v[:n-1])):
		return "점 붙음", v[:n-1]
	case n > 1 && v[n-1] == ',' && allDigits(string(v[:n-1])):
		return "쉼표 붙음", v[:n-1]
	case block == "Number Forms":
		return "로마 숫자", v
	case block == "Hangul Compatibility Jamo":
		return "호환 자모", v
	}
	return "그 밖의 호환", v
}
