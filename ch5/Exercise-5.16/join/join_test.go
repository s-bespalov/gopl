package join

import (
	"strings"
	"testing"
)

func TestJoin(t *testing.T) {
	arr := []string{"race", "vote", "lock", "bus", "alter", "polar", "elegant", "mirror", "laptop", "ring", "husband", "tent", "parade", "spoil", "eager"}
	exp := strings.Join(arr, "-")
	res := join("-", arr...)
	if exp != res {
		t.Fail()
	}
}
