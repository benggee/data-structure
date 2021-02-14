//
// Created by HEADS on 2021/2/5.
// No.209号问题
// 长度最小的子数组
// https://leetcode-cn.com/problems/minimum-size-subarray-sum/
//
#include<iostream>
#include<vector>
using namespace std;

// 滑动窗口
// 时间复杂度：O(n)
// 空间复杂度：O(1)
class Solution {
public:
    int minSubArrayLen(int target, vector<int>& nums) {
        int l = 0, r = -1;  // 初始窗口区间
        int sum = 0;  // 当前连续子集的值
        int res = nums.size() + 1; // 结果

        while (l < nums.size() ) {
            if (r + 1 < nums.size() && sum < target) {
                r ++;
                sum += nums[r];
            } else {
                sum -= nums[l];
                l++;
            }

            if (sum >= target) {
                res = min(res, r - l + 1);
            }
        }

        // 如果没有找到结果直接返回 0
        if (res == nums.size() + 1)
            return 0;

        return res;
    }
};

int main() {
    int target = 7;
    vector<int> nums = {2,3,1,2,4,3};

    Solution s = Solution();
    int res  = s.minSubArrayLen(target, nums);
    cout << res << endl;
}