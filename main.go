package main

import (
	"fmt"
	"go-bit-axon/bitset"
)

func main() {
	bs := bitset.New(100)
	err := bs.Set(50)
	if err != nil {
		panic(err)
	}

	fmt.Println(bs.Get(50)) // Output: true
	fmt.Println(bs.Get(51)) // Output: false

	err = bs.Set(150) // This will return an error
	if err != nil {
		fmt.Println(err) // Output: status : 400 , message : Index out of bounds
	}
}