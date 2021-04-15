package datastructure

type queue struct {
	data *linkList
}

func Queue() *queue {
	return &queue{
		data: Linklist(),
	}
}

func (q *queue) Push(data interface{}) {
	q.data.AddHead(data)
}

func (q *queue) Pop() interface{} {
	if q.data.size == 0 {
		return nil
	}
	return q.data.RemoveLast().data
}

func (q *queue) Size() int {
	return q.data.size
}