package treesearch

import (
	"math/rand"

	"github.com/google/uuid"
)

// Simple and quick and dirty generator of trees for testing
type generator struct {
	length       int
	target       uuid.UUID
	firstTarget  int
	secondTarget int
	targetDepth  int
}

// Creates a generator
//
// Uses length as the seed to make it easier to debug, and generates two indexes
// where we put the target ID. Very simple generator.
func makeGenerator(length int) generator {
	// Need a length of at least two to get duplicates, still generates the 0 and 1 lengths.
	if length >= 2 {
		r := rand.New(rand.NewSource(int64(length)))
		first := r.Intn(length)
		var second int
		// loop until to make sure 1st and 2nd aren't identical
		for second = r.Intn(length); first == second; second = r.Intn(length) {
		}

		return generator{
			length:       length,
			target:       uuid.New(),
			firstTarget:  first,
			secondTarget: second,
		}
	} else {
		return generator{
			length:       length,
			target:       uuid.New(),
			firstTarget:  0,
			secondTarget: 0,
		}
	}
}

// Generates the next node
//
// Returns a node with the target id if the current length left equals the first or second target,
// otherwise returns a newly created random id.
func (g *generator) next() *Node[uuid.UUID] {
	if g.length <= 0 {
		return nil
	}
	g.length--

	// Generating the target if
	// We need two targets and there's only two spots left
	// We need one target, and there's only one spot left
	// We randomize a chance
	if g.length == g.firstTarget || g.length == g.secondTarget {
		return &Node[uuid.UUID]{ID: g.target}
	}

	// Generate new uuid that isn't the target. The chances are very small, but if it did ever happen it would be a pain to debug
	var new uuid.UUID
	for new = uuid.New(); new == g.target; new = uuid.New() {
	}
	return &Node[uuid.UUID]{ID: new}
}

// Creates a tree with a root, nil if length zero, a target id, nil
// if the length is less than two, and the depth where we expect to find
// the duplicate target.
func (g *generator) makeTree() (*Node[uuid.UUID], *uuid.UUID, int) {
	root := g.next()
	if root == nil {
		return nil, nil, g.targetDepth
	}

	depth := []*Node[uuid.UUID]{root}
	targetDepth := 1
	for len(depth) > 0 {
		newDepth := []*Node[uuid.UUID]{}
		for _, node := range depth {
			breath := 1 + rand.Intn(5)

			for i := 0; i < breath; i++ {
				child := g.next()
				if child == nil {
					break // length nodes have been created, no new ones will be added
				}
				if child.ID == g.target {
					g.targetDepth = targetDepth
				}
				node.Children = append(node.Children, child)
			}
			newDepth = append(newDepth, node.Children...)
		}

		depth = newDepth
		targetDepth++
	}

	if len(root.Children) > 0 {
		return root, &g.target, g.targetDepth
	} else { // tree too small to have a target
		return root, nil, g.targetDepth
	}
}
