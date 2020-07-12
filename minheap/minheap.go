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
