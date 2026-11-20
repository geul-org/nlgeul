//ff:func feature=parse type=reader control=sequence
//ff:what nounSuffixOf — one suffix read after a nominal, from a line of its table: its form and what must stand before it; it is a particle named 접미:<form>, in the place of a suffix in chains (rule.LinkSuffix)
package parse

import (
	"github.com/geul-org/nlgeul/pkg/data"
	"github.com/geul-org/nlgeul/pkg/rule"
)

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
	return entry{name: "접미:" + form, form: form, class: classParticle, after: after, slot: rule.LinkSuffix}, nil
}
