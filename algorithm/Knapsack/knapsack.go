package knapsack

type Knapsack struct {
	// w      int   // 背包能够承受的重量
	// wt     []int //item的重量
	// values []int //item 的价值
	// n      int   // item的个数
}

func (k *Knapsack) knapsack(w int, wt []int, values []int, n int) int {
	if n == 0 || w == 0 {
		return 0
	}

	if wt[n-1] > w {
		return k.knapsack(w, wt, values, n-1)
	} else {
		return max(values[n-1]+k.knapsack(w-wt[n-1], wt, values, n-1), k.knapsack(w, wt, values, n-1))
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
