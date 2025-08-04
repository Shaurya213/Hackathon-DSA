package playlist

import (
	"fmt"
	"time"

	"playwise-cli/internal/models"
)

// Node represents a song node in the doubly linked list
type Node struct {
	Song *models.Song
	Prev *Node
	Next *Node
}

// Playlist manages the doubly linked list of songs
type Playlist struct {
	Head *Node
	Tail *Node
	Size int
}

// NewPlaylist creates and returns an empty playlist
func NewPlaylist() *Playlist {
	return &Playlist{}
}

// AddSong appends a song to the end of the playlist
func (pl *Playlist) AddSong(song *models.Song) {
	fmt.Printf("👉 Inserting into playlist: %s by %s\n", song.Title, song.Artist)
	song.AddedAt = time.Now()
	newNode := &Node{Song: song}

	if pl.Head == nil {
		pl.Head = newNode
		pl.Tail = newNode
	} else {
		pl.Tail.Next = newNode
		newNode.Prev = pl.Tail
		pl.Tail = newNode
	}
	pl.Size++
}

// DeleteSong removes a song at a given index
func (pl *Playlist) DeleteSong(index int) {
	if index < 0 || index >= pl.Size {
		fmt.Println("❌ Invalid index")
		return
	}

	curr := pl.Head
	for i := 0; i < index; i++ {
		curr = curr.Next
	}

	if curr.Prev != nil {
		curr.Prev.Next = curr.Next
	} else {
		pl.Head = curr.Next
	}

	if curr.Next != nil {
		curr.Next.Prev = curr.Prev
	} else {
		pl.Tail = curr.Prev
	}

	pl.Size--
}

// MoveSong moves a song from one index to another
func (pl *Playlist) MoveSong(from, to int) {
	if from < 0 || from >= pl.Size || to < 0 || to >= pl.Size {
		fmt.Println("❌ Invalid indices")
		return
	}
	if from == to {
		return
	}

	// Step 1: Detach node at `from`
	curr := pl.Head
	for i := 0; i < from; i++ {
		curr = curr.Next
	}
	// Remove curr from list
	if curr.Prev != nil {
		curr.Prev.Next = curr.Next
	} else {
		pl.Head = curr.Next
	}
	if curr.Next != nil {
		curr.Next.Prev = curr.Prev
	} else {
		pl.Tail = curr.Prev
	}

	// Step 2: Re-insert at `to`
	target := pl.Head
	for i := 0; i < to; i++ {
		target = target.Next
	}
	if to == 0 {
		curr.Next = pl.Head
		curr.Prev = nil
		pl.Head.Prev = curr
		pl.Head = curr
	} else {
		curr.Prev = target.Prev
		curr.Next = target
		target.Prev.Next = curr
		target.Prev = curr
	}
}

// Reverse reverses the playlist in place
func (pl *Playlist) Reverse() {
	curr := pl.Head
	var prev *Node
	pl.Tail = pl.Head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		curr.Prev = next
		prev = curr
		curr = next
	}
	pl.Head = prev
}

// Display prints the playlist
func (pl *Playlist) Display() {

	fmt.Println("📏 Playlist size:", pl.Size)

	// rest of the print logic...

	if pl.Head == nil {
		fmt.Println("🎧 Playlist is empty.")
		return
	}
	curr := pl.Head
	fmt.Println("🎵 Playlist:")
	i := 0
	for curr != nil {
		s := curr.Song
		fmt.Printf(" [%d] %s - %s (%ds)\n", i, s.Title, s.Artist, s.Duration)
		curr = curr.Next
		i++
	}
}

func (pl *Playlist) GetSongAt(index int) *models.Song {
	if index < 0 || index >= pl.Size {
		return nil
	}
	curr := pl.Head
	for i := 0; i < index; i++ {
		curr = curr.Next
	}
	if curr != nil {
		return curr.Song
	}
	return nil
}
