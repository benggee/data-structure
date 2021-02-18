//
// Created by HEADS on 2021/2/16.
//
#include <stdlib.h>
#include <stdio.h>
#include <time.h>
#include "list.h"
#include "list.c"
#include "dlist.h"
#include "dlist.c"

void test_list(); // 单链表
void destroy1(void*);
void print_list(ListElmt*);
void test_dlist(); // 又向链表
void print_dlist(DListElmt*);

int main() {
    // test_list();
    test_dlist();
    return 0;
}

void test_dlist() {
    int n = 10;
    srand(time(0));

    DList l1;
    dlist_init(&l1, &destroy1);

    clock_t  start_time = clock();

    for (int i = 0; i < n; i ++) {
        int* tmp = malloc(sizeof(int));
        *tmp = rand() % (i + 1);
        dlist_ins_next(&l1, dlist_head(&l1), tmp);
    }

    clock_t end_time = clock();

    printf("Insert %d nums element time:%2lf\n", n, ((double)(end_time-start_time))/CLOCKS_PER_SEC);

    print_dlist((dlist_head(&l1)));

    printf("\n");

    int *data;
    dlist_remove(&l1, dlist_head(&l1), (void **)&data);

    print_dlist(dlist_head(&l1));
    printf("\n");

    int* tmp = malloc(sizeof(int));
    *tmp = 100;
    dlist_ins_prev(&l1, dlist_head(&l1), tmp);

    print_dlist(dlist_head(&l1));

    dlist_destroy(&l1);
}

void test_list() {
    int n = 10;
    srand(time(0));

    List l1;

    list_init(&l1, &destroy1);


    printf("test list start...\n");

    clock_t start = clock();

    for (int i = 0; i < n; i ++) {
        int* tmp = malloc(sizeof(int));
        *tmp = rand() % (i + 1);
        list_ins_next(&l1, NULL, tmp);
    }

    clock_t end = clock();

    printf("Insert %d nums element time: %2lf\n", n, ((double)(end-start))/CLOCKS_PER_SEC);

    // 遍历链表
    print_list(list_head(&l1));
    printf("\n");

    // 删除元素
    int rmElm;
    list_rem_next(&l1, NULL, (void **)&rmElm);

    // 遍历链表
    print_list(list_head(&l1));

    list_destroy(&l1);

}

void destroy1(void *data) {
    free(data);
}

void print_list(ListElmt *head) {
    ListElmt *cur = head;
    while(cur != NULL) {
        if (cur->next != NULL)
            printf("%d->", *(int*)cur->data);
        else
            printf(" %d", *(int*)cur->data);

        cur = cur->next;
    }
}


void print_dlist(DListElmt *head) {
    DListElmt *cur = head;
    while(cur != NULL) {
        if (cur->next != NULL)
            printf("%d->", *(int*)cur->data);
        else
            printf(" %d", *(int*)cur->data);

        cur = cur->next;
    }
}
