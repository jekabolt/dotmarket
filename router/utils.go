package api

import "github.com/Kutabe/vk"

func (s *Server) AuthVK() (*vk.AuthResponse, error) {
	return vk.Auth(s.Login, s.Password)
}
