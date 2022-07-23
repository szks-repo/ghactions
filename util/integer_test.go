package util

import "testing"

func TestIntegerValue(t *testing.T) {
	t.Run("", func(t *testing.T) {
		x := int(123)
		if v := IntegerValue(&x); v != int(123) {
			t.Error("error")
		}
	})
}
