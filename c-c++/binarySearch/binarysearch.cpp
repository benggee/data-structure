//
// Created by HEADS on 2021/2/6.
//
// 使用二分查找，寻找一个半有序数组 [4, 5, 6, 7, 0, 1, 2] 中间无序的地方。同学们可以将自己的思路、代码写在学习总结中。
#include <iostream>
#include <vector>
using namespace std;

class BinarySearch {
public:
    int res;
    int search(vector<int> &nums, int start, int end) {
        if (end <= start) {
            res = 0;
            return 0;
        }
        if (end - start == 1 && nums[start] > nums[end]) {
            cout << "NUMS1:" << nums[start] << "-" << start << " NUMS2:" << nums[end] << "-" << end << endl;
            res = end;
            return end;
        }


        int mid = start + (end - start)/2;
        int l = search(nums, start, mid);
        int r = search(nums, mid + 1, end);

        if (r - l == 1 && nums[l] > nums[r]) {
            return r;
        }
        return 0;

        // cout << "l:" << l << "r:" << r << endl;




//        if (start == end)
//            return 0;

//        int l = start, r = end; // [0...i] 前闭后闭
//        while (l < r) {
//            int mid = l + (r - l)/2;
//            if (nums[l] > nums[r] && r - l == 1) {
//                return r;
//            }
//
//            if (nums[mid] > nums[r]) {
//                r = mid;
//            } else if (nums[mid] < nums[l]) {
//                l = mid;
//            } else {
//                l = search(nums, l, mid);
//                r = search(nums, mid+1, r);
//                break;
//            }
//        }
//
//        if (nums[l] < nums[r]) {
//            return r;
//        }
//
//        return 0;
    }
};

int main() {
    vector<vector<int>> nums = {
            {4, 5, 6, 7, 0, 1, 2},
            {5, 6, 7, 0, 1, 2, 4},
            {6, 7, 0, 1, 2, 4, 5},
            {0, 1, 2, 4, 6, 7, 5},
            {0, 1, 2, 4, 5, 6, 7},
            {0, 1},
            {1, 0},
            {0}
    };
    BinarySearch b = BinarySearch();

    for (int i = 0; i < nums.size(); i ++) {
        int res = b.search(nums[i], 0, nums[i].size()-1);
        cout << res << endl;
    }
}