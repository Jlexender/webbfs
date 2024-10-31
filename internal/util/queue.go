package util

type Comparable[T any] interface {
	Compare(other T) bool
}

type PriorityQueue[T Comparable[T]] struct {
	items []T
	size  int
}

func (pq *PriorityQueue[T]) Len() int {
	return pq.size
}

func (pq *PriorityQueue[T]) siftUp(i int) {
	for i > 0 {
		parent := (i - 1) / 2
		if pq.items[parent].Compare(pq.items[i]) {
			break
		}
		pq.Swap(i, parent)
		i = parent
	}
}

func (pq *PriorityQueue[T]) siftDown(i int) {
	for left := 2*i + 1; left < pq.size; left = 2*i + 1 {
		right := left + 1
		j := left
		if right < pq.size && pq.items[right].Compare(pq.items[left]) {
			j = right
		}
		if pq.items[i].Compare(pq.items[j]) {
			break
		}
		pq.Swap(i, j)
		i = j
	}
}

func (pq *PriorityQueue[T]) Push(item T) {
	pq.items = append(pq.items, item)
	pq.size++
	pq.siftUp(pq.size - 1)
}

func (pq *PriorityQueue[T]) Pop() *T {
	if pq.size == 0 {
		return nil
	}
	item := pq.items[0]
	pq.items[0] = pq.items[pq.size-1]
	pq.size--
	pq.siftDown(0)
	return &item
}

func (pq *PriorityQueue[T]) Peek() *T {
	if pq.size == 0 {
		return nil
	}
	return &pq.items[0]
}

func (pq *PriorityQueue[T]) Less(i, j int) bool {
	return pq.items[i].Compare(pq.items[j])
}

func (pq *PriorityQueue[T]) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
}

func NewPriorityQueue[T Comparable[T]]() *PriorityQueue[T] {
	return &PriorityQueue[T]{}
}
