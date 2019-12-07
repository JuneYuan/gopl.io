package geometry

import (
	"testing"
)

func Test_methodRule(t *testing.T) {
	path := Path{{0, 0}, {1, 0}}
	t.Log(path.Distance())
}
