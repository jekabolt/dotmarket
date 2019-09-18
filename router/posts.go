package api

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/Kutabe/vk"
)

func (s *Server) GetAllPostsChan(postType string) <-chan Post {
	//TODO: add errors
	// make return channel
	output := make(chan Post)

	parameters := make(map[string]string)
	random_id := rand.Intn(999999999999999-100000000000000) + 100000000000000
	parameters["random_id"] = strconv.Itoa(random_id)
	parameters["v"] = "5.101" // VK API version
	parameters["owner_id"] = s.GroupID
	parameters["filter"] = postType

	go func() {
		resp, err := vk.Request("wall.get", parameters, s.User)
		if err != nil {
			log.Printf("GetAllPosts:vk.Request:err [%v]", err.Error())

		}

		var response Response

		if err := json.Unmarshal(resp, &response); err != nil {
			log.Printf("GetAllPosts:json.Unmarshal:err [%v]", err.Error())
		}

		// posts := []Post{}
		for _, item := range response.Data.Items {
			datePublished := time.Time{}
			dateScheduled := time.Time{}
			postOwner := ""

			if postType == "suggests" {
				datePublished = time.Unix(int64(item.Date), 0)
				dateScheduled = time.Unix(int64(0), 0)
				postOwner = fmt.Sprintf("https://vk.com/id%d", item.FromID)
			}

			if postType == "postponed" {
				dateScheduled = time.Unix(int64(item.Date), 0)
				postOwner = fmt.Sprintf("https://vk.com/id%d", -item.FromID)
			}

			post := Post{
				DatePublished: datePublished,
				DateScheduled: dateScheduled,
				PostOwner:     postOwner,
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

			output <- post

		}
		close(output)
	}()

	return output
}

func merge(outputsChan ...<-chan Post) <-chan Post {

	var wg sync.WaitGroup

	merged := make(chan Post)

	wg.Add(len(outputsChan))

	output := func(sc <-chan Post) {
		for sqr := range sc {
			merged <- sqr
		}
		wg.Done()
	}

	for _, optChan := range outputsChan {
		go output(optChan)
	}

	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}
