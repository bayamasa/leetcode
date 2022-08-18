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
	if (head == nil) {
		return false
	}
	ptr := head.Next
    var ptrSlice []*ListNode
	ptrSlice = append(ptrSlice, head)
    for ptr != nil {
        for i := 0; i < len(ptrSlice); i++ {
            if (ptr == ptrSlice[i]) {
                return true
            }
        }
		ptrSlice = append(ptrSlice, ptr)
        ptr = ptr.Next
    }
    return false
}
// @lc code=end

