//ff:func feature=parse type=rule control=sequence
//ff:what listed — the criterion of rule.Listed: a named list and the fact looked up in it (rule.List…)
package parse

import "github.com/geul-org/nlgeul/pkg/rule"

func listed(name string, of int, items []string) *rule.List {
	return &rule.List{Name: name, Of: of, Items: items}
}
