//ff:func feature=token type=reader control=sequence
//ff:what TestCoreOf — a plain unit loses the signs at its edges; a unit with a head keeps it and what follows loses its own
package token

import "testing"

func TestCoreOf(t *testing.T) {
	if string(coreOf(unit{text: "ㆍ곰이ㆍ"})) != "곰이" || string(coreOf(unit{text: ")를ㆍ", head: 1})) != ")를" {
		t.Fatal("a named sign that is a letter for Unicode")
	}
	if string(coreOf(unit{text: "곰이,"})) != "곰이" || string(coreOf(unit{text: ")를,", head: 1})) != ")를" || string(coreOf(unit{text: "(", head: 1})) != "(" {
		t.Fatal(string(coreOf(unit{text: ")를,", head: 1})))
	}
}
