# Trees
This package implements some common tree data structures and its extensions.

- Binary Search Tree
- Max(Min) Heap
- Indexed Priority Queue (using Heap)


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

- A heap is a `Complete Binary Tree`
- A `Complete Binary Tree` is a tree that:
  - Every level is completely filled, except that the last level can be partially filled.
  - If a node is partially filled, its child must be the left child.
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

## Segment Tree

Segment Tree basically is used to speed up the following operation:
- For a given array of size N, from a range i to j, where 0 <= i < N and i <= j < N, query the minimum, maximum or sum from the range.

A naive approach is to traverse from i to j, and find the minimum, maximum or calculate the sum, this will introduce a time complexity of O(N), assuming N is very large.

The segment tree however, can speed up this range query to a time complexity of O(logN), by reducing the searching range, it's actually a binary search tree.

The segment tree can be defined as follows:
- For an array A of a length of N, the segment tree has 2N - 1 nodes.
- The segment tree is a complete binary tree.
- Each node of a segment tree store a segment or interval of the array.
- The leaf nodes are elements of array from A[0] to A[N-1].
- The root of will represent the whole array A[0~N).
- The internal nodes represent the union of elementary intervals/segment for A[i~j].
- For a node that represent the range A[i~j), the range is broken into A[i~(i+j)/2), A[(i+j)/2, j], represented by the left and right child respectively.

For example, the root node represents the interval of A[0-N), its left child represents the interval of A[0-N/2) and the right child represents the interval of A[N/2, N), so on and so forth.

> Tips: since segment tree is a complete binary tree, the best way to store it is to use an array and start the node at index 1, then for each node i, the left node is at i*2, and right node is at i*2+1.

For a segment tree, there are basically two things we can do about it once we built it, update and query.

### Build the segment tree
The algorithm to build the segment tree can be done using recursive approach, from bottom to up. There are three types of range query basically, for each type, we just need to adjust a little bit. For A[i] and A[i+1], the parent is sum(A[i]+A[i+1]), min(A[i], A[i+1]) or max(A[i], A[i+1]).
```
build(node, start, end, array) {
    if start == end {
        tree[node] = array[start]
    } else {
        mid = (start + end) / 2
        // build the left and right sub-tree recursively.
        build(node*2, start, mid, array)
        build(node*2+1, mid+1, end, array)

        // operation is min, max or sum.
        tree[node] = operation(tree[node*2], tree[node*2+1])
    }
}
```

### Update the segment tree
Updating a segment tree means update a val at a given index, the algorithm works similar as build, we firstly find it and from bottom to up update the parents recursively.
```
update(node, start, end, i, val) {
    if start == end {
        tree[node] = val
    } else {
        mid = (start + end) / 2
        if i >= start && i <= mid {
            update(node*2, start, mid, i, val)
        } else {
            update(node*2+1, mid+1, end, i, val)
        }
        // after updating the subtree, we update the node.
        // operation is min, max or sum.
        tree[node] = operation(tree[node*2], tree[node*2+1])
    }
}
```

### Query the segment tree
The query operation search the min, max or get sum from i to j. From the segment tree 
