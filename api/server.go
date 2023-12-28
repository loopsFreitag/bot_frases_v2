package api

import (
	"encoding/json"
	"example/bot-frases/utilis"
	"net/http"
	"strings"
	"unicode"
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
	http.Handle("/", http.FileServer(http.Dir("./static")))
	return http.ListenAndServe(s.listenAddr, nil)
}

func (s *Server) handleGetRandomPhrase(w http.ResponseWriter, r *http.Request) {
	phrase := utilis.BuildPhrase()

	lowercase := strings.ToLower(phrase)
	result := string(unicode.ToUpper(rune(lowercase[0]))) + lowercase[1:]

	response := map[string]interface{}{
		"message": result,
	}

	// Use json.NewEncoder to encode the map and write it to the response writer
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}
