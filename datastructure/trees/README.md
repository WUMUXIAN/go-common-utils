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


## Max(Min) Heap

Heap has the following properties:

- A heap is a Complete Binary Tree
- For a given node in a heap, the values of the nodes for its children are always smaller&bigger than this node's value, for Max and Min heap respectively.
- The sub-trees of a node are also heaps.

> Notes: The best way to store a heap is using an array list.
> for a given node with index i, the left child is i * 2, the right child is i * 2 + 1
> for a given node with index i, the parent node is i / 2.
> The root node is at index 1.
> We can use list[0] as a functional slot in heap operations.

For the following section, we assume it's a Max heap, for Min heap it's all the same, only the comparison is the other way round.

### Build Heap For a List

To build a heap from a given list, we do the following:
- Start from the parent of the last node, and loop back toward the root of the heap, for each of them:
    - If the node doesn't have left node, finish.
    - Find the node with larger value, let it be left or right.
    - If this value is not bigger than the value of the node, finish.
    - Otherwise switch the value, and let the current node to be this node.
    - Repeat.

    The time complexity of this operation is O(N).
    The space complexity of this operation is O(N).

### Insert

To insert a node into a heap, just attach it to the end of the list, which makes it the last leaf node. Then we do the following to make sure the heap is still valid:
- Start from the last leaf node you just added.
- Compare the node with its parent's value, if this value is bigger than the value of its parent, switch the value.
- Let the parent node be the current one, repeat the last operation unless we reach the root.

The time complexity of this operation is O(logN).
The space complexity of this operation is O(N).


### Peek/Pop

Peek is just to get the value of the top node.
Pop is to assign the value of the last node to the top node, and remove the last node.
Then do the following:
- Start from the root where index = 1
- If the node has no children, finish.
- Compare the value with its left and right child's value, get the large one, if it's the node itself, finish.
- If the largest value is the left or right child, switch this value with the node, and left node = node with largest value, repeat.

- The node is a leaf: we point the parent of this node to nil.
- The node has only left child: we point parent of this node to its left child.
- The node has only right child: we point parent of this node to its right child.
- The node has both children:
    - We find the maximum value in the node's left sub-tree.
    - We assign this value to the node.
    - We point the node's parent to the maximum node's left sub-tree.

The time complexity of this operation is O(logN).
The space complexity of this operation is O(N).
