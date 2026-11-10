//ff:func feature=mini type=reader control=sequence
//ff:what TestSlotsMark — each slot is a piece; no slots, nothing
package mini

import "testing"

func TestSlotsMark(t *testing.T) {
	if slotsMark([]rune("었겠더")) != "+었+겠+더" || slotsMark(nil) != "" {
		t.Fatal("slotsMark")
	}
}
