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

> Tips: since segment tree is a complete binary tree, the best way to store it is to use an array and start the node at index 1, then for each node i, the left node is at i*2, and right node is at i*2+1.

For a segment tree, there are basically two things we can do about it once we built it, update and query.

### Build the segment tree
The best way to build the segment tree is to build it from bottom up because we know from tree[n] to tree[2n-1], these leaf nodes represents A[i] to A[i-1], which are the values of the array. And also for a complete binary tree, a node's left child is node*2, right child is node * 2 + 1. The following code assumes we are calculating sum:
```
build(array) {
  for i = 0; i < len(array); i++ {
    tree[i+n] = array[i]
  }
  node := n - 1
  for node > 0 {
    tree[node] = tree[node*2] + tree[node*2+1]
  }
}
```

### Update the segment tree
Update a value is very straightforward as we will just update it and propagate the updates from bottom to up. The following code assumes we are calculating sum.
```
update(i val) {
    // take the update.
    tree[i+n] = val
    node = i + n
    // if node is not the root.
    for node > 1 {
      left, right = node, node
      if node % 2 == 0 {
        // if node is the left child.
        right += 1
      } else {
        left -= 1
      }
      // update parent
      tree[node/2] = sum(tree[left] + tree[right])
      node /= 2
    }
}
```

### Query the segment tree
The query operations query from i to j, which corresponds to the n+i and n+j leaf nodes. We call it left and right.
As long as left <= right, we can keep looking back to their parents. But there are two cases we need to take care of:
- If left is the right child, then we can't take its parent's value as the left child is out of the scope.
- If right is the left child, then we can't take its parent's value as the right child is out of the scope.
The following code assumes we are getting sum.
```
query(i, j) {
  left = n+i
  right = n+j
  for left <= right {
    if left % 2 == 1 {
      sum += tree[left]
      left++
    }
    if right % 2 == 0 {
      sum += tree[right]
      right--
    }
    // go to their parents
    left /= 2
    right /= 2
  }
  return sum
}
```

## Binary Indexed Tree
Imaging a scenario, you are given an array of N, and your task is to query the sum of from 0 to i and update the value of i, where i < N. How would you do it? The most naive way would be store the array of N as it is, then
- To get the sum from 0 to i, the time complexity will be O(N).
- To update a value at i, the time complexity will be O(1).
If you are going to update the array very frequently and rarely calculate the sum, this approach might be ok. However, if you wanna speed up the get sum operation, you could store the prefix-accumulated sum in another array B, so that B[i] = A[0] + ... + A[i], then
- To get the sum from 0 to 9, the time complexity will be O(1).
- To update a value at i, you have to update all elements in B[i] to B[N-1], thus the time complexity is O(N).

Is there a way to make the get sum and the update operation both faster? Yes, that's where Binary Indexed Tree comes in to play. Instead of storing the prefix-accumulated sum, the idea of Binary Indexed Tree is to store the postfix-accumulated sum in an array C, where:
```
C[i] = A[i] + A[i-1] + ... + A[i-lowbit(i)+1]
```

> Note that here i is 1-indexed, not 0-indexed.

So the key the `lowbit` function, what it does is, for a given integer x, turn it into binrary format, then turn all bits into 0, except for the first 1 encountered, from low bit to high bit. e.g. 6 => 0110 => lowbit(0110) => 0010.

If we visualize the Indexed Binrary Tree, we will see the following:

