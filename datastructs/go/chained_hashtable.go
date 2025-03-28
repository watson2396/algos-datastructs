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
func (h *ChHtbl) hash(str string) int {
	return HashString(str, h.buckets)
}

// Chained Hashtable Match
func (h *ChHtbl) Match(str string, n Node) bool {
	if str == n.Name {
		return true
	} else {
		return false
	}
}

// Chained Hashtable Insert

// Chained Hashtable Remove

// Chained Hashtable Lookup
func (h *ChHtbl) Lookup(str string) (string, error) {

	bucket := h.hash(str)
	list := h.table[bucket]

	for element := list.ListHead(); element != nil; element = list.ListNext(*element) {
		if h.Match(str, *element) {
			return str, nil
		}
	}

	return "", nil
}
