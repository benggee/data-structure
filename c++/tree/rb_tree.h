//
// Created by HEADS on 2021/3/4.
//

#ifndef RB_TREE_H
#define RB_TREE_H

const bool RED = true;
const bool BLACK = false;

template<typename T>
typedef struct RBNode {
    T val;
    RBNode *left;
    RBNode *right;
    bool color;
    RBNode(T val):val(val),left(NULL),right(NULL),color(RED){}
    RBNode():left(NULL),right(NULL),color(RED){}
};


template<typename T>
class RBTree {
public:
    RBTree() {
        root = new RBNode<T>();
        size = 0;
    }

    ~RBTree() {}

    void add(T val) {
        root = add(root, val);
        root->color = BLACK;
    }

private:
    RBNode<T> *root;
    int size;

    RBNode<T> add(RBNode<T> &node, T val) {
        if (node == NULL) {
            size++;
            return new RBNode<T>(val);
        }

        if (val < node.val) {
            node->left = add(node->left, val);
        } else if (val > node.val) {
            node->right = add(node->right, val);
        } else {
            node->val = val;
        }

        if (isRed(node->right) && !isRed(node->left)) {
            node = leftRotate(node);
        }

        if (isRed(node->left) && isRed(node->left->left)) {
            node = rightRotate(node);
        }

        if (isRed(node.left) && isRed(node.right)) {
            flipColors(node);
        }

        return node;
    }

    // 颜色翻转
    void flipColors(RBNode<T> &node) {
        node->color = RED;
        node->left->color = BLACK;
        node->right->color = BLACK;
    }

    //   node                     x
    //  /   \     左旋转         /  \
    // T1   x   --------->   node   T3
    //     / \              /   \
    //    T2 T3            T1   T2
    void leftRotate(RBNode<T> &node) {
        RBNode<T> x = node.right;

        // 左旋转
        node.right = x.left;
        x.left = node;

        // 染色
        x.color = node.color;
        node.color = RED;

        return x;
    }

    //     node                   x
    //    /   \     右旋转       /  \
    //   x    T2   ------->   y   node
    //  / \                       /  \
    // y  T1                     T1  T2
    void rightRotate(RBNode<T> &node) {
        RBNode<T> x = node.left;

        // 右旋转
        node->left = x->right;
        x->right = node;

        // 染色
        x->color = node->color;
        node->color = RED;

        return x;
    }

    bool isRed(const RBNode<T> &node) {
        if (node == NULL)
            return false;
        return node.color;
    }
};

#endif //RB_TREE_H
