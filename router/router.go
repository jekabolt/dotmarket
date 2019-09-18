package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Kutabe/vk"
	"github.com/caarlos0/env"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Server struct {
	Port     string `env:"SERVER_PORT" envDefault:"7005"`
	Login    string `env:"VK_LOGIN" envDefault:"jekabolt@gmail.ru"`
	Password string `env:"VK_PASSWORD" envDefault:"test"`
	GroupID  string `env:"VK_GROUP_ID" envDefault:"-170017193"`
	CertDir  string `env:"CERT_DIR" envDefault:""`
	User     *vk.AuthResponse
	Posts    Posts
}

type Posts struct {
	Postponed ChachePost
	Suggests  ChachePost
	All       ChachePost
}

type ChachePost struct {
	Posts  []Post
	Change time.Time
}

func InitServer() (*Server, error) {
	s := &Server{}
	err := env.Parse(s)

	user, err := s.AuthVK()
	if err != nil {
		return nil, err
	}
	if user.Error != "" {
		return nil, err
	} else {
		s.User = user
	}
	return s, err
}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}

func (s *Server) Serve() error {
	r := chi.NewRouter()

	// Init middlewares
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(15 * time.Second))

	r.Get("/health", s.healthCheck)

	workDir, _ := os.Getwd()
	fmt.Println("workDir ", workDir)
	filesDir := filepath.Join(workDir, "static")
	FileServer(r, "/", http.Dir(filesDir))
	// r.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("static"))))
	// r.NotFoundHandler = http.StripPrefix("/", http.FileServer(http.Dir("static/404.html")))

	r.Route("/api/v1.0", func(r chi.Router) {
		// can be "suggests" "postponed" "all"
		r.Get("/{postType}", s.getPosts)
	})

	if s.CertDir != "" {
		log.Println("Listening secure on :" + s.Port)
		return http.ListenAndServeTLS(":"+s.Port, s.CertDir+"/cert.pem", s.CertDir+"/privkey.pem", r)
	} else {
		log.Println("Listening on http://localhost:" + s.Port)
		return http.ListenAndServe(":"+s.Port, r)
	}

}
