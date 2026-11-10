//ff:func feature=cli type=command control=iteration dimension=1
//ff:what TestNewRootCmd — parse and token are the only subcommands besides cobra's help; the completion command is off
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
	if root.Name() != "nlgeul" || strings.Join(names, " ") != "parse token" {
		t.Fatal(root.Name(), names)
	}
}
