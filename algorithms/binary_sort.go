package algorithms

import "gabriel/algos/structures"

/**
 * Put the input array in a BinarySearchTree (BST), element by element.
 * By definition, the BST will keep all elements sorted.
 * Run through the BST to produce the sorted array.
 *
 * Performance:
 * We'll have to go through each of the N elements of the input array and insert them in the BST.
 * Each insert will roughly take logN.
 * And the final run through will also take N.
 * So roughly O(NlogN), which isn't super bad but not super good either.
 * Obviously all caveats of a BST apply here (if data is already sorted, inserting will take N, etc).
 * And it's going to require a BST with N space complexity, alongside the initial array.
 */
func BinarySort(input []int) []int {
	bst := structures.BinarySearchTree[int]()
	for _, val := range input {
		bst.Insert(val)
	}
	return bst.ToArray()
}
