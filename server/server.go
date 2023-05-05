package server

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/awilson506/med-bridge-hep/api"
	"github.com/gorilla/mux"
)

// hold some clients for the server to use
type Server struct {
	client *api.Client
	server *http.Server
	mux    *mux.Router
}

// for some added logging...
type Logger struct {
	handler http.Handler
}

// ServeHTTP
func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	l.handler.ServeHTTP(w, r)
	log.Println("handling route: ", r.URL.Path)
}

// NewServer get a new instance of the server
func NewServer(db *sql.DB) *Server {
	s := &Server{
		client: api.NewClient(db),
		mux:    mux.NewRouter(),
	}

	s.server = &http.Server{
		Addr:    ":8080",
		Handler: &Logger{s.mux},
	}

	// create the group and routes
	v1 := s.mux.PathPrefix("/v1").Subrouter()

	// TODO: add roles to users and scope the requests to user based on patient and or therapist ids/auth
	v1.HandleFunc("/exercises", s.client.GetExercises()).Methods("GET")
	v1.HandleFunc("/exercises/{id}", s.client.GetExercise()).Methods("GET")
	v1.HandleFunc("/exercises", s.client.CreateExercise()).Methods("POST")
	v1.HandleFunc("/exercises/{id}", s.client.UpdateExercise()).Methods("PUT")

	v1.HandleFunc("/sessions", s.client.GetSessions()).Methods("GET")
	v1.HandleFunc("/sessions/{id}", s.client.GetSession()).Methods("GET")
	v1.HandleFunc("/sessions", s.client.CreateSession()).Methods("POST")
	v1.HandleFunc("/sessions/{id}", s.client.UpdateSession()).Methods("PUT")

	v1.HandleFunc("/patient-program/{programId}", s.client.GetPatientExerciseProgram()).Methods("GET")
	//TODO: create the rest of the handlers for patients

	return s
}

// start up the server
func (s *Server) Start() error {
	return s.server.ListenAndServe()
}
