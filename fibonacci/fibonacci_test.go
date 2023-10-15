package fibonacci

import "testing"

func TestFibonacci(t *testing.T) {

	i := Fibonacci(9)

	if i != 34 {
		t.Fail()
	}
}
