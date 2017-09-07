package go_sat_solver

import "testing"


func TestChildrenCount(t *testing.T) {
	cases := []struct {
		expr Expr
		want int
	}{
		{And{[]Expr{Symbol{"x1"}, Symbol{"x2"}}}, 2},
		{Or{[]Expr{Symbol{"x1"}, Symbol{"x2"}, Symbol{"x3"}}}, 3},
		{Not{[1]Expr{Symbol{"x1"}}}, 1},
		{Symbol{"x1"}, 0},
		{Literal{true}, 0},
		{Literal{false}, 0},
	}
	for _, c := range cases {
		got := len(c.expr.Children())
		if got != c.want {
			t.Errorf("len(Children(%#v)) == %v, want %v", c.expr, got, c.want)
		}
	}
}
