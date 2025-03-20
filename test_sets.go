package main

import (
	"errors"
	"fmt"

	datastructs "github.com/watson2396/algos-datastructs/datastructs/go"
)

func generateSet(size int, stringSize int) (datastructs.Set, error) {

	if !(size > 0) || !(stringSize > 0) {
		return datastructs.Set{}, errors.New("need a non-zero set size and string size provided")
	}

	s := datastructs.SetInit()

	for i := 0; i <= size; i++ {
		name := randomString(stringSize)
		s.ListAppend(name)
	}

	return s, nil
}

func TestSetIsEqual() {
	s1, err := generateSet(5, 5)
	if err != nil {
		fmt.Printf("Cooked: %s", err)
	}

	s2, err := generateSet(5, 5)
	if err != nil {
		fmt.Printf("Cooked: %s", err)
	}

	if s1.SetIsEqual(&s2) {
		fmt.Println("Sets equal")
	} else {
		fmt.Println("Sets are not equal")
	}
}
