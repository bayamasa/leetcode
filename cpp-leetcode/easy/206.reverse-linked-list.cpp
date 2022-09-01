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
        ListNode *tmp = head;
        stack<ListNode *> v;
        while (tmp != nullptr)
        {
            v.push(tmp);
            tmp = tmp->next;
        }
        ListNode *loop = new ListNode();
        ListNode *res = loop;
        while (!v.empty())
        {
            loop->next = new ListNode(v.top()->val);
            v.pop();
            loop = loop->next;
        }
        return res->next;
    }
};
// @lc code=end

