package api

import (
	"encoding/json"
	"example/bot-frases/utilis"
	"net/http"
)

type Server struct {
	listenAddr string
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/phrase", s.handleGetRandomPhrase)
	return http.ListenAndServe(s.listenAddr, nil)
}

func (s *Server) handleGetRandomPhrase(w http.ResponseWriter, r *http.Request) {
	phrase := utilis.BuildPhrase()
	json.NewEncoder(w).Encode(phrase)
}
