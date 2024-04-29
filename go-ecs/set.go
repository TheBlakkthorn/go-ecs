package main

import "golang.org/x/exp/constraints"

// Set (unordered_set C++) for anything that can be treated as an int
type Set[T constraints.Integer] struct {
	Store map[T]bool
}

// Add T
func (s *Set[T]) Add(value T) {
	s.Store[value] = true
}

// Contains checks if the set contains T
func (s *Set[T]) Contains(value T) bool {
	return s.Store[value]
}

// Delete removes T
func (s *Set[T]) Delete(value T) {
	delete(s.Store, value)
}
