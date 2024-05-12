package treesearch

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// Basic test to make sure that the empty case is handled correctly.
func TestEmptyRoot(t *testing.T) {
	id, depth := CheckDuplicateIDs[int](nil)

	assert.Nil(t, id)
	assert.Zero(t, depth)
}

// Basic test to make sure that just a single root node is handled correct,
// should simply return nil, 0 as there can't be any duplicates.
// Also tests that a string is correctly handled.
func TestOnlyRoot(t *testing.T) {
	root := &Node[string]{
		ID: "This is a test",
	}

	id, depth := CheckDuplicateIDs(root)

	assert.Nil(t, id)
	assert.Zero(t, depth)
}

// Basic test to make sure that the most basic found case is covered,
// just two nodes with the same ID. Should return that ID plus a depth
// of 1 (i.e. the duplicate).
// Also tests that a uuid is correctly handled.

func TestMinimalTree(t *testing.T) {
	id := uuid.New()
	root := &Node[uuid.UUID]{
		ID: id,
		Children: []*Node[uuid.UUID]{
			{
				ID: id,
			},
		},
	}

	foundId, depth := CheckDuplicateIDs(root)

	assert.NotNil(t, foundId)
	assert.Equal(t, id, *foundId)
	assert.Equal(t, 1, depth)
}

// Find the 5 on depth 2 (0, 1, 2) as it's the only duplicate (and thus the shallowest)
//
//	     5
//	    / \
//	   /   \
//	  3     9
//	 / \   / \
//	1   2 5   7
func TestMediumTreeBottom(t *testing.T) {
	root := &Node[int]{
		ID: 5,
		Children: []*Node[int]{
			{
				ID: 3,
				Children: []*Node[int]{
					{ID: 1},
					{ID: 2},
				},
			},
			{
				ID: 9,
				Children: []*Node[int]{
					{ID: 5},
					{ID: 7},
				},
			},
		},
	}
	foundId, depth := CheckDuplicateIDs(root)

	assert.NotNil(t, foundId)
	assert.Equal(t, 5, *foundId)
	assert.Equal(t, 2, depth)
}

// Find the 5 on depth 1 (0, 1) as it's the shallowest duplicate.
// Should ignore the id on level 2 as it is not the shallowest
//
//	     5
//	    / \
//	   /   \
//	  5     9
//	 / \   / \
//	1   2 5   7
func TestMediumTree(t *testing.T) {
	root := &Node[int]{
		ID: 5,
		Children: []*Node[int]{
			{
				ID: 3,
				Children: []*Node[int]{
					{ID: 1},
					{ID: 2},
				},
			},
			{
				ID: 9,
				Children: []*Node[int]{
					{ID: 5},
					{ID: 7},
				},
			},
		},
	}
	foundId, depth := CheckDuplicateIDs(root)

	assert.NotNil(t, foundId)
	assert.Equal(t, 5, *foundId)
	assert.Equal(t, 2, depth)
}

// Testing pseudo random trees of sizes from 0 to 999
//
// Using a generator help generate trees, and provide a target and a target depth
func TestLargeTree(t *testing.T) {
	for length := 0; length < 1000; length++ {
		g := makeGenerator(length)
		root, target, targetDepth := g.makeTree()

		found, depth := CheckDuplicateIDs(root)

		if target == nil {
			assert.Nil(t, found)
			assert.Zero(t, depth)
		} else {
			assert.NotNil(t, found)
			assert.Equal(t, target, found)
			assert.Equal(t, targetDepth, depth)
		}
	}
}
