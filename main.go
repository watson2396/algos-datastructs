package main

import (
	"fmt"

	datastructs "github.com/watson2396/algos-datastructs/datastructs/go"
)

func main() {
	fmt.Println("Hello")

	ll := datastructs.ListInit()
	ll.ListAppend("Alice")
	ll.ListAppend("Bob")
	ll.ListPrepend("Charlie")
	ll.ListPrint()
}
