package main

import (
	"fmt"
	"math"
)

func main() {
	myarray := [5]int{1, 2, 3, 4, 5}

	var visited [len(myarray)]bool
	var app []int
	num := 0

	rec(myarray[:], visited[:], app, &num)

	totalPermutation := Factorial(len(myarray))

	fmt.Println(num)
	fmt.Printf("%f", float64(num)/float64(totalPermutation)*float64(100))
	fmt.Println()

}

func rec(myarray []int, visited []bool, app []int, num *int) {
	count := 0

	if len(myarray) == len(app) {
		for i := 0; i < len(app)-1; i++ {
			// in this task I changed calculations and condition
			// the rest is the same as in previous
			count += app[i] - app[i+1]
		}
		if float64(count) > math.Pow(float64(len(myarray)), 2)/2 {
			*num++
		}

		return
	}

	for i := 0; i < len(myarray); i++ {
		if !visited[i] {
			visited[i] = true
			rec(myarray, visited, append(app, myarray[i]), num)
			visited[i] = false
		}
	}
}

func Factorial(n int) (result int) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}
