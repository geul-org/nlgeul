//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what shapeNames — the names the table of shapes (special-forms.tsv) gives a letter, one for each of its lines (（: 전각:기호); none for a letter that table lacks, whatever the other tables of special letters call it
package parse

func shapeNames(r rune) []string {
	var out []string
	for _, f := range forms[r] {
		if f.origin == madeTable+".tsv" {
			out = append(out, f.name())
		}
	}
	return out
}
