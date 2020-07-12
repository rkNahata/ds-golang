package heaps

type Heap struct {
	Items []int
}

func (h *Heap) GetLeftChildIndex(parentIndex int) int {
	return 2*parentIndex + 1
}

func (h *Heap) GetRightChildIndex(parentIndex int) int {
	return 2*parentIndex + 2
}

func (h *Heap) GetParentIndex(childIndex int) int {
	return childIndex / 2
}

func (h *Heap) GetLeftChild(parentIndex int) int {
	return h.Items[h.GetLeftChildIndex(parentIndex)]
}

func (h *Heap) GetRightChild(parentIndex int) int {
	return h.Items[h.GetRightChildIndex(parentIndex)]
}

func (h *Heap) GetParent(child int) int {
	return h.Items[h.GetParentIndex(child)]
}

func (h *Heap) HasLeft(parentIndex int) bool {
	return h.GetLeftChildIndex(parentIndex) < len(h.Items)
}

func (h *Heap) HasRight(parentIndex int) bool {
	return h.GetRightChildIndex(parentIndex) < len(h.Items)
}

func (h *Heap) HasParent(childIndex int) bool {
	return h.GetParentIndex(childIndex) >= 0
}

func (h *Heap)Swap(idx1,idx2 int){
	h.Items[idx1],h.Items[idx2] = h.Items[idx2],h.Items[idx1]
}