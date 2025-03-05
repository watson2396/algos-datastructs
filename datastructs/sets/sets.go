package sets

import (
	"github.com/watson2396/algos-datastructs/datastructs/linkedlists"
)

type Set struct{ linkedlists.List }

func SetInit() Set {
	return Set{
		linkedlists.List{
			Size: 0,
			Head: nil,
		},
	}
}

// Set insert

// Set remove

// Set union

// Set intersection

// Set difference

// Set is member

func (s *Set) SetIsMember() {

}

// Set is subset

// Set is equal
