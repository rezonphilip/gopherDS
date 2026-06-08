package invoke

import (
	"fmt"
	"math/rand"
	"time"

	hashmap "github.com/rezonphilip/gopherDS/HashMap"
	bst "github.com/rezonphilip/gopherDS/Tree/BST"
	"github.com/rezonphilip/gopherDS/recursion"
)

func TestBasicHashMap() {
	basicHashMap := hashmap.NewBasicHashMap(10)

	for n := range 100 {
		basicHashMap.Put(fmt.Sprintf("key-%d", n), n)
	}

	fmt.Printf("the capacity of the hash is now %d\n", basicHashMap.Size)

	for i := range 10 {
		fmt.Printf("attempt_%d\n", i)
		key := fmt.Sprintf("key-%d", rand.Intn(151))
		fmt.Printf("attempting to get value for %s\n", key)
		val, ok := basicHashMap.Get(key)
		if ok {
			fmt.Printf("got -> %v\n", val)
		} else {
			fmt.Println("didnt get the value from the hash")
		}
	}
}

func TestBST() {
	bst := bst.NewBST()

	for range 100 {
		bst.Insert(rand.Intn(100))
	}

	bst.Inorder()
	fmt.Println()

	bst.Preorder()
	fmt.Println()

	bst.Postorder()
	fmt.Println()

	for i := range 10 {
		fmt.Printf("attempt_%d\n", i)
		searchVal := rand.Intn(151)
		fmt.Printf("attempting to get value for %d\n", searchVal)
		ok := bst.Search(searchVal)
		if ok {
			fmt.Printf("got -> %v\n", searchVal)
		} else {
			fmt.Println("didnt get the value from the bst")
		}
	}
}

func TestGoRoutines() {
	for range 100000 {
		go func() {
			time.Sleep(time.Second * 100)
		}()
	}
}

func TestRecursionFactorial() {

	for i := range 10 {
		fmt.Printf("the factorial of %d, is %d\n", i, recursion.Factorial(i))
	}
}

func TestBranchingRecursionFibonacci() {
	for i := range 10 {
		fmt.Printf("%d ", recursion.Fibonacci(i))
	}
	fmt.Println()
}

func TestRecursionSumOfN(n int) {
	fmt.Println(recursion.Sum(n))
}

func TestRecursionExponent(base, exp int) {
	fmt.Println(recursion.Exponent(base, exp))
}

func TestRecursiveReverse(str string) {
	fmt.Println(recursion.Reverse(str))
}
