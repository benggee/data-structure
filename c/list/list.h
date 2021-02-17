//
// Created by HEADS on 2021/2/16.
//

#ifndef C___LINKLIST_H
#define C___LINKLIST_H

#include <stdlib.h>

typedef struct ListElmt_
{
    void *data;
    struct ListElmt_ *next;
} ListElmt;

typedef struct List_
{
    int size;
    int (*match)(const void *key1, const void *key2);
    void (*destroy)(void *data);
    ListElmt *head;
    ListElmt *tail;
}List;

void list_init(List *list, void (*destroy)(void *data));
void list_destroy(List *list);
int list_ins_next(List *list, ListElmt *element, const void *data);
int list_rem_next(List *list, ListElmt *element, void **data);

#define list_size(list) ((list)->head)
#define list_tail(list) ((list)->tail)
#define list_is_head(list, element) ((element) == (list)->head ? 1 : 0)
#define list_is_tail(element) ((element)->next == NULL ? 1 : 0)
#define list_data(element) ((element)->data)
#define list_next(element) ((element)->next)

#endif //C___LINKLIST_H
