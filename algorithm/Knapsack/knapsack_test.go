package knapsack

import "testing"

func TestKnapsack(t *testing.T) {
	val := []int{60, 100, 120}
	wt := []int{10, 20, 30}
	w := 50
	n := len(val)
	k := &Knapsack{}
	res := k.knapsack(w, wt, val, n)
	if res != 220 {
		t.Fail()
	}
}
