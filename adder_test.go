package adder

import "testing"

func TestAdd(t *testing.T) {
	//per default no assertions possible - you can use testify import for that
	result := Add(1, 2)
	if result != 3 {
		t.Fatalf("FATAL: got %v, wanted %v", result, 3)
	}

	//table driven tests:
	var cases = []struct {
		a int
		b int
		r int
	}{
		{a: 1, b: 2, r: 3},
		{a: 2, b: 2, r: 4},
		{a: 5, b: 11, r: 16},
		{a: 2, b: 2, r: 44},
	}
	for _, c := range cases {
		r := Add(c.a, c.b)
		if r != c.r {
			t.Fatalf("FATAL: got %v, wanted %v", r, c.r)
		}
	}
}
