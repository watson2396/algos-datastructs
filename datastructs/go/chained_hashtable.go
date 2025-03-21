package datastructs

import "errors"

// Chained Hashtable struct
type ChHtbl struct {
	buckets int
	size    int
	table   []List
}

// Chained Hashtable Init
func ChHtblInit(buckets int) (ChHtbl, error) {
	if !(buckets > 0) {
		return ChHtbl{}, errors.New("Chained Hashtable error: need positive int for buckets input")
	}

	h := ChHtbl{
		buckets: buckets,
		size:    0,
		table:   nil,
	}

	for i := 0; i < h.buckets; i++ {
		h.table[i] = ListInit()
	}

	return h, nil
}

// Chained Hashtable Hash

// Chained Hashtable Match

// Chained Hashtable Insert

// Chained Hashtable Remove

// Chained Hashtable Lookup
