//
// Created by HEADS on 2021/3/3.
//

#ifndef DATA_STRUCTURE_AVL_TREE_H
#define DATA_STRUCTURE_AVL_TREE_H

#include <iostream>
#include <vector>
using namespace std;

template<typename T>
typedef struct AVLNode {
    T val;
    AVLNode *left;
    AVLNode *right;
    int height;
    AVLNode(T val):val(val), left(NULL), right(NULL), height(1){};
    AVLNode(){};
};

template<typename T>
class AVLTree {
public:
    AVLTree() {
        root = NULL;
        size = 0;
    }

    ~AVLTree(){}

    // 向以node为根的树中添加元素
    AVLNode<T> add(AVLNode<T> *node, T val) {
        if (node == NULL) {
            size++;
            return new AVLNode<T>(val);
        }

        if (val < node->val) {
            node->left = add(node->left, val);
        } else if (val > node->val) {
            node->right = add(node->right, val);
        } else {
            node->val = val;
        }

        // 更新height
        node->height = 1 + max(getHeight(node->left), getHeight(node->right));

        // 计算平衡因子
        int balanceFactor = getBalanceFactor(node);

        // 如果不平衡发生在左子树，进行右旋转，然后返回，进入到父节点作同样操作
        if (balanceFactor > 1 && getBalanceFactor(node->left) > 0)
            return leftRotate(node);
        // 如果不平衡发生在右子树，进行左旋转
        else if (balanceFactor < -1 && getBalanceFactor(node->right) > 0)
            return rightRotate(node);
        // LR
        else if (balanceFactor > 1 && getBalanceFactor(node->left) < 0) {
            node->left = leftRotate(node->left);
            return rightRotate(node);
        }
        // RR
        else if (balanceFactor < -1 && getBalanceFactor(node->right) > 0) {
            node->right = rightRotate(node->right);
            return leftRotate(node);
        }

        return node;
    }

    void remove(T val) {
        AVLNode<T> node = AVLNode<T>(root, val);
        if (node == NULL) {
            root = remove(root, val);
        }
        return;
    }

    // 判断当前树是否是二分搜索树
    bool isBST() {
        vector<T> vals = vector<T>(size);

        // 利用中序遍历判断是否是有序数组
        inOrder(root, vals);

        for (int i = 1; i < vals.size(); i ++) {
            if (vals[i-1] > vals[i]) {
                return false;
            }
        }

        return true;
    }

    // 判断当前树是否是平衡二叉树
    bool isBalanced() {
        return isBalanced(root);
    }

private:
    AVLNode<T> *root;
    int size;

    int getHeight(AVLNode<T> *node) {
        if (node == NULL)
            return 0;
        return node->height;
    }

    int getBalanceFactor(AVLNode<T> *node) {
        if (node == NULL)
            return 0;
        return getHeight(node->left) - getHeight(node->right);
    }

    // 中序遍历
    void inOrder(AVLNode<T> *node, vector<T> &vals) {
        if (node == NULL)
            return;

        inOrder(node->left, vals);
        vals.push_back(node->val);
        inOrder(node->right, vals);
    }

    // 判断是否是平衡二叉树
    bool isBalanced(AVLNode<T> *node) {
        if (node == NULL)
            return true;

        int balanceFactor = getBalanceFactor(node);
        if (abs(balanceFactor) > 1)
            return false;

        return isBalanced(node->left) && isBalanced(node->right);
    }

    // 右旋转
    // 对节点y进行向右旋转操作，返回旋转后新的根节点x
    //        y                              x
    //       / \                           /   \
    //      x   T4     向右旋转 (y)        z     y
    //     / \       - - - - - - - ->    / \   / \
    //    z   T3                       T1  T2 T3 T4
    //   / \
    // T1   T2
    AVLNode<T> rightRotate(AVLNode<T> &y) {
        AVLNode<T> x = y.left;
        AVLNode<T> t3 = x.right;

        // 向右旋转
        x.right = y;
        y.left = t3;

        // 更新height
        y.height = max(getHeight(y.left), getHeight(y.right)) + 1;
        x.height = max(getHeight(x.left), getHeight(x.right)) + 1;

        return x;
    }


    // 左旋转
    // 对节点y进行向左旋转操作，返回旋转后新的根节点x
    //    y                             x
    //  /  \                          /   \
    // T1   x      向左旋转 (y)       y     z
    //     / \   - - - - - - - ->   / \   / \
    //   T2  z                     T1 T2 T3 T4
    //      / \
    //     T3 T4
    AVLNode<T> leftRotate(AVLNode<T> &y) {
        AVLNode<T> x = y.right;
        AVLNode<T> t3 = x.left;

        // 向左旋转
        x.left = y;
        y.right = t3;

        // 更新height
        y.height = max(getHeight(y.left), getHeight(y.right)) + 1;
        x.height = max(getHeight(x.left), getHeight(x.right)) + 1;

        return x;
    }

    AVLNode<T> remove(AVLNode<T> &node, T val) {
        if (node == NULL)
            return NULL;

        AVLNode<T> retNode;
        if (val < node.val) {
            node.left = remove(node.left, val);
            retNode = node;
        } else if (val > node.val) {
            node.right = remove(node.right, val);
            retNode = node;
        } else { // 如果当前结节是要删除的节点
            if (node.left == NULL) {
                AVLNode<T> rightNode = node.right;
                node.right = NULL; // 为了垃圾回收
                size --;
                retNode = rightNode;
            } else if (node.right == NULL) {
                AVLNode<T> leftNode = node.left;
                node.left = NULL;  // 为了垃圾回收
                size--;
                retNode = leftNode;
            } else { // 如果要删除的节点左、右子节点都不为NULL
                // 找到比待删除节点大的子树中最小节点，即待删除节点右子树的最小节点
                // 然后用这个节点替换被删除节点
                AVLNode<T> successor = minimum(node.right);
                // 因为最小节点已经替换成了被删除节点，所以要把最小节点也删除
                successor.right = remove(node.right, successor.val);
                successor.left = node.left;

                node.left = node.right = NULL;  // 为了垃圾回收

                retNode = successor;
            }
        }

        if (retNode == NULL)
            return NULL;

        retNode.height = 1 + max(getHeight(retNode.left), getHeight(retNode.right));

        // 计算平衡因子
        int balanceFactor = getBalanceFactor(retNode);

        // 平衡维护
        // LL
        if (balanceFactor > 1 && getBalanceFactor(retNode.left) >= 0) {
            return rightRotate(retNode);
        }
        // RR
        if (balanceFactor < -1 && getBalanceFactor(retNode.right) <= 0) {
            return leftRotate(retNode);
        }
        // LR
        if (balanceFactor > 1 && getBalanceFactor(retNode.left) < 0) {
            retNode.left = leftRotate(retNode.left);
            return rightRotate(retNode);
        }
        // RL
        if (balanceFactor < -1 && getBalanceFactor(retNode.right) > 0) {
            retNode.right = rightRotate(retNode.right);
            return leftRotate(retNode);
        }

        return retNode;
    }

    AVLNode<T> minimum(AVLNode<T> node) {
        if (node.left == NULL)
            return node;
        return minimum(node.left);
    }
};

#endif //DATA_STRUCTURE_AVL_TREE_H
