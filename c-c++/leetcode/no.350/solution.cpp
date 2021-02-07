//
// Created by HEADS on 2021/2/6.
//
#include <iostream>
#include <vector>
#include <map>
using namespace std;

// 使用map实现
class Solution {
public:
    vector<int> intersect(vector<int>& nums1, vector<int>& nums2) {
        map<int,int> record;
        for (int i = 0; i < nums1.size(); i ++)
            if (record.find(nums1[i]) == record.end()) { // 不存在
                record.insert(make_pair(nums1[i], 1));
            } else {
                record[nums1[i]] ++;
            }

        vector<int> result;
        for (int i = 0; i < nums2.size(); i ++) {
            if (record.find(nums2[i]) != record.end() && record[nums2[i]] > 0) {
                result.push_back(nums2[i]);
                record[nums2[i]] --;
                if (record[nums2[i]] == 0)
                    record.erase(nums2[i]); // 删除
            }
        }
        return result;
    }
};

int main() {
    vector<int> nums1 = {1,2,2,1};
    vector<int> nums2 = {2,2};

    Solution s = Solution();
    vector<int> res = s.intersect(nums1, nums2);

    for (int i = 0; i < res.size(); i ++) {
        cout << res[i] << endl;
    }
}