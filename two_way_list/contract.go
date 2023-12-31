package two_way_list

import "ooap/linked_list"

type Contract interface {
	linked_list.Contract

	// Left
	// Precondition: list is not empty & previous node exists
	// Postcondition: cursor shifted 1 node left
	Left()
}
