//
// Created by HEADS on 2021/2/18.
//

#ifndef SET_H
#define SET_H

#include <stdlib.h>
#include "list.h"

type List Set;

void set_init(Set *set, int (*match)(const void *key1, const void key2), void (*destroy)(void *data));
int set_insert(Set *set, const void *data);
int set_remove(Set *set, void **data);
int set_union(Set *setu, const Set *set1, const Set *set2);
int set_intersection(Set *seti, const Set *set1, const Set *set2);
int set_difference(Set *setd, const Set *set1, const Set *set2);
int set_is_member(const Set *set, const void *data);
int set_is_subset(const Set *set1, const set *set2);
int set_is_equal(const Set *set1, const set *set2);

#define set_destroy list_destroy
#define set_size(set) ((set)->size)

#endif //SET_H
