package snapshot

import (
	"fmt"
	"playwise-cli/internal/models"
	"sort"
)

// SnapshotEngine holds the state required for generating system snapshot
type SnapshotEngine struct {
	AllSongs        []*models.Song
	PlaybackHistory []*models.Song // Most recent at the end
}

// NewSnapshotEngine creates an empty snapshot engine
func NewSnapshotEngine() *SnapshotEngine {
	return &SnapshotEngine{
		AllSongs:        make([]*models.Song, 0),
		PlaybackHistory: make([]*models.Song, 0),
	}
}

// AddSongToSnapshot adds a song to the engine's global song list
func (s *SnapshotEngine) AddSongToSnapshot(song *models.Song) {
	s.AllSongs = append(s.AllSongs, song)
}

// LogPlayback adds a song to the recent playback history
func (s *SnapshotEngine) LogPlayback(song *models.Song) {
	s.PlaybackHistory = append(s.PlaybackHistory, song)
}

// TopLongestSongs returns the top N longest songs
func (s *SnapshotEngine) TopLongestSongs(n int) []*models.Song {
	if len(s.AllSongs) == 0 {
		return nil
	}

	sorted := make([]*models.Song, len(s.AllSongs))
	copy(sorted, s.AllSongs)

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Duration > sorted[j].Duration
	})

	if n > len(sorted) {
		n = len(sorted)
	}
	return sorted[:n]
}

// MostRecentlyPlayed returns the N most recently played songs
func (s *SnapshotEngine) MostRecentlyPlayed(n int) []*models.Song {
	l := len(s.PlaybackHistory)
	if n > l {
		n = l
	}
	return s.PlaybackHistory[l-n:]
}

// SongCountByRating returns a map of rating → count
func (s *SnapshotEngine) SongCountByRating() map[int]int {
	count := make(map[int]int)
	for _, song := range s.AllSongs {
		count[song.Rating]++
	}
	return count
}

// ExportSnapshot prints all snapshot data
func (s *SnapshotEngine) ExportSnapshot() {
	fmt.Println("\n📊 System Snapshot Dashboard")

	// Top 5 longest
	fmt.Println("\n🎶 Top 5 Longest Songs:")
	top := s.TopLongestSongs(5)
	for i, song := range top {
		fmt.Printf(" %d. %s - %s (%ds)\n", i+1, song.Title, song.Artist, song.Duration)
	}

	// Recently played
	fmt.Println("\n🕓 Recently Played Songs:")
	recent := s.MostRecentlyPlayed(5)
	for i := len(recent) - 1; i >= 0; i-- {
		song := recent[i]
		fmt.Printf(" - %s by %s\n", song.Title, song.Artist)
	}

	// Song count by rating
	fmt.Println("\n⭐ Song Count by Rating:")
	countMap := s.SongCountByRating()
	for rating := 1; rating <= 5; rating++ {
		fmt.Printf(" %d star: %d songs\n", rating, countMap[rating])
	}
}
