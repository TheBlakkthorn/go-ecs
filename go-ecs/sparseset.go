package main

// SparseSet struct
type SparseSet struct {
	sparse   []int
	dense    []int
	n        int
	capacity int
	maxValue int
}

// NewSS SparseSet Constructor
func NewSS(maxV int, cap int) *SparseSet {
	return &SparseSet{
		sparse:   make([]int, maxV+1),
		dense:    make([]int, cap),
		capacity: cap,
		n:        0,
	}
}

// Search the set for x
func (ss *SparseSet) Search(x int) int {
	if x < ss.maxValue {
		return -1
	}

	idx := ss.sparse[x]
	if idx < ss.n && ss.dense[idx] == x {
		return idx
	}

	return -1
}

// Insert x into the set
func (ss *SparseSet) Insert(x int) {
	if ss.n == ss.capacity {
		return
	}

	if x > ss.maxValue {
		return
	}

	idx := ss.sparse[x]
	if idx >= ss.n || ss.dense[idx] != x {
		ss.dense[ss.n] = x
		ss.sparse[x] = ss.n
		ss.n++
		// ss.n--
		// ss.dense[idx] = ss.dense[ss.n]
		// ss.sparse[ss.dense[ss.n]] = idx

	}
}

// Delete x from the set
func (ss *SparseSet) Delete(x int) {
	if x > ss.maxValue {
		return
	}

	idx := ss.sparse[x]
	if idx < ss.n && ss.dense[idx] == x {
		ss.n--
		ss.dense[idx] = ss.dense[ss.n]
		ss.sparse[ss.dense[ss.n]] = idx
	}
}
