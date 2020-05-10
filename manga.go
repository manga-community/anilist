package anilist

import "strconv"

// Manga ...
type Manga struct {
	ID    int `json:"id"`
	MALID int `json:"idMal"`
	Title struct {
		Romaji  string `json:"romaji"`
		English string `json:"english"`
		Native  string `json:"native"`
	} `json:"title"`
}

// Link returns the permalink to that manga.
func (manga *Manga) Link() string {
	return "https://anilist.co/manga/" + strconv.Itoa(manga.ID)
}
