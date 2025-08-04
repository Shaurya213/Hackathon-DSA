package playback

import (
	"fmt"
	"playwise-cli/internal/models"
)

// PlaybackStack holds recently played songs (LIFO)
type PlaybackStack struct {
	stack []*models.Song
}

// NewPlaybackStack initializes an empty playback stack
func NewPlaybackStack() *PlaybackStack {
	return &PlaybackStack{
		stack: make([]*models.Song, 0),
	}
}

// Push adds a song to the playback stack
func (ps *PlaybackStack) Push(song *models.Song) {
	ps.stack = append(ps.stack, song)
}

// Pop removes and returns the last played song (undo)
func (ps *PlaybackStack) Pop() *models.Song {
	n := len(ps.stack)
	if n == 0 {
		fmt.Println("⚠️ No song to undo.")
		return nil
	}
	last := ps.stack[n-1]
	ps.stack = ps.stack[:n-1]
	return last
}

// Peek returns the last played song without removing it
func (ps *PlaybackStack) Peek() *models.Song {
	n := len(ps.stack)
	if n == 0 {
		return nil
	}
	return ps.stack[n-1]
}

// IsEmpty checks if the stack is empty
func (ps *PlaybackStack) IsEmpty() bool {
	return len(ps.stack) == 0
}

// Display shows stack contents
func (ps *PlaybackStack) Display() {
	if ps.IsEmpty() {
		fmt.Println("🎵 No recent playback history.")
		return
	}
	fmt.Println("🕓 Playback History:")
	for i := len(ps.stack) - 1; i >= 0; i-- {
		s := ps.stack[i]
		fmt.Printf(" %d. %s - %s\n", len(ps.stack)-i, s.Title, s.Artist)
	}
}
