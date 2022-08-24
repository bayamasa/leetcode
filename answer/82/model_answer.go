/*
 * @lc app=leetcode id=82 lang=golang
 *
 * [82] Remove Duplicates from Sorted List II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{0, head}
	previous := dummy
	current := head
	for current != nil {
		nextNode := current.Next
		for nextNode != nil && current.Val == nextNode.Val {
			nextNode = nextNode.Next
		}
		if current.Next == nextNode {
			// 値が変わっていないということ
			previous = current
		} else {
			previous.Next = current
		}
		current = nextNode
	}
	return dummy.Next
}
// @lc code=end

