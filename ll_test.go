package go_linked_list

import "testing"

func makeLinkedList() *Node {
	node4 := &Node{4, nil}
	node3 := &Node{3, node4}
	node2 := &Node{2, node3}
	head := &Node{1, node2}
	return head
}

func areLinkedListsEqual(head1 *Node, head2 *Node) bool {
	curr1, curr2 := head1, head2

	for curr1 != nil && curr2 != nil {
		if curr1.Value != curr2.Value {
			return false
		}
		curr1 = curr1.Next
		curr2 = curr2.Next
	}

	if curr1 != nil || curr2 != nil {
		return false
	}
	return true
}

func TestContructLinkedList(t *testing.T) {
	cases := []struct{
		input []int
		expected *Node
	}{
		{[]int{1, 2, 3, 4}, makeLinkedList()},
	}
	for _, v := range cases {
		output := ConstructListFromSlice(v.input)
		if !areLinkedListsEqual(output, v.expected) {
			t.Error("Test failed")
		}
	}
}
