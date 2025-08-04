package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"playwise-cli/internal/models"
	"playwise-cli/internal/playback"
	"playwise-cli/internal/playlist"
	"playwise-cli/internal/recommend"
	"playwise-cli/internal/snapshot"
)

func main() {
	pl := playlist.NewPlaylist()
	reader := bufio.NewReader(os.Stdin)

	// Integrated modules
	pb := playback.NewPlaybackStack()
	snap := snapshot.NewSnapshotEngine()
	rec := recommend.NewRecommender()

	// Load songs from JSON
	err := loadSampleData(pl, snap, rec, "testdata/sample_songs.json")
	if err != nil {
		fmt.Println("Failed to load sample data:", err)
	}

	for {
		fmt.Println("\n🎧 PlayWise CLI Menu:")
		fmt.Println("1. Add Song to Playlist")
		fmt.Println("2. Delete Song by Index")
		fmt.Println("3. Move Song")
		fmt.Println("4. Reverse Playlist")
		fmt.Println("5. Display Playlist")
		fmt.Println("6. Play Song by Index")
		fmt.Println("7. View Playback History")
		fmt.Println("8. Show Snapshot Dashboard")
		fmt.Println("9. Recommend Similar Songs")
		fmt.Println("0. Exit")

		fmt.Print("Enter choice: ")
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var id, title, artist string
			var duration int

			fmt.Print("Enter Song ID: ")
			fmt.Scan(&id)

			fmt.Print("Enter Title: ")
			reader.ReadString('\n')
			title, _ = reader.ReadString('\n')
			title = trim(title)

			fmt.Print("Enter Artist: ")
			artist, _ = reader.ReadString('\n')
			artist = trim(artist)

			fmt.Print("Enter Duration (in sec): ")
			fmt.Scan(&duration)

			song := &models.Song{ID: id, Title: title, Artist: artist, Duration: duration}
			pl.AddSong(song)
			rec.AddToLibrary(song)
			snap.AddSongToSnapshot(song)
			fmt.Println("✅ Song added!")

		case 2:
			pl.Display()
			var index int
			fmt.Print("Enter index to delete: ")
			fmt.Scan(&index)
			pl.DeleteSong(index)

		case 3:
			pl.Display()
			var from, to int
			fmt.Print("Move from index: ")
			fmt.Scan(&from)
			fmt.Print("Move to index: ")
			fmt.Scan(&to)
			pl.MoveSong(from, to)

		case 4:
			pl.Reverse()
			fmt.Println("🔄 Playlist reversed.")

		case 5:
			pl.Display()

		case 6:
			pl.Display()
			var index int
			fmt.Print("Enter index to play: ")
			fmt.Scan(&index)
			song := pl.GetSongAt(index)
			if song == nil {
				fmt.Println("❌ Invalid index.")
				break
			}
			fmt.Printf("🎶 Now playing: %s by %s\n", song.Title, song.Artist)
			pb.Push(song)
			snap.LogPlayback(song)
			rec.AddToHistory(song)

		case 7:
			pb.Display()

		case 8:
			snap.ExportSnapshot()

		case 9:
			pl.Display()
			var index int
			fmt.Print("Pick reference index: ")
			fmt.Scan(&index)
			ref := pl.GetSongAt(index)
			if ref == nil {
				fmt.Println("❌ Invalid index.")
				break
			}
			rec.PrintRecommendations(ref)

		case 0:
			fmt.Println("👋 Exiting PlayWise CLI.")
			return

		default:
			fmt.Println("❌ Invalid choice")
		}
	}
}

func trim(s string) string {
	return strings.TrimSpace(s)
}

func loadSampleData(pl *playlist.Playlist, snap *snapshot.SnapshotEngine, rec *recommend.Recommender, filepathStr string) error {
	fmt.Println("📂 Trying to load:", filepathStr)

	file, err := os.Open(filepath.Clean(filepathStr))
	if err != nil {
		return err
	}
	defer file.Close()

	var songs []*models.Song
	if err := json.NewDecoder(file).Decode(&songs); err != nil {
		return err
	}

	fmt.Printf("✅ Loaded %d songs from JSON\n", len(songs))

	for _, song := range songs {
		fmt.Printf("🎵 Adding song: %+v\n", song)
		pl.AddSong(song)
		snap.AddSongToSnapshot(song)
		rec.AddToLibrary(song)
	}

	fmt.Printf("🎼 Playlist Size after loading: %d\n", pl.Size)
	return nil
}
