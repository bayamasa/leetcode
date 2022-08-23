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
	is_deleted := false
	for v != nil && v.Next != nil {
		is_deleted = false
		if v.Val == v.Next.Val {
			(*v).Next = v.Next.Next
			is_deleted = true
		}
		if is_deleted {
			v = head
		} else {
			v = v.Next
		}
	}
	return head
	
}
// @lc code=end

