package main

import (
	"github.com/rezonphilip/gopherDS/invoke"
)

func main() {
	invoke.TestRecursionFactorial()
	invoke.TestBranchingRecursionFibonacci()
	invoke.TestRecursionSumOfN(10)
	invoke.TestRecursionExponent(3, 5)
	invoke.TestRecursiveReverse("rezon")

	invoke.TestRecursiveGCD(12, 18)

	invoke.TestRecursiveBinarySearch()
	invoke.TestRecursiveDirectoryTraversal("/Users/rezon/Desktop")
}
