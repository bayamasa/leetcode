/*
 * @lc app=leetcode id=142 lang=golang
 *
 * [142] Linked List Cycle II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func detectCycle(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
		return nil
	}
	if head == head.Next {
		return head
	}
	hare := head.Next
	tortoise := head
	isCycle := false
	cycleMaxSize := 0
	for hare != nil && hare.Next != nil {
		if hare == tortoise{
			isCycle = true
			break
		}
		hare = hare.Next.Next
		tortoise = tortoise.Next
		cycleMaxSize += 2
	}
	if !isCycle {
		return nil
	}
	tortoise = head
	for i := 0; i < cycleMaxSize; i++ {
		hare = tortoise.Next
		for j := i; j < cycleMaxSize; j++ {
			if (tortoise == hare) {
				return tortoise
			}
			hare = hare.Next
		}
		tortoise = tortoise.Next
	}
	return nil
}

// @lc code=end

