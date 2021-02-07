//
// Created by HEADS on 2021/2/7.
//
#include <iostream>
#include <vector>
#include <queue>
using namespace std;


struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

class Solution {
public:
    vector<vector<int>> levelOrder(TreeNode* root) {
        vector<vector<int>> res;
        if (root == NULL)
            return res;

        queue<pair<TreeNode*, int>> q;
        q.push(make_pair(root, 0));

        while(!q.empty()) {
            TreeNode* cur = q.front().first;
            int level = q.front().second;
            q.pop();

            if (level == res.size()) {
                res.push_back(vector<int>());
            }
            res[level].push_back(cur->val);

            if (cur->left)
                q.push(make_pair(cur->left, level + 1));
            if (cur->right)
                q.push(make_pair(cur->right, level + 1));
        }

        return res;
    }
};

