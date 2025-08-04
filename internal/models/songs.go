package models

import "time"

type Song struct {
	ID       string // Unique ID for lookup
	Title    string
	Artist   string
	Duration int       // Duration in seconds
	AddedAt  time.Time // For sorting recently added
	Genre    string    // For recommendations
	BPM      int       // Beats Per Minute, used in similarity
	Rating   int       // 1 to 5 stars
}
