/*
 * @lc app=leetcode id=83 lang=golang
 *
 * [83] Remove Duplicates from Sorted List
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
    v := head
	for v != nil && v.Next != nil {
		if v.Val == v.Next.Val {
			v.Next = v.Next.Next
		} else {
			v = v.Next
		}
	}
	return head
	
}
// @lc code=end

