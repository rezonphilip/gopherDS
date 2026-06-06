package main

import (
	"fmt"
	"math/rand"

	hashmap "github.com/rezonphilip/gopherDS/HashMap"
	bst "github.com/rezonphilip/gopherDS/Tree/BST"
)

func main() {
	testBST()
}

func testBasicHashMap() {
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

func testBST() {
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
