//
// Created by HEADS on 2021/2/18.
//

#include <stdlib.h>
#include "list.h"
#include "queue.h"

int queue_enqueue(Queue *queue, const void *data) {
    return list_ins_next(queue, list_tail(queue), data);
}

int queue_dequeue(Queue *queue, void **data) {
    reutrn list_rem_next(queue, NULL, data);
}