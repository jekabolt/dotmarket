package api

import "time"

type Item struct {
	ID          int    `json:"id"`
	FromID      int    `json:"from_id"`
	OwnerID     int    `json:"owner_id"`
	Date        int    `json:"date"`
	Type        string `json:"type"`
	Text        string `json:"text"`
	Attachments []struct {
		Type  string `json:"type"`
		Photo Photo  `json:"photo"`
	} `json:"attachments"`
}

type Photo struct {
	ID      int `json:"id"`
	AlbumID int `json:"album_id"`
	OwnerID int `json:"owner_id"`
	UserID  int `json:"user_id"`
	Sizes   []struct {
		Type   string `json:"type"`
		URL    string `json:"url"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
	} `json:"sizes"`
	Text      string `json:"text"`
	Date      int    `json:"date"`
	AccessKey string `json:"access_key"`
}

type Response struct {
	Data struct {
		Count    int    `json:"count"`
		Items    []Item `json:"items"`
		NextFrom string `json:"next_from"`
	} `json:"response"`
}

type Post struct {
	DatePublished time.Time
	DateScheduled time.Time
	PostOwner     string
	Text          string
	Images        []string
}
