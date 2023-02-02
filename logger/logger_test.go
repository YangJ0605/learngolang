package logger

import "testing"

func TestConstLevel(t *testing.T) {
	t.Logf("%v %T\n", Debug, Debug)
}
