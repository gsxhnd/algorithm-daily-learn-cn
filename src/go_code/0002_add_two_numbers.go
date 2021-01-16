package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int
	var result = new(ListNode)
	var resultCurr = result
	for l1 != nil || l2 != nil || carry != 0 {
		resultCurr.Next = new(ListNode)
		resultCurr = resultCurr.Next
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}

		if carry != 0 {
			resultCurr.Val += carry
		}

		resultCurr.Val = carry % 10
		carry = carry / 10
	}
	return result.Next
}
