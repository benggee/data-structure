//
// Created by HEADS on 2021/2/18.
// 双向链表
//

#include <stdlib.h>
#include <string.h>

#include "dlist.h"

void dlist_init(DList *list, void(*destroy)(void *data)) {
    list->size = 0;
    list->destroy = destroy;
    list->head = NULL;
    list->tail = NULL;

    return;
}

void dlist_destroy(DList *list) {
    void *data;

    while(dlist_size(list) > 0) {
        if (dlist_remove(list, dlist_tail(list), (void **)&data) == 0 && list->destroy != NULL) {
            list->destroy(data);
        }
    }

    memset(list, 0, sizeof(list));
    return;
}

int dlist_ins_next(DList *list, DListElmt *element, const void *data) {
    DListElmt *new_element;

    if (element == NULL && dlist_size(list) != 0) {
        return -1;
    }

    if ((new_element = (DListElmt *)malloc(sizeof(DListElmt))) == NULL)
        return -1;

    new_element->data = (void *)data;
    if (dlist_size(list) == 0) {
        list->head = new_element;
        list->head->prev = NULL;
        list->head->next = NULL;
        list->tail = new_element;
    } else {
        new_element->next = element->next;
        new_element->prev = element;

        if (element->next == NULL)
            list->tail = new_element;
        else
            element->next->prev = new_element;

        element->next = new_element;
    }

    list->size++;
    return 0;
}

int dlist_ins_prev(DList *list, DListElmt *element, const void *data) {
    DListElmt  *new_element;

    if (element == NULL && dlist_size(list) != 0)
        return -1;

    if ((new_element = (DListElmt *)malloc(sizeof(DListElmt))) == NULL)
        return -1;

    new_element->data = (void *)data;
    if (dlist_size(list) == 0) {
        list->head = new_element;
        list->head->prev = NULL;
        list->head->next = NULL;
        list->tail = new_element;
    } else {
        new_element->next = element;
        new_element->prev = element->prev;

        if (element->prev == NULL)
            list->head = new_element;
        else
            element->prev->next = new_element;

        element->prev = new_element;
    }

    list->size++;
    return 0;
}

int dlist_remove(DList *list, DListElmt *element, void **data) {
    if (element == NULL || dlist_size(list) == 0)
        return -1;

    *data = element->data;

    if (element == list->head) {
        list->head = element->next;

        if (list->head == NULL)
            list->tail = NULL;
        else
            element->next->prev = NULL;
    } else {
        element->prev->next = element->next;
        if (element->next == NULL)
            list->tail = element->prev;
        else
            element->next->prev = element->prev;
    }

    free(element);

    list->size--;
    return 0;
}