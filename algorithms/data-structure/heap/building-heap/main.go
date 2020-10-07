package main

import "fmt"

func main() {
	// ランダムな列
	a := []int{5, 8, 2, 20, 1, 3, 9, 6, 22, 15, 8, 10}
	fmt.Println(a)
	buildHeap(a)
	fmt.Println(a)

}

func buildHeap(a []int) {
	n := len(a) - 1
	for i := n/2 - 1; i >= 0; i-- {
		downHeap(a, i)
	}
}

func downHeap(a []int, i int) {
	n := len(a) - 1
	l := left(i)
	r := right(i)
	largest := i

	if l < n && a[l] > a[i] { // 親(自分)、左の子、右の子の仲で最大のノードを見つける
		largest = l
	} else {
		largest = i
	}

	if r < n && a[r] > a[largest] {
		largest = r
	}

	if largest != i { // どちらかの子が最大の場合
		swap(&a[i], &a[largest])
		downHeap(a, largest) // 再帰によってダウンヒープを繰り返す
	}
}

func swap(a, b *int) {
	var t int
	t = *a
	*a = *b
	*b = t
}

func left(n int) int {
	return n*2 + 1
}

func right(n int) int {
	return n*2 + 2
}
