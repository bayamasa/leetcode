/*
 * @lc app=leetcode id=206 lang=cpp
 *
 * [206] Reverse Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    ListNode* reverseList(ListNode* head) {
        ListNode *prev = nullptr, *next = nullptr;
        while (head) {
            next = head->next; //次のptrの値を保持
            head->next = prev; //headのnextが指すptrを逆にする
            prev = head; // prevを先にすすめる
            head = next; // 保持していた値をreturn
        }
        return prev;
    }
};
// @lc code=end

