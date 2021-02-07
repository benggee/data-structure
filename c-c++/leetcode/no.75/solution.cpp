//
// Created by HEADS on 2021/2/5.
// 颜色分类问题
// https://leetcode-cn.com/problems/sort-colors/
//
#include <iostream>
#include <map>
#include <vector>
using namespace std;

// 记录0、1、2的数量，然后依次铺开
// 时间复杂度：O(n)
// 空间复杂度：O(1)
class Solution {
public:
    void sortColors(vector<int>& nums) {
        vector<int> tmp = vector<int>(3, 0);
        for (int i = 0; i < nums.size(); i ++) {
            assert(nums[i] <= 2);
            tmp[nums[i]]++;
        }

        int k = 0;
        for (int i = 0; i < tmp[0]; i ++)
            nums[k++] = 0;

        for (int i = 0; i < tmp[1]; i ++)
            nums[k++] = 1;

        for (int i = 0; i < tmp[2]; i ++)
            nums[k++] = 2;
    }
};

// 优化了三重循环
class Solution1 {
public:
    void sortColors(vector<int> &nums) {
        vector<int> tmp = vector<int>(3, 0);
        for (int i = 0; i < nums.size(); i ++) {
            assert(nums[i] <= 2);
            tmp[nums[i]] ++;
        }

        int k = 0;
        for (int i = 0; i < 3; i ++)
            for (int j = 0; j < tmp[i]; j ++)
                nums[k++] = i;
    }
};

// 使用三路快排的思想
// 时间复杂度：O(n)
// 空间复杂度：O(1)
class Solution2 {
public:
    void sortColors(vector<int> &nums) {
        int zero = -1; // 0 所在区域初始化下标[0...zero] == 0
        int two  = nums.size(); // 2所在获取下标 [two...n-1] == 2

        for (int i = 0; i < two; ) {
            if (nums[i] == 1) {
                i ++;
            } else if (nums[i] == 2) {
                two --;
                swap(nums[i], nums[two]);
            } else {
                assert(nums[i] == 0);
                zero ++;
                swap(nums[zero], nums[i]);
                i++;
            }
        }
    }
};



int main() {
    vector<int> test = {0,2,1,1,0,0,2,1};

//    Solution s = Solution();
//    s.sortColors(test);

//    Solution1 s1 = Solution1();
//    s1.sortColors(test);

    Solution2 s2 = Solution2();
    s2.sortColors(test);

    for (int i = 0; i < test.size(); i ++) {
        cout << test[i] << endl;
    }
}