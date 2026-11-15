//ff:type feature=rule type=model
//ff:what Around — a candidate of a word judged with the words of its sentence around it: the candidate itself, the words before it (the nearest first) and the words after it up to the predicate (the nearest first), each as its first candidate reads it
package rule

type Around struct {
	Self   Near
	Before []Near
	After  []Near
}
