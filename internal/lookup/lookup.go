package lookup

import (
	"fmt"
	"playwise-cli/internal/models"
)

// LookupTable holds two maps for fast access
type LookupTable struct {
	ByID    map[string]*models.Song
	ByTitle map[string]*models.Song
}

// NewLookupTable creates a new lookup table
func NewLookupTable() *LookupTable {
	return &LookupTable{
		ByID:    make(map[string]*models.Song),
		ByTitle: make(map[string]*models.Song),
	}
}

// AddSong indexes a song by ID and title
func (lt *LookupTable) AddSong(song *models.Song) {
	lt.ByID[song.ID] = song
	lt.ByTitle[song.Title] = song
}

// GetByID returns the song with given ID
func (lt *LookupTable) GetByID(id string) *models.Song {
	song, exists := lt.ByID[id]
	if !exists {
		fmt.Println("🔍 Song ID not found.")
		return nil
	}
	return song
}

// GetByTitle returns the song with given title
func (lt *LookupTable) GetByTitle(title string) *models.Song {
	song, exists := lt.ByTitle[title]
	if !exists {
		fmt.Println("🔍 Song title not found.")
		return nil
	}
	return song
}

// DeleteSong removes a song from both maps
func (lt *LookupTable) DeleteSong(song *models.Song) {
	delete(lt.ByID, song.ID)
	delete(lt.ByTitle, song.Title)
}

// PrintAll displays all indexed songs
func (lt *LookupTable) PrintAll() {
	fmt.Println("📚 Lookup Table Contents:")
	for id, song := range lt.ByID {
		fmt.Printf(" - [%s] %s by %s\n", id, song.Title, song.Artist)
	}
}
