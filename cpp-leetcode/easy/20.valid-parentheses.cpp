/*
 * @lc app=leetcode id=20 lang=cpp
 *
 * [20] Valid Parentheses
 */

// @lc code=start
class Solution {
public:
    bool isValid(string s) {
        stack<char> v;
        auto it = s.begin();
        while (it != s.end())
        {
            if (*it == '(' || *it == '{' || *it == '[')
            {
                v.push(*it);
            }
            if (*it == ')' || *it == '}' || *it == ']')
            {
                if (v.empty())
                    return false;
                if (*it == ')' && v.top() != '(')
                    return false;
                if (*it == '}' && v.top() != '{')
                    return false;
                if (*it == ']' && v.top() != '[')
                    return false;
                v.pop();
            }
            it++;
        }
        if (!v.empty())
            return false;
        return true;
    }
};
// @lc code=end

