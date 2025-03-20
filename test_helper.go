package main

import "math/rand"

func randomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)

	for i := 0; i < n; i++ {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
