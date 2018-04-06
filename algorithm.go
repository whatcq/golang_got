package main

import "fmt"

func fibonacci() func() int {
	var sum int
	prev, next := 0, 1
	count := 0
	return func() int {
		count++
		if count < 2 {
			return count
		}
		sum, prev = prev+next, next
		next = sum
		//一并赋值，先取出后面的值，然后逐个赋值！写成先这样是bug
		//sum, prev, next = prev+next, next, sum
		return sum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
	var a, b, c int
	a, b, c = i(1), i(2), i(3)
	fmt.Println(a, b, c)
}

func i(i int) int {
	fmt.Println(i)
	return i
}
