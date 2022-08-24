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
	// まずheadの判定から
	if head == nil || head.Next == nil {
		return head
	}
	for head.Val == head.Next.Val {
		for head.Val == head.Next.Val {
			head.Next = head.Next.Next
            if head == nil || head.Next == nil{
                return nil
            }
		}
		head = head.Next
        if head.Next == nil {
            break
        }
	}
	
	//ここでheadが重複しないことが確定。tmpを削除しなくてよいことが確定する
    tmp := head
	for tmp != nil && tmp.Next != nil && tmp.Next.Next != nil{
		for tmp.Next.Val == tmp.Next.Next.Val {
			for tmp.Next.Val == tmp.Next.Next.Val {
				tmp.Next.Next = tmp.Next.Next.Next
                if tmp.Next.Next == nil {
                    break
                }
			}
			tmp.Next = tmp.Next.Next
            if tmp.Next == nil || tmp.Next.Next == nil {
                break
            }
		}
		tmp = tmp.Next
	}
	return head
}
// @lc code=end

