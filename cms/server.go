package cms

import (
	"log"
	"net/http"
)

type Server struct {
	mux *http.ServeMux
}

func NewServer() *Server {
	return &Server{
		mux: http.NewServeMux(),
	}
}

func (s *Server) prepareContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer:  w,
		Request: r,
	}
}

func (s *Server) handle(url string, handler Handler) {
	s.mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		c := s.prepareContext(w, r)
		_ = handler(c) // TODO error handling
	})
}

func (s *Server) Run(addr string) error {
	// TODO display correct address
	log.Printf("Starting server on http://localhost%s\n", addr)
	return http.ListenAndServe(addr, s.mux)
}
