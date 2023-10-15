package factorial

import "testing"

func TestFactorial(t *testing.T) {

	i := factorial(5)

	if i != 120 {
		t.Fail()
	}
}
