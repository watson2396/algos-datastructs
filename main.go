package main

import (
	"fmt"
	"github.com/watson2396/algos-datastructs/linkedlists"
)

func main() {
	fmt.Println("Hello")

	ll := linkedlists.Init()
	ll.AppendSingle("Alice")
	ll.AppendSingle("Bob")
	ll.PrependSingle("Charlie")
	ll.PrintSingle()
}
