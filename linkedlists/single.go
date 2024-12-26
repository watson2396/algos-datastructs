// creating and managing a single linked listi

package linkedlists

type node struct {
	next *node
	data *data
}

type data struct {
	first_name string
	last_name  string
}

func init_list(nameList []data) []node {

	var llist []node

	var prevName node
	for i := 0; i < len(nameList); i++ {
		var name = nameList[i]

		n := node{}
		llist = append(llist)
	}
}
