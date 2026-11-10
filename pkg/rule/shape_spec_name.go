//ff:func feature=rule type=model control=sequence
//ff:what SpecName — the criterion's name: "shape:" and the reading it speaks of
package rule

func (s *Shape) SpecName() string {
	return "shape:" + s.Of
}
