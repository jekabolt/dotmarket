package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi"
)

func (s *Server) healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (s *Server) getPosts(w http.ResponseWriter, r *http.Request) {

	postType := chi.URLParam(r, "postType")
	posts := []Post{}

	switch postType {
	case "postponed":
		change := time.Now().Unix() - s.Posts.Postponed.Change.Unix()
		if change > 30 && len(s.Posts.Postponed.Posts) > 0 {
			rawPosts, err := json.Marshal(s.Posts.Postponed.Posts)
			if err != nil {
				http.Error(w, "bad marshal", http.StatusBadRequest)
			}
			w.WriteHeader(http.StatusOK)
			w.Write(rawPosts)
			return
		}

		postCh := s.GetAllPostsChan(postType)
		for post := range postCh {
			posts = append(posts, post)
		}
		s.Posts.Postponed = ChachePost{posts, time.Now()}
	case "suggests":
		change := time.Now().Unix() - s.Posts.Suggests.Change.Unix()
		if change > 30 && len(s.Posts.Suggests.Posts) > 0 {
			rawPosts, err := json.Marshal(s.Posts.Suggests.Posts)
			if err != nil {
				http.Error(w, "bad marshal", http.StatusBadRequest)
			}
			w.WriteHeader(http.StatusOK)
			w.Write(rawPosts)
			return
		}
		postCh := s.GetAllPostsChan(postType)
		for post := range postCh {
			posts = append(posts, post)
		}
		s.Posts.Suggests = ChachePost{posts, time.Now()}
	case "all":
		change := time.Now().Unix() - s.Posts.All.Change.Unix()
		if change < 30 && len(s.Posts.All.Posts) > 0 {
			rawPosts, err := json.Marshal(s.Posts.All.Posts)
			if err != nil {
				http.Error(w, "bad marshal", http.StatusBadRequest)
			}
			w.WriteHeader(http.StatusOK)
			w.Write(rawPosts)
			return
		}
		postponedCh := s.GetAllPostsChan("postponed")
		suggestsCh := s.GetAllPostsChan("suggests")
		chanMergedPosts := merge(postponedCh, suggestsCh)

		for post := range chanMergedPosts {
			posts = append(posts, post)
		}
		s.Posts.All = ChachePost{posts, time.Now()}

	default:
		http.Error(w, "wrong postType", http.StatusBadRequest)
		return
	}

	rawPosts, err := json.Marshal(posts)
	if err != nil {

	}
	w.WriteHeader(http.StatusOK)
	w.Write(rawPosts)
}
