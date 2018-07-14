# Trees
This readme introduces some common trees including

- Binary Search Tree
- Max(Min) Heap


## Binary Search Tree

Binary Search Tree has the following properties:

- The values in left sub-tree nodes are always smaller than the value in the root.
- The values in the right sub-tree nodes are always bigger than the value in the root.

### Insert

To insert a node into a binary search tree, just compare it with the current root, if it's smaller than the root, insert it into the left sub-tree, otherwise, insert into the right sub-tree. The time complexity is O(logN), the worse case scenario is O(N), in which case, the tree is extremely unbalanced and is essentially a linked list. The space complexity is O(N).


### Search

To search a node, it's the same of insertion, starting from the root, if the value matches, return; if the value is smaller than the root's value, search it recursively in the left sub-tree; if the value is bigger than the root's value, search it recursively in the right sub-tree. The time complexity is O(logN), the worse case scenario is O(N), in which case, the tree is extremely unbalanced and is essentially a linked list. The space complexity is O(N).

### Deletion

Deleting a node is tricky, first of all, find the node by the searching method. When the node is found, there are a few cases:

- The node is a leaf: we point the parent of this node to nil.
- The node has only left child: we point parent of this node to its left child.
- The node has only right child: we point parent of this node to its right child.
- The node has both children:
    - We find the maximum value in the node's left sub-tree.
    - We assign this value to the node.
    - We point the node's parent to the maximum node's left sub-tree.

> Tips: To make the deletion work, we make a fake parent who's right child is the root to start with.

The time complexity is O(logN), but the same as the above, worse case is O(N). The space complexity is O(N).
