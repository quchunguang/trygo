package main

import (
	"fmt"
)

func main() {
	fmt.Println("# Algorithms in Go")

	// FastSearch(10)
	// FastMerge(10)
	// Bernoulli()
	// ClosePoint()
	// SingleLink()
	// MapReduce()
	// Permutation()
	// GenPrime(100)
	// MatchRegexp()
	// SingleRing()
	// Searching()
	// FastSort1()
	// PostfixCalc(Infix2Postfix("2+(10+(2*4))+1"))
	// PrefixCalc("+ * 2 3 3 ")
	// fmt.Println(Gcd(314159, 271828))
	// fmt.Println(Puzzle(3))
	// fmt.Println(Max([]int{6, 7, 8, 9, 1, 2, 3, 4, 5}, 0, 8))
	// Hanoi(5, 1)
	// fmt.Println("F(92) = ", F(92))
	// fmt.Println(Knapsack(17))
	// fmt.Printf("%c", GenMax("VABCDE").Item.(byte))

	// With input "KEVIN", the tree would be,
	//     V
	//   V   N
	//  K V I N
	// K E
	root := GenMax("KEVIN")
	// BTTraverseFront(root)
	// BTTraverseMiddle(root)
	// BTTraverseLast(root)
	BTTraverseLevel(root)
	// BTTraverseFrontNR(root)
}
