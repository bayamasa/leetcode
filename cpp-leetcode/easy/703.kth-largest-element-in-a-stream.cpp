/*
 * @lc app=leetcode id=703 lang=cpp
 *
 * [703] Kth Largest Element in a Stream
 */

// @lc code=start
class KthLargest {
public:
    priority_queue<int, vector<int>, greater<int>> q;
    int size;
    KthLargest(int k, vector<int>& nums) {
        size = k;
        for (int n : nums) {
            q.push(n);
            if (q.size() > size) q.pop();
        }
    }
    
    int add(int val) {
        q.push(val);
        if (q.size() > size) q.pop();
        return q.top();
    }
};

/**
 * Your KthLargest object will be instantiated and called as such:
 * KthLargest* obj = new KthLargest(k, nums);
 * int param_1 = obj->add(val);
 */
// @lc code=end

