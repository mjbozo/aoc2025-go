package utils

import (
	"fmt"
)

// Node data structure for binary heap
type heapNode[T comparable] struct {
	value  T
	left   *heapNode[T]
	right  *heapNode[T]
	parent *heapNode[T]
}

// Default formatting for binary heap node
func (n heapNode[T]) String() string {
	return fmt.Sprintf("HeapNode[Val=%v, Left=%v, Right=%v]", n.value, n.left, n.right)
}

// Helper method for swapping two node values, and returning the new swapped node
func (n *heapNode[T]) swap(other *heapNode[T]) *heapNode[T] {
	temp := n.value
	n.value = other.value
	other.value = temp
	n = other
	return n
}

// Error for Binary Heap operations
type HeapError struct {
	msg string
}

// Error implementation for HeapError
func (e *HeapError) Error() string {
	return e.msg
}

// Binary heap data structure
//
// Can be used as either min-heap or max-heap depending in the user-defined comparator function
type binaryHeap[T comparable] struct {
	Comparator func(a, b T) int
	root       *heapNode[T]
	size       int
}

func BinaryHeap[T comparable](comparator func(a, b T) int) *binaryHeap[T] {
	return &binaryHeap[T]{
		Comparator: comparator,
		root:       nil,
		size:       0,
	}
}

func BinaryHeapFrom[T comparable](comparator func(a, b T) int, vals []T) *binaryHeap[T] {
	heap := &binaryHeap[T]{
		Comparator: comparator,
		root:       nil,
		size:       0,
	}

	heap.InsertAll(vals)
	return heap
}

func (h binaryHeap[T]) Size() int {
	return h.size
}

// Default format for binary heap
func (h binaryHeap[T]) String() string {
	return fmt.Sprintf("BinaryHeap[Size=%d, Root=%v]", h.size, h.root)
}

// Array representation of binary heap
func (h binaryHeap[T]) Array() []T {
	arr := make([]T, 0)
	queue := make([]*heapNode[T], 0)
	queue = append(queue, h.root)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node != nil {
			arr = append(arr, node.value)
			queue = append(queue, node.left)
			queue = append(queue, node.right)
		}
	}

	return arr
}

// Helper method for maintaining valid heap structure
//
// Swaps node with parent as long as node value is greater than parent value, as determined by heap comparator
func (h *binaryHeap[T]) shiftUp(node *heapNode[T]) {
	for node.parent != nil && h.Comparator(node.value, node.parent.value) > 0 {
		node = node.swap(node.parent)
	}
}

// Helper method for maintaining valid heap structure
//
// Swaps node with largest child as long as node value is less than child value, as determined by heap comparator
func (h *binaryHeap[T]) shiftDown(node *heapNode[T]) {
	for (node.left != nil && h.Comparator(node.value, node.left.value) < 0) || (node.right != nil && h.Comparator(node.value, node.right.value) < 0) {
		if node.right != nil {
			// then i should have left node too
			if h.Comparator(node.left.value, node.right.value) >= 0 {
				node = node.swap(node.left)
			} else {
				node = node.swap(node.right)
			}
		} else {
			node = node.swap(node.left)
		}
	}
}

// Insert a value into binary heap
func (h *binaryHeap[T]) Insert(val T) {
	defer func() { h.size += 1 }()

	if h.root == nil {
		h.root = &heapNode[T]{value: val}
		return
	}

	var insertedNode *heapNode[T]
	queue := make([]*heapNode[T], 0)
	queue = append(queue, h.root)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.left == nil {
			node.left = &heapNode[T]{value: val, parent: node}
			insertedNode = node.left
			break
		}

		if node.right == nil {
			node.right = &heapNode[T]{value: val, parent: node}
			insertedNode = node.right
			break
		}

		queue = append(queue, node.left)
		queue = append(queue, node.right)
	}

	h.shiftUp(insertedNode)
}

// Inserts all elements of a slice into the heap
func (h *binaryHeap[T]) InsertAll(vals []T) {
	for _, val := range vals {
		h.Insert(val)
	}
}

// Pop the element with highest priority from the heap and return pointer to the value
func (h *binaryHeap[T]) Pop() (T, error) {
	if h == nil || h.root == nil {
		var zero T
		return zero, &HeapError{msg: "Cannot pop empty or nil heap"}
	}

	defer func() {
		if h != nil {
			h.size -= 1
		}
	}()

	maxValue := h.root.value

	// find right-most element in last layer
	last := h.root
	queue := make([]*heapNode[T], 0)
	queue = append(queue, h.root)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.left == nil {
			break
		}
		last = node.left

		if node.right == nil {
			break
		}
		last = node.right

		queue = append(queue, node.left)
		queue = append(queue, node.right)
	}

	// root is the only value in the heap
	if last == h.root {
		fmt.Println(h, maxValue)
		h.root = nil
		return maxValue, nil
	}

	h.root.value = last.value
	if last.parent.right != nil {
		last.parent.right = nil
	} else {
		last.parent.left = nil
	}

	h.shiftDown(h.root)

	return maxValue, nil
}

// Return the highest priority value in the heap, without removing it
func (h *binaryHeap[T]) Peek() (T, error) {
	if h == nil || h.root == nil {
		var zero T
		return zero, &HeapError{msg: "Cannot peek empty or nil heap"}
	}

	return h.root.value, nil
}
