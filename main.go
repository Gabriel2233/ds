package main

import (
	"fmt"

	"github.com/Gabriel2233/ds/array"
)

func main() {
	arr := array.New()

	arr.Add(2)
	arr.Add(5)
	arr.Add(10)
	arr.Add(9)

	arr.Remove(2)

	arr.Add(50)

	fmt.Println(arr.Get(2))
	fmt.Println(arr.Get(10))

	arr.Set(0, 100)

	arr.Print()
}
