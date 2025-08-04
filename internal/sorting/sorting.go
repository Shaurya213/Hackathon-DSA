package sorting

import (
	"playwise-cli/internal/models"
)

// SortByTitle sorts the songs alphabetically by title
func SortByTitle(songs []*models.Song, ascending bool) []*models.Song {
	return mergeSort(songs, func(a, b *models.Song) bool {
		if ascending {
			return a.Title < b.Title
		}
		return a.Title > b.Title
	})
}

// SortByDuration sorts the songs by duration
func SortByDuration(songs []*models.Song, ascending bool) []*models.Song {
	return mergeSort(songs, func(a, b *models.Song) bool {
		if ascending {
			return a.Duration < b.Duration
		}
		return a.Duration > b.Duration
	})
}

// SortByRecentlyAdded sorts songs by AddedAt timestamp (latest first)
func SortByRecentlyAdded(songs []*models.Song) []*models.Song {
	return mergeSort(songs, func(a, b *models.Song) bool {
		return a.AddedAt.After(b.AddedAt)
	})
}

// Generic merge sort using comparator
func mergeSort(arr []*models.Song, less func(a, b *models.Song) bool) []*models.Song {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid], less)
	right := mergeSort(arr[mid:], less)

	return merge(left, right, less)
}

func merge(left, right []*models.Song, less func(a, b *models.Song) bool) []*models.Song {
	result := make([]*models.Song, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if less(left[i], right[j]) {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}
