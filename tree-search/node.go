package treesearch

// Node for a tree.
//
// # Using a generic ID so that any type of ID can be used
//
// ID a generic id, needs to be comparable
// Children array of
type Node[T comparable] struct {
	ID       T
	Children []*Node[T]
}
