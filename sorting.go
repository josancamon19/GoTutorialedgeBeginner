package main

import (
	"fmt"
	"sort"
)

type Kid struct {
	name string
	age  int
}

// Custom sort, override 3 functions
// Len() - returns the length of the array of items
// Swap() - a function which swaps the position of two elements in a sorted array
// Less() - a function which returns a bool value depending on whether the item at position i is less than the item at position j.

type kidsByAge []Kid

func (k kidsByAge) Len() int {
	return len(k)
}

func (k kidsByAge) Swap(i, j int) {
	k[i], k[j] = k[j], k[i]
}

func (k kidsByAge) Less(i, j int) bool {
	return k[i].age < k[j].age
}

func main7() {
	values := []int{4, 3, 2, 1}
	sort.Ints(values)
	fmt.Println(values)

	kids := []Kid{
		{
			name: "Joan",
			age:  0,
		},
		{
			name: "Santiago",
			age:  12,
		},
		{
			name: "Monroy",
			age:  9,
		},
		{
			name: "Cabezas",
			age:  -1,
		},
	}

	sort.Sort(kidsByAge(kids))
	fmt.Println(kids)
}
