//ff:func feature=cli type=command control=iteration dimension=1
//ff:what TestNewRootCmd — hypo, parse and token (in cobra's name order) are the only subcommands besides cobra's help; the completion command is off
package cli

import (
	"strings"
	"testing"
)

func TestNewRootCmd(t *testing.T) {
	root := newRootCmd()
	root.InitDefaultCompletionCmd()
	var names []string
	for _, c := range root.Commands() {
		names = append(names, c.Name())
	}
	if root.Name() != "nlgeul" || strings.Join(names, " ") != "hypo parse token" {
		t.Fatal(root.Name(), names)
	}
}
