#define COLOR_RED  1
#define COLOR_BLACK  0

typedef int KEY_TYPE;

#define RBTREE_ENTRY(name, type) \
    struct name {                \
        struct type *right;      \
        struct type *left;       \
        struct type *parent;     \
        unsigned char color;     \
    }

typedef struct _rbtree_node {
    KEY_TYPE key;
    void *value;
#if 0
    struct _rbtree_node *right;
    struct _rbtree_node *left;

    struct _rbtree_node *parent;

    unsigned char color;
#else
    RBTREE_ENTRY(, _rbtree_node) _rbtree_node;
#endif
} rbtree_node;

typedef struct _rbtree {
    struct _rbtree_node *root;
    struct _rbtree_node *nil;
} rbtree;

// 这里的T是一个公共节点，其中root节点用来表示整颗红黑树的根节点
// nil表示所有叶子节点（也就是所有叶子节点都指向这个指针）
void rbtree_left_rotate(rbtree *T, rbtree_node *x) {
    // 先将要旋转的右节点取出来
    rbtree_node *y = x->right;

    // 改变x节点的右节点
    x->right = y->left;
    // 如果y的左节点不是一个叶子节点
    // 将parent指向x
    if (y->left != T->nil) {
        y->left->parent = x;
    }

    // 改变y的parent节点
    y->parent = x->parent;
    // 如果x的parent为空，说明x是根节点
    // 所以，y就变成了根节点
    if (x->parent == T->nil) {
        T->root = y;
    }
        // 到这里，说明x不是根节点，并且此时我们是不知道x是左子树还是右子树
        // 如果是左子树，就将parent的左子树指向y
    else if (x == x->parent->left) {
        x->parent->left = y;
    } else { // 不是左子树就一定是右子树
        x->parent->right = y;
    }

    // 改变y的左子树
    y->left = x;
    x->parent = y;
}

void rbtree_right_rotate(rbtree *T, rbtree_node *y) {
    rbtree_node *x = y->left;

    // 改变y节点的左子树
    y->left = x->right;
    // 如果x的右子树不是一个叶子节点
    if (x->right != T->nil) {
        x->right->parent = y;
    }

    // 改变x的parent节点
    x->parent = y->parent;
    // 如果y的parent为空，说明y是根节点
    // 所以，x此时就变成了根节点
    if (y->parent == T->nil) {
        T->root = x;
    }
        // 到这里，说明y不是根节点，此时y是左子树还是右子树我们不知道
    else if (y == y->parent->right) {
        y->parent->right = x;
    } else {
        y->parent->left = x;
    }

    // 改变x的左子树
    x->right = y;
    y->parent = x;
}

void rbtree_insert_fixup(rbtree *T, rbtree_node *z) {

    // 这里z本身是红色的，如果父节点是红色的，违背规则，继续调整
    while (z->parent->color == COLOR_RED) {
        // z的叔父节点
        rbtree_node *y = z->parent->parent->right;

        // 反转两层的颜色
        if (y->color == COLOR_RED) {
            z->parent->color = COLOR_BLACK;
            y->color = COLOR_BLACK;
            z->parent->parent->color = COLOR_RED;

            z = z->parent->parent; // z时刻是红色的
        } else { // y 是黑色
            if (z == z->parent->left) {
                z->parent->color = black;
                z->parent->parent->color = COLOR_RED;

                rbtree_right_rotate(T, z->parent->parent);
            } else {
                if (z == z->parent->right) {
                    z = z->parent;
                    rbtree_left_rotate(T, z);
                }

                z->parent->color = COLOR_BLACK;
                z->parent->parent->color = COLOR_RED;

                rbtree_right_rotate(T, z->parent->parent);
            }
        }
    }
}

void rbtree_insert(rbtree *T, rbtree_node *z) {
    rbtree_node *y = T->nil;
    rbtree_node *x = T->root;

    while(x != T->nil) {
        y = x;
        if (z->key < x->key) {
            x = x->left;
        } else if (z->key > x->key) {
            x = x->right;
        } else {
            // 已存在
            return;
        }
    }

    if (y == T->nil) {
        T->root = z;
    } else {

        if (z->key < y->key) {
            y->left = z;
        } else {
            y->rifht = z;
        }
    }

    z->parent = y;

    z->left = T->nil;
    z->right = T->nil;
    z->color = COLOR_RED;
}
