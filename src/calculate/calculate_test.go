package calculate

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	tt := []struct {
		A, B, Expected float64
	}{
		{A: 1, B: 1, Expected: 2},
		{A: 1, B: 2, Expected: 3},
		{A: 2, B: 1, Expected: 3},
		{A: 2, B: 2, Expected: 4},
	}

	for i, x := range tt {
		t.Run(fmt.Sprintf("test (%d)", i), func(st *testing.T) {
			got := Add(x.A, x.B)
			if got != x.Expected {
				st.Errorf("expected %v, got %v", x.Expected, got)
			}
		})
	}
}

func TestSubstract(t *testing.T) {
	tt := []struct {
		A, B, Expected float64
	}{
		{A: 1, B: 1, Expected: 0},
		{A: 1, B: 2, Expected: -1},
		{A: 2, B: 1, Expected: 1},
		{A: 2, B: 2, Expected: 0},
	}

	for i, x := range tt {
		t.Run(fmt.Sprintf("test (%d)", i), func(st *testing.T) {
			got := Substract(x.A, x.B)
			if got != x.Expected {
				st.Errorf("expected %v, got %v", x.Expected, got)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tt := []struct {
		A, B, Expected float64
	}{
		{A: 1, B: 1, Expected: 1},
		{A: 1, B: 2, Expected: 2},
		{A: 2, B: 1, Expected: 2},
		{A: 2, B: 2, Expected: 4},
	}

	for i, x := range tt {
		t.Run(fmt.Sprintf("test (%d)", i), func(st *testing.T) {
			got := Multiply(x.A, x.B)
			if got != x.Expected {
				st.Errorf("expected %v, got %v", x.Expected, got)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tt := []struct {
		A, B, Expected float64
	}{
		{A: 1, B: 1, Expected: 1},
		{A: 1, B: 2, Expected: 0.5},
		{A: 2, B: 1, Expected: 2},
		{A: 2, B: 2, Expected: 1},
	}

	for i, x := range tt {
		t.Run(fmt.Sprintf("test (%d)", i), func(st *testing.T) {
			got := Divide(x.A, x.B)
			if got != x.Expected {
				st.Errorf("expected %v, got %v", x.Expected, got)
			}
		})
	}
}

func TestRound(t *testing.T) {
	tt := []struct {
		F, Expected float64
	}{
		{1.009, 1.01},
		{1.019, 1.02},
		{1.011, 1.01},
		{1, 1},
	}

	for i, x := range tt {
		t.Run(fmt.Sprintf("test (%d)", i), func(st *testing.T) {
			got, _ := round(x.F)
			if got != x.Expected {
				st.Errorf("expected %v, got %v", x.Expected, got)
			}
		})
	}
}
