//ff:func feature=parse type=reader control=sequence
//ff:what nounSuffixOf — one suffix read after a nominal, from a line of its table: its form, what must stand before it and whether particles follow it; it is a particle named 접미:<form>, in the place of a suffix in chains (rule.LinkSuffix) or of a suffix that ends its word (rule.LinkLast)
package parse

import "github.com/geul-org/nlgeul/pkg/data"

func nounSuffixOf(row data.Row) (entry, error) {
	lead, form, err := FormOf(row, 0)
	if err != nil {
		return entry{}, err
	}
	if lead != 0 || form == "" {
		return entry{}, row.Errf(0, "접미의 꼴은 한글 음절로만 적는다")
	}
	after, err := row.Pick(1, AfterNames)
	if err != nil {
		return entry{}, err
	}
	slot, err := row.Pick(2, SuffixNames)
	if err != nil {
		return entry{}, err
	}
	return entry{name: "접미:" + form, form: form, class: classParticle, after: after, slot: slot}, nil
}
