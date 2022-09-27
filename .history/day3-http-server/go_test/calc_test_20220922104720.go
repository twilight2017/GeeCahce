package main

import "testing"

func TestAdd(t *testing.T) {
	if ans := Add(1, 2); ans != 3 {
		t.Errorf("1+2 expected be 3 , but %d got", ans)
	}
	if ans := Mul(2, 3); ans != 6 {
		t.Errorf("2*3 expected be 6, but %d got", ans)
	}
}

func TestMul(t *testing.T) {
	t.Run("pos", func(t *testing.T) {
		if Mul(2, 3) != 6 {
			t.Fatal("fail")
		}
	})

	t.Run("neg", func(t *testing.T) {
		if Mul(1, 1) != 1 {
			t.Fatal("fail")
		}
	})
}

func TestMul2(t *testing.T){
	cases := [] struct{
		Name string
		A, B, Expected, int
	}{
		{"pos", 2, 3, 6},
		{"neg", 1, 1, 1},
		{"zero", 6, 0, 0},
		{"neg2", -1, -9, 9},
	}

	for _,c := range(cases){
		t,Run(c.Name, func(t *testing.T){
            if ans :=Mul(c.A, c.B);ans != c.Expected{
				t.Fatal("fail")
			}
		})
	}
}
