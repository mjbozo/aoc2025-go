package utils

import (
	"fmt"
)

// A Set, implemented as a wrapper around Go's map type
type HashSet[T comparable] map[T]bool

func (s HashSet[T]) String() string {
	var output string
	for k := range s {
		if len(output) > 0 {
			output += ", "
		}
		output += fmt.Sprintf("%v", k)
	}
	return fmt.Sprintf("[%s]", output)
}

// Checks if key exists in the set
func (s *HashSet[T]) Contains(key T) bool {
	_, exists := (*s)[key]
	return exists
}

// Inserts a key in the set
//
// Returns true if the key is newly inserted into the set
func (s *HashSet[T]) Insert(key T) bool {
	exists := s.Contains(key)
	(*s)[key] = true
	return !exists
}

// Removes a key from the set
//
// Returns true if the key was successfully removed
func (s *HashSet[T]) Remove(key T) bool {
	if s.Contains(key) {
		delete(*s, key)
		return true
	}

	return false
}
