//
// Created by HEADS on 2021/2/16.
//

#ifndef C___LINKLIST_H
#define C___LINKLIST_H

#include <stdio.h>
#include <stdlib.h>

type struct Node {
    int data;
    struct Node *next;
}NODE,*PNODE;

PNODE create_list();
bool is_empty(PNODE);
int length_list(PNODE);
bool insert_list(PNODE, int, int);
bool delete_list(PNODE, int, int);
void destroy(PNODE);

#endif //C___LINKLIST_H
