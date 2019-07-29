package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/Kutabe/vk"
	"github.com/caarlos0/env"
)

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

type Config struct {
	Port     string `env:"SERVER_PORT" envDefault:"7005"`
	Login    string `env:"VK_LOGIN" envDefault:""`
	Password string `env:"VK_PASSWORD" envDefault:""`
	GroupID  string `env:"VK_GROUP_ID" envDefault:"-170017193"`
}

func (c *Config) GetAllPosts(postType string, user *vk.AuthResponse) ([]Post, error) {
	parameters := make(map[string]string)

	random_id := rand.Intn(999999999999999-100000000000000) + 100000000000000
	parameters["random_id"] = strconv.Itoa(random_id)
	parameters["v"] = "5.101" // VK API version
	parameters["owner_id"] = c.GroupID
	parameters["filter"] = postType

	resp, err := vk.Request("wall.get", parameters, user)
	if err != nil {
		log.Printf("GetAllPosts:vk.Request:err [%v]", err.Error())
		return nil, err
	}

	var response Response

	if err := json.Unmarshal(resp, &response); err != nil {
		log.Printf("GetAllPosts:json.Unmarshal:err [%v]", err.Error())
		return nil, err
	}

	posts := []Post{}
	for _, item := range response.Data.Items {

		datePublished := time.Time{}
		dateScheduled := time.Time{}
		if postType == "suggests" {
			datePublished = time.Unix(int64(item.Date), 0)
			dateScheduled = time.Unix(int64(0), 0)
		}

		if postType == "postponed" {
			dateScheduled = time.Unix(int64(item.Date), 0)
		}

		post := Post{
			DatePublished: datePublished,
			DateScheduled: dateScheduled,
			PostOwner:     fmt.Sprintf("https://vk.com/id%d", item.FromID),
			Text:          item.Text,
		}

		for _, attachment := range item.Attachments {
			if attachment.Type == "photo" {
				for _, size := range attachment.Photo.Sizes {
					if size.Type == "z" {
						post.Images = append(post.Images, size.URL)
					}
				}
			}
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (c *Config) AuthVK() (*vk.AuthResponse, error) {
	return vk.Auth(c.Login, c.Password)
}

func main() {

	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		log.Fatalf("main:env.Parse: [%v]", err.Error())
	}

	user, err := cfg.AuthVK()
	if err != nil {
		log.Fatalf("main:cfg.AuthVK: [%v]", err.Error())
	}

	// suggests
	// postponed

	if user.Error != "" {
		log.Fatalf("main:user.Error: [%v]", user.Error)
	} else {
		postsScheduled, err := cfg.GetAllPosts("postponed", user)
		if err != nil {
			//TODO: retry with new user
			log.Fatalf("main:cfg.GetAllPosts postponed: [%v]", err.Error())
		}

		// postsSuggests, err := cfg.GetAllPosts("postponed", user)
		// if err != nil {
		// 	//TODO: retry with new user
		// 	log.Fatalf("main:cfg.GetAllPosts postponed: [%v]", err.Error())
		// }

		for _, p := range postsScheduled {
			fmt.Println(p, "\n")
		}
	}
}
