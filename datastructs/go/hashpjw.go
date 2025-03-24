package datastructs

import "strconv"

// based on the C hashpjw, modified to work in Go
func HashString(str string, tblSize int) int {
	val := 0
	key, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}

	for i := 0; i < key; i++ {
		tmp := 0

		val = (val << 4) + key

		tmp = (val & 0xf000000)

		if tmp > 1 {
			val = val ^ (tmp >> 24)
			val = val ^ tmp
		}
	}

	return val % tblSize
}
