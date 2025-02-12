package main

import (
	"fmt"

	"github.com/watson2396/algos-datastructs/linkedlists"
)

func main() {
	fmt.Println("Hello")

	ll := linkedlists.Init()
	ll.SListAppend("Alice")
	ll.SListAppend("Bob")
	ll.SListPrepend("Charlie")
	ll.SListPrint()
}
