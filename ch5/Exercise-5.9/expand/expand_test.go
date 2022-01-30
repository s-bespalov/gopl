package expand

import (
	"fmt"
	"testing"
	"time"
)

func TestExpand(t *testing.T) {
	utime := time.Now().Unix()
	s := "Hello %s. How are you, %s. Unix time is %s"
	tmplt := fmt.Sprintf(s, "$name", "$name", "$time")
	expected := fmt.Sprintf(s, "Mark", "Mark", fmt.Sprint(utime))
	r := Expand(tmplt, func(s string) string {
		if s == "time" {
			return fmt.Sprint(utime)
		}
		if s == "name" {
			return "Mark"
		}
		return "$" + s
	})
	if r != expected {
		t.Errorf("should be equal:\n%s\n%s\n", expected, r)
	}
}
