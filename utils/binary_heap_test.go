package utils

import (
	"fmt"
	"testing"
)

func TestMaxHeap(t *testing.T) {
	heap := BinaryHeapFrom(func(a, b int) int { return a - b }, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})

	if heap.Size() != 9 {
		t.Fatalf(Red(fmt.Sprintf("Expected size = 9, got = %d", heap.Size())))
	}

	max, err := heap.Peek()
	if err != nil {
		t.Fatalf(Red(fmt.Sprintf("Peek failed: %s", err.Error())))
	}
	if max != 9 {
		t.Fatalf(Red(fmt.Sprintf("Expected max = 9, got = %d", max)))
	}

	popped, err := heap.Pop()
	if err != nil {
		t.Fatalf(Red(fmt.Sprintf("Pop failed: %s", err.Error())))
	}
	if popped != 9 {
		t.Fatalf(Red(fmt.Sprintf("Expected max = 9, got = %d", max)))
	}

	max, err = heap.Peek()
	if err != nil {
		t.Fatalf(Red(fmt.Sprintf("Peek failed: %s", err.Error())))
	}
	if max != 8 {
		t.Fatalf(Red(fmt.Sprintf("Expected max = 8, got = %d", max)))
	}

	heap.Insert(100)
	max, err = heap.Peek()
	if err != nil {
		t.Fatalf(Red(fmt.Sprintf("Peek failed: %s", err.Error())))
	}
	if max != 100 {
		t.Fatalf(Red(fmt.Sprintf("Expected max = 100, got = %d", max)))
	}

	heap.Insert(20)
	max, err = heap.Peek()
	if err != nil {
		t.Fatalf(Red(fmt.Sprintf("Peek failed: %s", err.Error())))
	}
	if max != 100 {
		t.Fatalf(Red(fmt.Sprintf("Expected max = 100, got = %d", max)))
	}
}

func TestMinHeap(t *testing.T) {
	heap := BinaryHeapFrom(func(a, b int) int { return b - a }, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})

	if heap.Size() != 9 {
		t.Fatalf(Red(fmt.Sprintf("Expected size = 9, got = %d", heap.Size())))
	}

	min, err := heap.Peek()
	if err != nil {
		t.Fatalf(Red(fmt.Sprintf("Peek failed: %s", err.Error())))
	}
	if min != 1 {
		t.Fatalf(Red(fmt.Sprintf("Expected min = 1, got = %d", min)))
	}

	popped, err := heap.Pop()
	if err != nil {
		t.Fatalf(Red(fmt.Sprintf("Pop failed: %s", err.Error())))
	}
	if popped != 1 {
		t.Fatalf(Red(fmt.Sprintf("Expected min = 1, got = %d", min)))
	}

	min, err = heap.Peek()
	if err != nil {
		t.Fatalf(Red(fmt.Sprintf("Peek failed: %s", err.Error())))
	}
	if min != 2 {
		t.Fatalf(Red(fmt.Sprintf("Expected max = 2, got = %d", min)))
	}

	heap.Insert(-100)
	min, err = heap.Peek()
	if err != nil {
		t.Fatalf(Red(fmt.Sprintf("Peek failed: %s", err.Error())))
	}
	if min != -100 {
		t.Fatalf(Red(fmt.Sprintf("Expected min = -100, got = %d", min)))
	}

	heap.Insert(20)
	min, err = heap.Peek()
	if err != nil {
		t.Fatalf(Red(fmt.Sprintf("Peek failed: %s", err.Error())))
	}
	if min != -100 {
		t.Fatalf(Red(fmt.Sprintf("Expected min = -100, got = %d", min)))
	}
}

func TestPeekAndPopEmptyHeap(t *testing.T) {
	heap := BinaryHeap(func(a, b int) int { return a - b })
	_, err := heap.Peek()
	if err == nil {
		t.Fatalf(Red("Expected peek err != nil, got err = nil"))
	}

	_, err = heap.Pop()
	if err == nil {
		t.Fatalf(Red("Expected pop err != nil, got err = nil"))
	}
}
