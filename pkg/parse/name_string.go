//ff:func feature=parse type=helper control=sequence
//ff:what String — a candidate printed: kind:form, or the kind alone when the form is empty
package parse

func (n Name) String() string {
	if n.Form == "" {
		return n.Kind
	}
	return n.Kind + ":" + n.Form
}
