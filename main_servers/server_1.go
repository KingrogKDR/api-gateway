package mainservers

import "github.com/KingrogKDR/api-gateway/models"

// main server that takes request from the reverse proxy (thinking of using two rate limiters for each server)
// what if there is another rate limiter for the reverse proxy
type Server struct {
	repository *models.Repo
}

func NewServer(repository *models.Repo) *Server {
	return &Server{repository: repository}
}

// Todo: add routes and methods to the server based on different types of request and also a start func to start the server
