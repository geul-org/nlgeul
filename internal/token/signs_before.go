//ff:func feature=token type=helper control=iteration dimension=1
//ff:what signsBefore — where the run of signs right before a place begins, of those a name may begin with (~ / . - @ # $ _); the place itself when none stands there
package token

import "strings"

func signsBefore(all []rune, from int) int {
	for from > 0 && strings.ContainsRune("~/.-@#$_", all[from-1]) {
		from--
	}
	return from
}
