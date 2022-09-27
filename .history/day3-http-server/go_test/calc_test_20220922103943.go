package go_test

import "testing"

func TestAdd(t *testing.T) {
	if ans := Add(1, 2); ans != 3 {
		t.Errorf("1+2 expected be 3 , but %d got", ans)
	}
	if ans := Mul(2, 3); ans != 6 {
		t.Errorf("2*3 expected be 6, but %d got", ans)
	}
}
