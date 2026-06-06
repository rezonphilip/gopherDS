/*
Package hashmap implements a simple hash table using separate chaining
with linked lists for collision resolution.

Hashing
--------

Keys are mapped to buckets using the FNV-1a (Fowler-Noll-Vo) hashing
algorithm. FNV-1a is a fast, non-cryptographic hash function commonly
used in hash table implementations due to its simplicity and good
distribution characteristics.

The algorithm works by:

  1. Starting with a fixed offset basis value.
  2. XORing each character of the key into the hash.
  3. Multiplying the hash by a carefully chosen prime number.
  4. Repeating for every character in the key.

This process produces a deterministic hash value, meaning the same key
will always generate the same hash.

Bucket Selection
----------------

The resulting hash value is converted into a valid bucket index using:

	hash % number_of_buckets

For example:

	hash("apple") -> 123456789
	123456789 % 10 -> bucket 9

Collision Handling
------------------

Different keys may occasionally produce the same bucket index. This is
known as a collision.

To handle collisions, each bucket stores a linked list of entries
(separate chaining). When multiple keys hash to the same bucket, new
entries are added to the bucket's chain and lookups traverse that chain
until the requested key is found.

Average-case performance:

	Insert: O(1)
	Lookup: O(1)
	Delete: O(1)

Worst-case performance (many collisions):

	Insert: O(n)
	Lookup: O(n)
	Delete: O(n)
*/

package hashmap

import "fmt"

// entry represents a single key-value pair.
// Collisions are handled using a linked list of entries.
type entry struct {
	key   string
	value any
	next  *entry
}

// BasicHashMap uses separate chaining to resolve collisions.
type BasicHashMap struct {
	buckets []*entry
	Size    int
}

// NewBasicHashMap creates a hash map with the given number of buckets.
func NewBasicHashMap(capacity int) *BasicHashMap {
	return &BasicHashMap{
		buckets: make([]*entry, capacity),
		Size:    0,
	}
}

// hash computes the bucket index for a given key using FNV-1a.
//
// FNV-1a works by:
//  1. Starting with a fixed offset value.
//  2. XORing each character into the hash.
//  3. Multiplying by a prime to spread bit patterns.
//  4. Mapping the resulting hash to a bucket using modulo.
//
// This provides fast lookups and helps distribute keys evenly
// across buckets, reducing collisions.
func (h *BasicHashMap) hash(key string) int {
	var hash uint32 = 2166136261 // FNV offset basis

	for _, letter := range key {
		hash ^= uint32(letter)
		hash *= 16777619 // FNV prime
	}

	return int(hash) % len(h.buckets)
}

// Put inserts a new key-value pair or updates an existing one.
func (h *BasicHashMap) Put(key string, val any) {
	index := h.hash(key)

	// Check if the key already exists in the bucket.
	for e := h.buckets[index]; e != nil; e = e.next {
		if e.key == key {
			e.value = val
			fmt.Printf("updated key %s in bucket %d\n", e.key, index)
			return
		}
	}

	// Insert new entry at the head of the bucket's linked list.
	h.buckets[index] = &entry{
		key:   key,
		value: val,
		next:  h.buckets[index],
	}

	fmt.Printf("inserted key into bucket %d\n", index)
	h.Size++
}

// Get looks up a key and returns its value if found.
func (h *BasicHashMap) Get(key string) (any, bool) {
	index := h.hash(key)

	for e := h.buckets[index]; e != nil; e = e.next {
		if e.key == key {
			fmt.Printf("key found in bucket %d\n", index)
			return e.value, true
		}
	}

	fmt.Println("key not found")
	return nil, false
}
