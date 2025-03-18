package main

import datastructs "github.com/watson2396/algos-datastructs/datastructs/go"

func TestLists() {
	ll := datastructs.ListInit()
	ll.ListAppend("Alice")
	ll.ListAppend("Bob")
	ll.ListPrepend("Charlie")
	ll.ListPrint()
}
