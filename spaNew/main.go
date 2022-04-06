package main

// libs to print and some math calculations
import (
	"fmt"
	"math"
)

// main function to drive program
func main() {
	// fixed size array, that will be PERMUTED
	myarray := [5]int{1, 2, 3, 4, 5}

	// temporary variables to make backtracking
	// bactracking - the classic method to solve tasks like permutations, sudokos etc.
	// it's very helpful when we need to iterate through all possible options
	// in this case, variable visited declared to assign that certain element in the array were "visited", so we need to change next variable
	var visited [len(myarray)]bool
	// app array need to be temporary array that keeps values that were appended
	var app []int
	// for futher calculation, like what part of permutations it is
	num := 0

	// calling function rec, that is recursively calls itself till the end
	rec(myarray[:], visited[:], app, &num)

	// we need find, x1 * x2 * x3 what part of all the permutation
	// so, to find the number of all possible permutations we just need to find factorial of it
	totalPermutation := Factorial(len(myarray))

	// printing percenatge
	fmt.Println(num)
	fmt.Printf("%f", float64(num)/float64(totalPermutation)*float64(100))
	fmt.Println()

}

// backtracking function

func rec(myarray []int, visited []bool, app []int, num *int) {
	// count needed to keep value
	count := 1

	// condition to check whether appended array is filled already, if it is we go to check the numbers of permutations
	// that meets condition
	if len(myarray) == len(app) {
		// loop to sum up all numbers in the permutation
		for i := 0; i < len(app)-1; i++ {
			count *= app[i] - app[i+1]
		}
		// checking if the total sum is smaller than sqrt(N!)
		// if it is, global variable "num" will be increamented by +1
		if float64(count) < math.Sqrt(float64(Factorial(len(myarray)))) {
			*num++
		}

		return
	}

	// main backtraking snippet
	// loop to go through all array

	for i := 0; i < len(myarray); i++ {
		// if the certain element were visited, it skips untill find that were not visited
		if !visited[i] {
			// after finding element that were not visited, function sets value to true
			visited[i] = true
			// goes to rec funciton again
			rec(myarray, visited, append(app, myarray[i]), num)
			// if initially all elements were visited, it returns from current RECURSIVE FUNCTION and sets value to false
			visited[i] = false
		}
	}
}

// recursively founded Factorial fucntion

func Factorial(n int) (result int) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}
