package array

import "fmt"

type Array struct {
	arr   []int
	realL int
	fakeL int
}

func New() *Array {
	return &Array{
		arr:   make([]int, 0),
		realL: 0,
		fakeL: 0,
	}
}

func (a *Array) Get(i int) int {
	if i < 0 || i > a.realL {
		return -1
	}
	return a.arr[i]
}

func (a *Array) Set(i, v int) bool {
	for idx := range a.arr {
		if idx == i {
			a.arr[i] = v
			return true
		}
	}

	return false
}

func (a *Array) Add(v int) {
	if a.realL == 0 {
		a.realL++
		newarr := make([]int, a.realL)
		newarr[0] = v
		a.arr = newarr
	} else if a.fakeL+1 >= a.realL {
		a.realL *= 2
		newarr := make([]int, a.realL)
		for i := 0; i < a.fakeL; i++ {
			newarr[i] = a.arr[i]
		}
		newarr[a.fakeL] = v
		a.arr = newarr
	} else {
		a.arr[a.fakeL] = v
	}
	a.fakeL++
}

func (a *Array) Remove(v int) {
	// [1, 2, 3, 4] -> remove 3
	// 1 is 3, no, so add it as its current index, which is zero
	// 2 is 3, no, so add it as its current index, which is one
	// 3 is 3, yes, dont add it to the array, and dont update the current position, which is two
	// 4 is 3, no, so add it as the position left, which is two
	newarr := make([]int, a.fakeL-1)
	pos := 0
	for i := 0; i < a.fakeL; i++ {
		if a.arr[i] == v {
			continue
		} else {
			newarr[pos] = a.arr[i]
			pos++
		}
	}

	a.fakeL--
	a.realL = a.fakeL
	a.arr = newarr
}

func (a *Array) Print() {
	fmt.Println("[")
	for i := 0; i < a.realL; i++ {
		fmt.Printf("%d: %d\n", i, a.arr[i])
	}
	fmt.Println("]")
}
