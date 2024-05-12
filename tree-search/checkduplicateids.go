package treesearch

// Adding basic set struct to make the code a bit easier to read. It's trivial to
// write, but in a proper project it would be in a separate file(s) with tests etc.
//
// Using a struct{}{} as the value of the map to avoid wasting memory. Afaik it's
// the standard implementation for golang.
//
// Hopefully a set is added to the standard golang library in the future.
type set[T comparable] struct {
	members map[T]struct{}
}

// Creates an empty set
func makeSet[T comparable]() set[T] {
	return set[T]{
		members: map[T]struct{}{},
	}
}

// Adds a member to the set
func (s *set[T]) add(member T) {
	s.members[member] = struct{}{}
}

// Return true if the member is in the set, otherwise false
func (s *set[T]) has(member T) bool {
	_, has := s.members[member]
	return has
}

// Finds the shallowest duplicate node, starting from the root, and returns that node's
// ID and its depth/level. If no duplicate can be found nil and 0 will be returned.
//
// If multiple duplicates exists on the same depth, the left most will be returned.
func CheckDuplicateIDs[T comparable](root *Node[T]) (*T, int) {
	// Empty trees can't have duplicates
	if root == nil {
		return nil, 0
	}

	// See READMEJOHN.md for more information about the algorithm
	seen := makeSet[T]()
	children := []*Node[T]{root} // Preloading the first node, i.e. the root
	for depth := 0; len(children) > 0; depth++ {
		// New array for each level, easier than reusing the children array
		nextDepthChildren := []*Node[T]{}

		for _, node := range children {
			if !seen.has(node.ID) {
				seen.add(node.ID)
			} else {
				// Stop search when a duplicate is found
				return &node.ID, depth
			}

			nextDepthChildren = append(nextDepthChildren, node.Children...)
		}

		children = nextDepthChildren
	}

	return nil, 0
}
