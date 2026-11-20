//ff:type feature=parse type=model
//ff:what alt — one candidate of a changed stem form: the stem of the word table it stands for (in changed: the changed form itself), and the kind of change (irr…)
package parse

type alt struct {
	stem string
	kind int
}
