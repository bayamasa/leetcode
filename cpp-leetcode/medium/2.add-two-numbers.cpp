/*
 * @lc app=leetcode id=2 lang=cpp
 *
 * [2] Add Two Numbers
 */

// @lc code=start
//  struct ListNode {
//       int val;
//       ListNode *next;
//       ListNode() : val(0), next(nullptr) {}
//       ListNode(int x) : val(x), next(nullptr) {}
//       ListNode(int x, ListNode *next) : val(x), next(next) {}
//   };

class Solution {
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        bool carry_forward = false;
        ListNode *tmp = new ListNode;
        ListNode *head = tmp;
        while (l1 != nullptr || l2 != nullptr) {
            // 計算パート
            int val;
                
            // どっちも値が存在
            if (l1 != nullptr && l2 != nullptr)
            {
                val = l1->val + l2->val;
            } else if (l1 == nullptr && l2 != nullptr) {
                val = l2->val;
            } else if (l1 != nullptr && l2 == nullptr) {
                val = l1->val;
            }
            if (carry_forward)
                val += 1;
            // 初期化
            carry_forward = false;
            //もし値が10以上だった場合
            if ((val / 10) != 0) {
                val %= 10;
                carry_forward = true;
            }
            tmp->val = val;
            if (l1 != nullptr)
                l1 = l1->next;
            if (l2 != nullptr)
                l2 = l2->next;
            if (l1 == nullptr && l2 == nullptr )
            {
                if (carry_forward)
                {
                    ListNode *next = new ListNode(1, nullptr);
                    tmp->next = next;
                }
                break;
            }
                
            // node作成パート
            ListNode *next = new ListNode;
            tmp->next = next;
            tmp = tmp->next;
        }
        return head;
    }
};
// @lc code=end

