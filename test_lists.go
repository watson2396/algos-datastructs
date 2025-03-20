package main

import (
	"errors"
	"fmt"

	datastructs "github.com/watson2396/algos-datastructs/datastructs/go"
)

func TestLists() {
	_, err := generateList(5, 5)
	if err != nil {
		fmt.Printf("Cooked: %s", err)
	}

	// fmt.Println("List:")
	// l.ListPrint()
}

func generateList(listn int, stringSize int) (datastructs.List, error) {
	if !(listn > 0) || !(stringSize > 0) {
		return datastructs.List{}, errors.New("need a non-zero set size and string size provided")
	}

	l := datastructs.ListInit()

	for i := 0; i <= listn; i++ {
		str := randomString(stringSize)
		l.ListAppend(str)
	}

	return l, nil
}
