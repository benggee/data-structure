//
// Created by HEADS on 2021/2/6.
//
#include <iostream>
#include <vector>
#include <set>
using namespace std;

// 第一种写法
// 使用迭代器遍历set
class Solution {
public:
    vector<int> intersection(vector<int>& nums1, vector<int>& nums2) {
        set<int> record;
        for (int i = 0; i < nums1.size(); i ++) {
            record.insert(nums1[i]);
        }

        set<int> result;
        for (int i = 0; i < nums2.size(); i ++) {
            if (record.find(nums2[i]) != record.end()) {
                result.insert(nums2[i]);
            }
        }

        vector<int> resultVector;
        for (set<int>::iterator iter = result.begin(); iter != result.end(); iter ++) {
            resultVector.push_back( *iter );
        }
        return resultVector;
    }
};

// 第二种写法
// 直接相互转换
class Solution1 {
public:
    vector<int> intersection(vector<int> &nums1, vector<int> &nums2) {
        // 将vector 直接转成set
        set<int> record(nums1.begin(), nums1.end());

        set<int> resultSet;
        for (int i = 0; i < nums2.size(); i ++)
            if (record.find(nums2[i]) != record.end())
                resultSet.insert(nums2[i]);

        return vector<int>(resultSet.begin(), resultSet.end());
    }
};


int main() {
    vector<int> nums1 = {1,2,2,1};
    vector<int> nums2 = {2,2};
//
//    Solution s = Solution();
//    vector<int> res = s.intersection(nums1, nums2);

    Solution1 s1 = Solution1();
    vector<int> res = s1.intersection(nums1, nums2);

    for (int i = 0; i < res.size(); i ++) {
        cout << res[i] << endl;
    }
}