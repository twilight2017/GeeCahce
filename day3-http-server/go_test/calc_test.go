package main

import (
	"fmt"
	"os"
	"testing"
)

// func TestAdd(t *testing.T) {
// 	if ans := Add(1, 2); ans != 3 {
// 		t.Errorf("1+2 expected be 3 , but %d got", ans)
// 	}
// 	if ans := Mul(2, 3); ans != 6 {
// 		t.Errorf("2*3 expected be 6, but %d got", ans)
// 	}
// }

// func TestMul(t *testing.T) {
// 	t.Run("pos", func(t *testing.T) {
// 		if Mul(2, 3) != 6 {
// 			t.Fatal("fail")
// 		}
// 	})

// 	t.Run("neg", func(t *testing.T) {
// 		if Mul(1, 1) != 1 {
// 			t.Fatal("fail")
// 		}
// 	})
// }

// func TestMul2(t *testing.T) {
// 	cases := []struct {
// 		Name           string
// 		A, B, Expected int
// 	}{
// 		{"pos", 2, 3, 6},
// 		{"neg", 1, 1, 1},
// 		{"zero", 6, 0, 0},
// 		{"neg2", -1, -9, 9},
// 	}

// 	for _, c := range cases {
// 		t.Run(c.Name, func(t *testing.T) {
// 			if ans := Mul(c.A, c.B); ans != c.Expected {
// 				t.Fatal("fail")
// 			}
// 		})
// 	}
// }

// type calcCase struct{ A, B, Expected int }

// func createMulTestCase(t *testing.T, c *calcCase) {
// 	//t.Helper()
// 	if ans := Mul(c.A, c.B); ans != c.Expected {
// 		t.Fatal("fail")
// 	}
// }

// func TestMul3(t *testing.T) {
// 	createMulTestCase(t, &calcCase{2, 5, 10})
// 	createMulTestCase(t, &calcCase{-2, 9, -18})
// 	createMulTestCase(t, &calcCase{0, 8, 0})
// }
func setup() {
	fmt.Println("Before all tests")
}

func teardown() {
	fmt.Println("after all tests")
}

func Test1(t *testing.T) {
	fmt.Println("this is test1")
}

func Test2(t *testing.T) {
	fmt.Println("this is test2")
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)

}
