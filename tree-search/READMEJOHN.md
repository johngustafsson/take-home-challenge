# Rational
## Nodes

Since I don't know what type of ID would go into the tree I thought that it would be easiest to use generics. Since the only requirement should be that it's comparable a fairly wide range of IDs could be used. Such as numbers, strings, or UUIDs.

Also since no requirements meantiond it being a binary tree I went with a tree with arbitrary number of children. It supports binary trees as input, and it should be easy to update the code if a hard requirement for binary trees were needed.

## Function




## Tests

Writing tests for the simplest cases before implementing the algorithm. Empty tree, only a root node, and finding a single duplicate at the bottom and finding the shallowest duplicate out of two possible ones.

More testing to harden the code will follow when the basic algorith passes.