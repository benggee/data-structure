//
// Created by HEADS on 2021/2/5.
//
#include <iostream>
#include <vector>
using namespace std;

// 1.使用额外的vector保存所有不为零的数
// 2.将不为零的数回填到原数组
// 3.将剩下的元素置为0
// 时间复杂度：O(n)
// 空间复杂度：O(n)
class Solution {
public:
    void moveZeroes(vector<int>& nums) {
        vector<int> tmp;

        for (int i = 0; i < nums.size(); i ++)
            if (nums[i] != 0)
                tmp.push_back(nums[i]);

        for (int i = 0; i < tmp.size(); i ++)
            nums[i] = tmp[i];

        for (int i = tmp.size(); i < nums.size(); i ++)
            nums[i] = 0;

    }
};

// 1. 交换0和非零
// 时间复杂度：O(n)
// 空间复杂度：O(1)
class Solution1 {
public:
    void moveZeroes(vector<int>& nums) {

        int k = 0;

        for (int i = 0; i < nums.size(); i ++) {
            if (nums[i] != 0)
                if (i != k) // 只有在i!=k的情况下才进行交换
                    swap(nums[k++], nums[i]);
                else
                    k++;
        }
    }
};

// 1. 将0值覆盖
// 2. 将最后值置为0
// 时间复杂度：O(n)
// 空间复杂度：O(1)
class Solution2 {
public:
    void moveZeroes(vector<int>& nums) {

        int k = 0;

        for (int i = 0; i < nums.size(); i ++) {
            if (nums[i] != 0)
                nums[k++] = nums[i];
        }

        for (int i = k; i < nums.size(); i ++) {
            nums[i] = 0;
        }
    }
};


int main() {
    vector<int> test = {1, 0, 3, 0, 5, 0, 12, 0};

//    Solution s = Solution();
//    s.moveZeroes(test);
//    for (int i = 0; i < test.size(); i ++) {
//        cout << test[i] << endl;
//    }

//    Solution1 s1 = Solution1();
//    s1.moveZeroes(test);
//
//    for (int i = 0; i < test.size(); i ++) {
//        cout << test[i] << endl;
//    }

    Solution2 s2 = Solution2();
    s2.moveZeroes(test);
    for (int i = 0; i < test.size(); i ++) {
        cout << test[i] << endl;
    }
}