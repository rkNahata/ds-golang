package minheap

import (
	"ds-golang/heaps"
	"log"
)

type MinHeap struct {
	*heaps.Heap
}

func New(input []int) *MinHeap {
	h := &MinHeap{&heaps.Heap{Items: input}}
	if len(h.Items) > 0 {
		h.buildHeap()
	}
	return h
}

func (m *MinHeap) buildHeap() {
	//we ignore leaf nodes since they always follow heap property.
	for i := len(m.Items)/2 - 1; i >= 0; i-- {
		m.minHeapifyDown(i)
	}
}

func (m *MinHeap) ExtractMin() int {
	if len(m.Items) == 0 {
		log.Fatal("heap is empty")
	}
	min := m.Items[0]
	m.Swap(0, len(m.Items)-1)
	m.Items = m.Items[:len(m.Items)-1]
	m.minHeapifyDown(0)
	return min
}

func (m *MinHeap) Insert(item int) *MinHeap {
	m.Items = append(m.Items, item)
	m.minHeapifyUp(len(m.Items) - 1)
	return m
}

func (m *MinHeap) Size() int {
	return len(m.Items)
}

func (m *MinHeap) minHeapifyDown(index int) {
	for (m.HasLeft(index) && m.Items[index] > m.GetLeftChild(index)) ||
		(m.HasRight(index) && m.Items[index] > m.GetRightChild(index)) {
		if (m.HasLeft(index) && m.Items[index] > m.GetLeftChild(index)) &&
			(m.HasRight(index) && m.Items[index] > m.GetRightChild(index)) {
			if m.GetLeftChild(index) < m.GetRightChild(index) {
				m.Swap(index, m.GetLeftChildIndex(index))
				index = m.GetLeftChildIndex(index)
			} else {
				m.Swap(index, m.GetRightChildIndex(index))
				index = m.GetRightChildIndex(index)
			}
		} else if m.HasLeft(index) && m.Items[index] > m.GetLeftChild(index) {
			m.Swap(index, m.GetLeftChildIndex(index))
			index = m.GetLeftChildIndex(index)
		} else {
			m.Swap(index, m.GetRightChildIndex(index))
			index = m.GetRightChildIndex(index)
		}

	}
}

func (m *MinHeap) minHeapifyUp(index int) {
	for m.HasParent(index) && m.Items[index] < m.GetParent(index) {
		m.Swap(m.GetParentIndex(index), index)
		index = m.GetParentIndex(index)
	}
}

func (h *MinHeap) minHeapifyDown1(idx int, count map[int]int) {
	for (h.HasLeft(idx) && count[h.GetLeftChild(idx)] < count[h.Items[idx]]) ||
		(h.HasRight(idx) && count[h.GetRightChild(idx)] < count[h.Items[idx]]) {
		if (h.HasLeft(idx) && count[h.GetLeftChild(idx)] < count[h.Items[idx]]) &&
			(h.HasRight(idx) && count[h.GetRightChild(idx)] < count[h.Items[idx]]) {
			if h.GetLeftChild(idx) < h.GetRightChild(idx) {
				h.Swap(idx, h.GetLeftChildIndex(idx))
				idx = h.GetLeftChildIndex(idx)
			} else {
				h.Swap(idx, h.GetRightChildIndex(idx))
				idx = h.GetRightChildIndex(idx)
			}
		} else if h.HasLeft(idx) && count[h.GetLeftChild(idx)] < count[h.Items[idx]] {
			h.Swap(idx, h.GetLeftChildIndex(idx))
			idx = h.GetLeftChildIndex(idx)
		} else {
			h.Swap(idx, h.GetRightChildIndex(idx))
			idx = h.GetRightChildIndex(idx)
		}
	}

}

func (h *MinHeap) maxHeapifyUp1(idx int, count map[int]int) {
	for h.HasParent(idx) && count[h.GetParent(idx)] > count[h.Items[idx]] {
		h.Swap(idx, h.GetParentIndex(idx))
		idx = h.GetParentIndex(idx)
	}
}

func (h *MinHeap) Insert1(item int, count map[int]int) *MinHeap {
	h.Items = append(h.Items, item)
	h.maxHeapifyUp1(len(h.Items)-1, count)
	return h
}

func (m *MinHeap) ExtractMin1(count map[int]int) int {
	if len(m.Items) == 0 {
		log.Fatal("heap is empty")
	}
	min := m.Items[0]
	m.Swap(0, len(m.Items)-1)
	m.Items = m.Items[:len(m.Items)-1]
	m.minHeapifyDown1(0,count)
	return min
}