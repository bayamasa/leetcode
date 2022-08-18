/*
 * @lc app=leetcode id=141 lang=golang
 *
 * [141] Linked List Cycle
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func hasCycle(head *ListNode) bool {
	if (head == nil || head.Next == nil) {
		return false
	}
	hare := head
	tortoise := head
	for hare != nil && hare.Next != nil {
		hare = hare.Next.Next
		tortoise = tortoise.Next
		if (hare == tortoise) {
			return true
		}
	}
	return false
}
// @lc code=end

