//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what formShaped — is r a letter the tables of special letters give the shape (꼴) asked for (Ⅲ is of the shape 로마 숫자)
package parse

func formShaped(r rune, shape string) bool {
	for _, f := range forms[r] {
		if f.shape == shape {
			return true
		}
	}
	return false
}
