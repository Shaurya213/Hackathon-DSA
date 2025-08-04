package recommend

import (
	"fmt"
	"math"
	"playwise-cli/internal/models"
	"time"
)

// Recommender stores all songs and recent play history
type Recommender struct {
	AllSongs        []*models.Song
	PlaybackHistory []*models.Song
	RecentLimit     time.Duration // e.g., 24h or 10m
}

// NewRecommender creates a new engine
func NewRecommender() *Recommender {
	return &Recommender{
		AllSongs:        make([]*models.Song, 0),
		PlaybackHistory: make([]*models.Song, 0),
		RecentLimit:     24 * time.Hour,
	}
}

// AddToHistory logs a song as played
func (r *Recommender) AddToHistory(song *models.Song) {
	r.PlaybackHistory = append(r.PlaybackHistory, song)
}

// AddToLibrary adds a new song to the global pool
func (r *Recommender) AddToLibrary(song *models.Song) {
	r.AllSongs = append(r.AllSongs, song)
}

// Recommend returns songs that are similar but not recently played
func (r *Recommender) Recommend(reference *models.Song) []*models.Song {
	recentSet := make(map[string]bool)
	thresholdTime := time.Now().Add(-r.RecentLimit)

	for _, song := range r.PlaybackHistory {
		if song.AddedAt.After(thresholdTime) {
			recentSet[song.ID] = true
		}
	}

	recommendations := make([]*models.Song, 0)
	for _, song := range r.AllSongs {
		if song.ID == reference.ID || recentSet[song.ID] {
			continue // skip same or recently played
		}

		if isSimilar(reference, song) {
			recommendations = append(recommendations, song)
		}
	}

	return recommendations
}

// isSimilar checks if 2 songs are similar by Genre, BPM, and Duration
func isSimilar(a, b *models.Song) bool {
	sameGenre := a.Genre == b.Genre
	bpmClose := math.Abs(float64(a.BPM-b.BPM)) <= 10
	durationClose := math.Abs(float64(a.Duration-b.Duration)) <= 15

	return sameGenre && bpmClose && durationClose
}

// PrintRecommendations displays results
func (r *Recommender) PrintRecommendations(ref *models.Song) {
	fmt.Printf("\n🤖 Recommendations similar to '%s' by %s:\n", ref.Title, ref.Artist)
	recs := r.Recommend(ref)

	if len(recs) == 0 {
		fmt.Println("⚠️ No suitable recommendations found.")
		return
	}
	for _, song := range recs {
		fmt.Printf(" - %s by %s [BPM: %d, Duration: %ds]\n", song.Title, song.Artist, song.BPM, song.Duration)
	}
}
