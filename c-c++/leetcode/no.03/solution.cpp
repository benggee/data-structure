//
// Created by HEADS on 2021/2/5.
// No.03
// 无重复字符的最长子串
// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
//
#include <iostream>
#include <vector>
#include <string>
using namespace std;

class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        int l = 0, r = -1;
        int res = 0;
        int freq[256] = {0};

        while(l < s.size()) {
            // 如果在滑动窗口没有和当前字符重复的字符，则滑动窗口继续向右滑动
            // 细节：freg[s[r+1]] 表示获取assic码为s[r+1]的位置的状态
            if (freq[s[r + 1]] == 0 && r + 1 < s.size()) {
                r ++;
                freq[s[r]] ++;
            } else {
                freq[s[l]] --;
                l ++;
            }

            res = max(res, r -l +1);
        }
        return res;
    }
};

int main() {
    string str = "aaabccdssl";

    Solution s = Solution();
    int res = s.lengthOfLongestSubstring(str);
    cout << res << endl;
}