//ff:func feature=parse type=helper control=sequence
//ff:what name — what a letter of the tables of special letters is called in the output: its shape and what it writes, as the table gives them (원:수, 괄호:한글)
package parse

func (f form) name() string {
	return f.shape + ":" + f.content
}
