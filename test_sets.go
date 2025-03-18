package main

import (
	"errors"
	"fmt"
	"math/rand"

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

func randomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)

	for i := 0; i <= n; i++ {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func TestSetIsEqual() {

	s1, err := generateSet(5, 5)
	if err != nil {
		fmt.Println("Cooked")
	}

	s2, err := generateSet(5, 5)
	if err != nil {
		fmt.Println("Cooked")
	}

	s1.SetIsEqual(&s2)
}
