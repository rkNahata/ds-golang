package maxheap

import (
	"ds-golang/heaps"
	"log"
)

type MaxHeap struct {
	*heaps.Heap
}

func New(input []int) *MaxHeap {
	h := &MaxHeap{&heaps.Heap{Items: input}}
	if len(h.Items) > 0 {
		h.buildHeap()

	}
	return h
}

func (h *MaxHeap) ExtractMax() int {
	if len(h.Items) == 0 {
		log.Fatal("heap is empty")
	}
	max := h.Items[0]
	lastIndex := len(h.Items) - 1
	h.Swap(0, lastIndex)
	h.Items = h.Items[:lastIndex]
	h.maxHeapifyDown(0)

	return max

}

func (h *MaxHeap) Insert(item int) *MaxHeap {
	h.Items = append(h.Items, item)
	h.maxHeapifyUp(len(h.Items) - 1)
	return h
}

func (h *MaxHeap) maxHeapifyDown(idx int) {
	for (h.HasLeft(idx) && h.GetLeftChild(idx) > h.Items[idx]) || (h.HasRight(idx) && h.GetRightChild(idx) > h.Items[idx]) {
		if (h.HasLeft(idx) && h.GetLeftChild(idx) > h.Items[idx]) && (h.HasRight(idx) && h.GetRightChild(idx) > h.Items[idx]) {
			if h.GetLeftChild(idx) > h.GetRightChild(idx) {
				h.Swap(idx, h.GetLeftChildIndex(idx))
				idx = h.GetLeftChildIndex(idx)
			} else {
				h.Swap(idx, h.GetRightChildIndex(idx))
				idx = h.GetRightChildIndex(idx)
			}
		} else if h.HasLeft(idx) && h.GetLeftChild(idx) > h.Items[idx] {
			h.Swap(idx, h.GetLeftChildIndex(idx))
			idx = h.GetLeftChildIndex(idx)
		} else {
			h.Swap(idx, h.GetRightChildIndex(idx))
			idx = h.GetRightChildIndex(idx)
		}
	}
}

func (h *MaxHeap) maxHeapifyUp(idx int) {
	for h.HasParent(idx) && h.GetParent(idx) < h.Items[idx] {
		h.Swap(idx, h.GetParentIndex(idx))
		idx = h.GetParentIndex(idx)
	}
}

func (h *MaxHeap) buildHeap() {
	for i := len(h.Items)/2 - 1; i >= 0; i-- {
		h.maxHeapifyDown(i)
	}
}

func (h *MaxHeap) Size() int {
	return len(h.Items)
}

func (h *MaxHeap) buildHead1(count map[int]int) {
	for i := len(h.Items)/2 - 1; i >= 0; i-- {
		h.maxHeapifyDown(i)
	}
}

func (h *MaxHeap) maxHeapifyDown1(idx int, count map[int]int) {
	for (h.HasLeft(idx) && count[h.GetLeftChild(idx)] > count[h.Items[idx]]) ||
		(h.HasRight(idx) && count[h.GetRightChild(idx)] > count[h.Items[idx]]) {
		if (h.HasLeft(idx) && count[h.GetLeftChild(idx)] > count[h.Items[idx]]) &&
			(h.HasRight(idx) && count[h.GetRightChild(idx)] > count[h.Items[idx]]) {
			if h.GetLeftChild(idx) > h.GetRightChild(idx) {
				h.Swap(idx, h.GetLeftChildIndex(idx))
				idx = h.GetLeftChildIndex(idx)
			} else {
				h.Swap(idx, h.GetRightChildIndex(idx))
				idx = h.GetRightChildIndex(idx)
			}
		} else if h.HasLeft(idx) && count[h.GetLeftChild(idx)] > count[h.Items[idx]] {
			h.Swap(idx, h.GetLeftChildIndex(idx))
			idx = h.GetLeftChildIndex(idx)
		} else {
			h.Swap(idx, h.GetRightChildIndex(idx))
			idx = h.GetRightChildIndex(idx)
		}
	}

}

func (h *MaxHeap) maxHeapifyUp1(idx int, count map[int]int) {
	for h.HasParent(idx) && count[h.GetParent(idx)] < count[h.Items[idx]] {
		h.Swap(idx, h.GetParentIndex(idx))
		idx = h.GetParentIndex(idx)
	}
}

func (h *MaxHeap) Insert1(item int,count map[int]int) *MaxHeap {
	h.Items = append(h.Items, item)
	h.maxHeapifyUp1(len(h.Items) - 1,count)
	return h
}