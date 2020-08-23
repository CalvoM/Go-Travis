package math

import (
	"testing"
)

func TestPower(t *testing.T) {
	actual := 8
	expected := power(2, 3)
	if actual != expected {
		t.Fail()
	}
	actual = 9
	expected = power(3, 2)
	if actual != expected {
		t.Fail()
	}
	actual = 81
	expected = power(3, 4)
	if actual != expected {
		t.Fail()
	}
}
