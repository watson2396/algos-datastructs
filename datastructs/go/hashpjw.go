package datastructs

func hashpjw_string(data string) int {

	var key string
	var val, PRIME_TBLSIZE int

	// TODO
	// 3 is certainly not right
	// find a better number
	PRIME_TBLSIZE = 3

	//  Hash the key by performing a number of bit operations on it
	val = 0
	key = data

	for pos, char := range key {
		var tmp int

		val = (val << 4) + (key)
	}

	return val % PRIME_TBLSIZE

}
