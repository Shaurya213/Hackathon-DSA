package ratingbst

import (
	"fmt"
	"playwise-cli/internal/models"
)

// TreeNode holds a list of songs for a specific rating
type TreeNode struct {
	Rating int
	Songs  []*models.Song
	Left   *TreeNode
	Right  *TreeNode
}

// RatingBST is the root of the BST
type RatingBST struct {
	Root *TreeNode
}

// NewRatingBST initializes an empty BST
func NewRatingBST() *RatingBST {
	return &RatingBST{}
}

// InsertSong adds a song under the correct rating node
func (bst *RatingBST) InsertSong(song *models.Song) {
	if song.Rating < 1 || song.Rating > 5 {
		fmt.Println("❌ Invalid rating. Must be between 1 and 5.")
		return
	}
	bst.Root = insert(bst.Root, song)
}

func insert(node *TreeNode, song *models.Song) *TreeNode {
	if node == nil {
		return &TreeNode{
			Rating: song.Rating,
			Songs:  []*models.Song{song},
		}
	}
	if song.Rating < node.Rating {
		node.Left = insert(node.Left, song)
	} else if song.Rating > node.Rating {
		node.Right = insert(node.Right, song)
	} else {
		node.Songs = append(node.Songs, song)
	}
	return node
}

// SearchByRating returns list of songs for a given rating
func (bst *RatingBST) SearchByRating(rating int) []*models.Song {
	node := bst.Root
	for node != nil {
		if rating < node.Rating {
			node = node.Left
		} else if rating > node.Rating {
			node = node.Right
		} else {
			return node.Songs
		}
	}
	return []*models.Song{}
}

// DeleteSong removes a song by its ID from the tree
func (bst *RatingBST) DeleteSong(songID string) {
	bst.Root = deleteSongByID(bst.Root, songID)
}

func deleteSongByID(node *TreeNode, id string) *TreeNode {
	if node == nil {
		return nil
	}

	// Traverse to find the song
	for i, song := range node.Songs {
		if song.ID == id {
			// Remove song from slice
			node.Songs = append(node.Songs[:i], node.Songs[i+1:]...)
			break
		}
	}

	// If no songs left at this rating, delete node
	if len(node.Songs) == 0 {
		if node.Left == nil {
			return node.Right
		}
		if node.Right == nil {
			return node.Left
		}
		// Replace with in-order successor
		successor := findMin(node.Right)
		node.Rating = successor.Rating
		node.Songs = successor.Songs
		node.Right = deleteRatingNode(node.Right, successor.Rating)
	}
	return node
}

func deleteRatingNode(node *TreeNode, rating int) *TreeNode {
	if node == nil {
		return nil
	}
	if rating < node.Rating {
		node.Left = deleteRatingNode(node.Left, rating)
	} else if rating > node.Rating {
		node.Right = deleteRatingNode(node.Right, rating)
	} else {
		if node.Left == nil {
			return node.Right
		}
		if node.Right == nil {
			return node.Left
		}
		temp := findMin(node.Right)
		node.Rating = temp.Rating
		node.Songs = temp.Songs
		node.Right = deleteRatingNode(node.Right, temp.Rating)
	}
	return node
}

func findMin(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

// PrintInOrder prints songs grouped by rating in sorted order
func (bst *RatingBST) PrintInOrder() {
	printInOrder(bst.Root)
}

func printInOrder(node *TreeNode) {
	if node == nil {
		return
	}
	printInOrder(node.Left)
	fmt.Printf("⭐ Rating %d:\n", node.Rating)
	for _, song := range node.Songs {
		fmt.Printf(" - %s by %s\n", song.Title, song.Artist)
	}
	printInOrder(node.Right)
}
