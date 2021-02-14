//
// Created by HEADS on 2021/2/5.
// No.167 两数之和II
// https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted/
//
#include <iostream>
#include <vector>
using namespace std;

// 暴力法
// 两重循环
// 时间复杂度：O(n^2)
// 空间复杂度：O(n)
class Solution {
public:
    vector<int> twoSum(vector<int>& numbers, int target) {
        for (int i = 0; i < numbers.size(); i ++) {
            for (int j = 0; j < numbers.size(); j ++) {
                if (numbers[i] + numbers[j] == target && i != j) {
                    int tmp[2] = {i + 1, j + 1};
                    return vector<int>(tmp, tmp + 2);
                }
            }
        }

        throw invalid_argument("not found result.");
    }
};

// 二分法
// 记录下已经遍历过的元素
// 时间复杂度：O(nlogn)
// 空间复杂度：O(1)
class Solution1 {
public:
    vector<int> twoSum(vector<int>& numbers, int target) {
        for (int i = 0; i < numbers.size(); i ++) {
            for (int j = i + 1; i < numbers.size(); j ++) {
                if (target - numbers[i] == numbers[j]) {
                    int res[2] = {i + 1, j + 1};
                    return vector<int>(res, res + 2);
                }
            }
        }

        throw invalid_argument("not found result.");
    }
};

// 使用对撞指针
// 时间复杂度：O(logn)
// 空间复杂度：O(1)
class Solution2 {
public:
    vector<int> twoSum(vector<int>& numbers, int target) {
        int l = 0, r = numbers.size() - 1;
        while(l < r) {
            if (numbers[l] + numbers[r] == target) {
                int res[2] ={l + 1, r + 1};
                return vector<int>(res, res + 2);
            } else if (numbers[l] + numbers[r] > target) {
                r --;
            } else {
                l ++;
            }
        }

        throw invalid_argument("not found result.");
    }
};

int main() {
    vector<int> numbers = {2,7,11,15};
    int target = 9;

//    Solution s = Solution();
//    vector<int> res = s.twoSum(numbers, target);

//    Solution1 s1 = Solution1();
//    vector<int> res = s1.twoSum(numbers, target);
    Solution2 s2 = Solution2();
    vector<int> res = s2.twoSum(numbers, target);

    for (int i = 0; i < res.size(); i ++) {
        cout << res[i] << endl;
    }
}