```                               
                                  C[8]
                                /   |
                              /     |
                            /       |
                          /         |
                        /           |
                      /             |
                    /               |
                  /                ...
              C[4]              /   |
             /  |             /     |
           /    |           /       |
         /      |         /         |
     C[2]      ...      C[6]       ...
     /|       / |      /  |      /  |
C[1] ...  C[3] ...  C[5] ...  C[7] ...
 |    |    |    |    |    |    |    |
A[1] A[2] A[3] A[4] A[5] A[6] A[7] A[8]
```
And we can see:
```
C[1] = A[1]                                     (1 => 001 => lowbit(001) => 001, C[1] = A[1] + ... + A[1-1+1])
C[2] = A[2]+A[1]                                (2 => 010 => lowbit(010) => 010, C[2] = A[2] + ... + A[2-2+1])
C[3] = A[3]                                     (3 => 011 => lowbit(001) => 001, C[2] = A[3] + ... + A[3-1+1])
C[4] = A[4]+A[3]+A[2]+A[1]                      (4 => 100 => lowbit(100) => 100, C[4] = A[4] + ... + A[4-4+1])
C[5] = A[5]                                     (5 => 101 => lowbit(001) => 001, C[1] = A[5] + ... + A[5-1+1])
C[6] = A[6]+A[5]                                (6 => 110 => lowbit(010) => 010, C[6] = A[6] + ... + A[6-2+1])
C[7] = A[7]                                     (7 => 111 => lowbit(001) => 001, C[7] = A[7] + ... + A[7-1+1])
C[8] = A[8]+A[7]+A[6]+A[5]+A[4]+A[3]+A[2]+A[1]  (8 => 1000 => lowbit(1000) => 1000, C[8] = A[8] + ... + A[8-8+1])
```
### Build the Binary Indexed Tree
Building the tree essentially is to construct the array C from the given input A.
```
build(C, A, N) {
  for i = 1; i <= N; i++ {
    // from tree bottom, update to up.
    val = A[i]
    for i <= N {
      C[i] += val
      i += lowbit(i)
    }
  }
}
// A and C is 1-indexed.
```
The time complexity is O(NlogN)

### Get Sum using the Binary Indexed Tree
```
getSum(C, j) int {
  res = 0
  for j > 0 {
    res += C[j]
    j -= lowbit(j)
  }
}
// C is 1-indexed
```
The time complexity is O(logN)

### Update a idx by delta using the Binary Indexed Tree
```
update(C, i, delta) {
  for i <= N {
    C[i] += delta
    i += lowbit(1)
  }
}
```
The time complexity is O(logN)

## Trie Tree
Trie Tree is also called **Radix Tree**, or **Prefix Tree**. It's a search tree that can be used store a collection of strings and provides the following functionality:
- Insert a string into the trie tree.
- Search exactly a string from the trie tree.
- Count the occurrence of strings.
- Sort strings.
- Auto-completion, as in, for a given prefix, find out all strings that share the same prefix.

We can define the Trie Tree as follows:
- The root node does not contain any character, it represents "".
- Each other node has only one character, assembling all the characters from the root node down to this node forms a prefix string.
- Each non-leaf node contains 0 - 26 children, representing a-z.
- All the strings that are represented by the children a node, share the the same prefix mentioned above.
- Do a pre-order traversing of the trie tree can prints all strings that are stored in it.

The following gives an example:
```
                                    root
                                    /  \
                                  'a'  'b'
                                  /      \
                                 'b'     'e'
                                /        / \
                               'u'      'e' 'd'
                              /
                            's'
                            /
                           'e'
```
The above trie tree represents the collection of strings ["abuse", "bee", "bed", "be"].
> Please note that, to be able to have "be" marked as a string, we need a flag in the tree node.

### Inserting into a trie tree.
Inserting into a trie is a process of a travering down the tree based on the string to be inserted.
- for each character in the string, 0 <= i < N, where N is the length of the string.
- At a given node, if it has a child that matches the current character, move the node to this child.
- At a given ndoe, if it does not have a child that matches the current character, create the child with the value of this character and move the node to this child.
- repeat until all character of the strings are inserted, when reaching the last character, mark the node using a flag to indicate it represents a string, optionally, we can have a count field in a node to count occurrence of the string.

> Storing all children nodes for a node using [26]\*TreeNode is a trick, because we only have a-z in alphabet.

### Search from a trie tree.
For a given prefix, we can find out all the strings that started with it. If prefix == "", it means all the strings.
- Starting from the root, we will firstly search down the tree to find out whether prefix exists by matching character by character.
- If prefix does not exist, return empty.
- If find the prefix, at a node, we will print all the strings represented by this node's subtrees.
- In order to achieve it, we can use a pre-order DFS traverse.

With the above operations, we can achieve what we want using trie tree.
