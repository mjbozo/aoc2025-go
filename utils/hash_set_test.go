package utils

import (
	"fmt"
	"testing"
)

func TestInsertElement(t *testing.T) {
	set := make(HashSet[int], 0)

	inserted := set.Insert(1)
	if !inserted {
		t.Fatalf(Red("Expected 1 to be inserted"))
	}

	inserted = set.Insert(2)
	if !inserted {
		t.Fatalf(Red("Expected 2 to be inserted"))
	}

	inserted = set.Insert(1)
	if inserted {
		t.Fatalf(Red("Expected 1 to not be inserted again"))
	}

	if len(set) != 2 {
		t.Fatalf(Red(fmt.Sprintf("Expected HashSet len = 2, got = %d", len(set))))
	}

	if !set.Contains(1) {
		t.Fatalf(Red("Expected HashSet contains 1"))
	}

	if !set.Contains(2) {
		t.Fatalf(Red("Expected HashSet contains 2"))
	}
}

func TestContainsElement(t *testing.T) {
	set := make(HashSet[int], 0)
	set.Insert(1)

	if !set.Contains(1) {
		t.Fatalf(Red("Expected HashSet contains 1"))
	}

	if set.Contains(2) {
		t.Fatalf(Red("Expected HashSet doesn't contain 2"))
	}

	if len(set) != 1 {
		t.Fatalf(Red(fmt.Sprintf("Expected HashSet len = 1, got = %d", len(set))))
	}
}

func TestRemoveElement(t *testing.T) {
	set := make(HashSet[int], 0)
	set.Insert(1)

	if !set.Contains(1) {
		t.Fatalf(Red("Expected HashSet contains 1"))
	}

	removed := set.Remove(1)
	if !removed {
		t.Fatalf(Red("Expected 1 to be removed"))
	}

	if set.Contains(1) {
		t.Fatalf(Red("Expected HashSet doesn't contain 1"))
	}

	if len(set) != 0 {
		t.Fatalf(Red(fmt.Sprintf("Expected HashSet len = 0, got = %d", len(set))))
	}

	removed = set.Remove(1)
	if removed {
		t.Fatalf(Red("Expected 1 to not be removed again"))
	}
}
