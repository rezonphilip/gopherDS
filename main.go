package main

import (
	"fmt"
	"math/rand"

	hashmap "github.com/rezonphilip/gopherDS/HashMap"
)

func main() {
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
