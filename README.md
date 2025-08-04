# 🎿 PlayWise CLI — DSA Hackathon Project

PlayWise is a CLI-based smart playlist manager built in Go as part of a DSA-focused hackathon challenge. It demonstrates real-world applications of core data structures like Linked Lists, Stacks, BSTs, Hash Maps, Trees, and Sorting algorithms.

---

## 🚀 Features Implemented

| ID  | Feature                                | DSA Concept            |
| --- | -------------------------------------- | ---------------------- |
| 1️⃣ | Playlist Engine                        | Doubly Linked List     |
| 2️⃣ | Playback History (Undo)                | Stack                  |
| 3️⃣ | Rating-Based Song Search               | Binary Search Tree     |
| 4️⃣ | Instant Song Lookup                    | HashMap                |
| 5️⃣ | Playlist Sorting (title/duration/date) | Merge Sort             |
| 6️⃣ | Playback Optimization                  | Time-Space Annotation  |
| 7️⃣ | Snapshot Dashboard                     | Aggregation + Map      |
| 8️⃣ | Playlist Explorer Tree                 | N-ary Tree             |
| 9️⃣ | Smart Song Recommendation              | Filtering + Similarity |

---

## 💂 Folder Structure

```
playwise-cli/
├── cmd/
│   └── main.go                 # CLI entrypoint
├── internal/
│   ├── models/                 # Shared Song struct
│   ├── playlist/               # Linked list playlist
│   ├── playback/               # Stack for playback history
│   ├── ratingbst/              # BST for ratings
│   ├── lookup/                 # HashMap lookup
│   ├── sorting/                # Merge sort engine
│   ├── snapshot/               # System dashboard
│   ├── explorer/               # Genre → Artist explorer tree
│   └── recommend/              # Smart recommender
├── testdata/
│   └── sample_songs.json       # Sample songs for loading
├── go.mod
├── go.sum
└── README.md
```

---

## 🧪 Running the App

### 1. Clone and enter the project

```bash
git clone https://github.com/Shaurya213/Hackathon-DSA.git
cd Hackathon-DSA
```

### 2. Run the CLI

```bash
go run cmd/main.go
```

### 3. CLI Menu

```
🎿 PlayWise CLI Menu:
1. Add Song to Playlist
2. Delete Song by Index
3. Move Song
4. Reverse Playlist
5. Display Playlist
6. Play Song by Index
7. View Playback History
8. Show Snapshot Dashboard
9. Recommend Similar Songs
0. Exit
```

---

## 📦 Sample Song Data

Located in `testdata/sample_songs.json`:

```json
[
  {
    "id": "s1",
    "title": "Bohemian Rhapsody",
    "artist": "Queen",
    "duration": 355,
    "genre": "Rock",
    "bpm": 72,
    "rating": 5
  },
  ...
]
```

---

## 📚 Concepts Practiced

* Pointer manipulation and traversal
* Real-world data modeling using structs
* Map-sync between data structures
* Recursion and traversal (DFS, BFS)
* Space-time tradeoff analysis
* Command-line interactivity

---

## 👨‍💼 Author

**Shaurya Gupta**
Golang Developer | Backend Enthusiast | DSA Explorer
GitHub: [@Shaurya213](https://github.com/Shaurya213)

---

## 📌 Note

This CLI project is intentionally kept modular for academic and demo purposes.
Each module is designed to reflect a distinct DSA concept mapped to a real feature.
