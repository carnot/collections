# Badgerodon Collections

Maps and slices go a long way in Go, but sometimes you need more. This is a collection of collections that may be useful.

## Splay Tree

A [splay tree](http://en.wikipedia.org/wiki/Splay_tree) is a type of binary search tree where every access to the tree results in the tree being rearranged so that the current node gets put on top.

## Ternary Search Tree

A [ternary search tree](http://en.wikipedia.org/wiki/Ternary_search_tree) is similar to a trie in that nodes store the letters of the key, but instead of either using a list or hash at each node a binary tree is used. Ternary search trees have the performance benefits of a trie without the usual memory costs.