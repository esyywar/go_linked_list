package go_linked_list

import "fmt"

func ConstructListFromSlice(vals []int) *Node {
	if len(vals) == 0 {
		return nil
	}

	head := &Node{vals[0], nil}
	prev := head
	for _, v := range vals[1:] {
		new_node := &Node{v, nil}
		prev.Next = new_node
		prev = new_node
	}
	return head
}

func PrintList(head *Node) {
	curr := head
	node_count := 0
	for curr != nil {
		fmt.Printf("Node %d value: %d\n", node_count, curr.Value)
		node_count++
		curr = curr.Next
	}
}

func Reverse(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}

	prev := head
	curr := head.Next
	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}

	head.Next = nil
	return prev
}

func ReverseInPlace(head **Node) {
	if head == nil || *head == nil || (*head).Next == nil {
		return
	}

	prev := *head
	curr := (*head).Next
	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}

	(*head).Next = nil
	*head = prev
}