# Rational

General rational for the algorithm plus comments.

## Nodes

Since I don't know what type of ID would go into the tree I thought that it would be easiest to use generics. Since the only requirement should be that it's comparable a fairly wide range of IDs could be used. Such as numbers, strings, or UUIDs.

Also since no requirements mentioned it being a binary tree I went with a tree with arbitrary number of children. It supports binary trees as input, and it should be easy to update the code if a hard requirement for binary trees were needed.

## Function

The function takes the root of a tree as it's only argument and returns the leftmost shallowest duplicate of a previous ID as a pointer and the depth (level) of the ID. That's unless it can't find a duplicate. The algorithm uses a set to remember which IDs that it has seen, searches one depth at a time, and the first duplicate found is returned. No need to keep looking.

I chose a generic ID (only requirement is that it's comparable) as I didn't know which type of ID might be used. Since it's an algorithm making it more generic isn't much of an issue and it's a fairly light use of generics.

The algorithm is a breadth first traversal of the tree. There's no indication that the tree has a structure that would make it quicker to search so it seems to fit well. A depth first traversal would in theory be the same, but as we specifically are looking for the depth in practice it can be quicker. When we have found a duplicate we immediately return it.

The time complexity would be O(n). Best case the duplicate would be the leftmost child on depth 1 and worst case we need to traverse the complete tree (n nodes). The work done for each node would be considered O(1) so it wouldn't affect the time complexity further.

The space complexity would be O(n). In the best case the set would only contain one member, so the space complexity would effectively by O(1), while if it's the last possible node visited that contains the duplicate it would need to save n-1 nodes in the set.

I wanted to avoid an O(n^2) brute force algorithm so I decided to go with a set to remember visited nodes. The brute force option would probably have used a fair bit of stack memory as well, so I don't think I would have saved on memory either as a trade off.

## Tests

Wrote tests for the most simple cases before implementing the algorithm. Empty tree, only a root node, and finding a single duplicate at the bottom and finding the shallowest duplicate out of two possible ones.

Created a simple tree generator, pretty quick and dirty, to be able to test more trees to increase the confidence that the algorithm is correct. For a real project it's very useful to have robust generators with a decent set of options for how to create test data. It's difficult to hand write and image larger test sets. Not to mention humans generally are poor random generators. Another approach can be to build a generator that writes to disk, and use those for future testing. Similar to how it's useful to have a sample size of real life data to test on.

Testing could also be expanded upon with benchmark and fuzz tests, since go supports both of those.