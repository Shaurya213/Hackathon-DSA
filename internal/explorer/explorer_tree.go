package explorer

import (
	"fmt"
	"playwise-cli/internal/models"
)

// TreeNode represents a generic classification node
type TreeNode struct {
	Name     string
	Children map[string]*TreeNode
	Songs    []*models.Song
	Level    string // Genre, Subgenre, Mood, Artist
}

// ExplorerTree is the root of the tree
type ExplorerTree struct {
	Root *TreeNode
}

// NewExplorerTree initializes the root node
func NewExplorerTree() *ExplorerTree {
	return &ExplorerTree{
		Root: &TreeNode{
			Name:     "Music",
			Children: make(map[string]*TreeNode),
			Level:    "Root",
		},
	}
}

// AddSong inserts a song based on its hierarchy
func (et *ExplorerTree) AddSong(genre, subgenre, mood, artist string, song *models.Song) {
	curr := et.Root

	levels := []struct {
		Name  string
		Label string
	}{
		{genre, "Genre"},
		{subgenre, "Subgenre"},
		{mood, "Mood"},
		{artist, "Artist"},
	}

	for _, level := range levels {
		if _, exists := curr.Children[level.Name]; !exists {
			curr.Children[level.Name] = &TreeNode{
				Name:     level.Name,
				Children: make(map[string]*TreeNode),
				Level:    level.Label,
			}
		}
		curr = curr.Children[level.Name]
	}
	curr.Songs = append(curr.Songs, song)
}

// DFS traversal to print tree with indentation
func (et *ExplorerTree) PrintTree() {
	printNode(et.Root, 0)
}

func printNode(node *TreeNode, depth int) {
	if node == nil {
		return
	}
	indent := ""
	for i := 0; i < depth; i++ {
		indent += "  "
	}
	fmt.Printf("%s📂 %s (%s)\n", indent, node.Name, node.Level)

	if len(node.Songs) > 0 {
		for _, song := range node.Songs {
			fmt.Printf("%s  🎵 %s - %s\n", indent, song.Title, song.Artist)
		}
	}

	for _, child := range node.Children {
		printNode(child, depth+1)
	}
}
