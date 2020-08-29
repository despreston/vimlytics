package main

import (
	"github.com/despreston/vimlytics/handlers/github"
	"github.com/despreston/vimlytics/handlers/vimrc"
	"github.com/despreston/vimlytics/internal/api"
	"github.com/despreston/vimlytics/mongo"
	"github.com/despreston/vimlytics/redis"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application-json")
		next.ServeHTTP(w, r)
	})
}

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s\n", r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func main() {
	router := mux.NewRouter()
	router.Use(cors)
	router.Use(logger)

	router.Handle("/api/vimrc", api.Handler(vimrc.Upload)).Methods("POST")
	router.Handle("/api/github/vimrc", api.Handler(github.Vimrc)).Methods("GET")
	router.Handle("/api/github/user", api.Handler(github.User)).Methods("GET")

	// serve static content
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./web/dist")))

	addr := os.Getenv("ADDR")

	if len(addr) < 1 {
		addr = "localhost:3001"
	}

	srv := &http.Server{
		Handler:      router,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Ensure that the redis conn works
	redis.Client()

	// Ensure that the mongo conn works.
	mongo.Db()

	log.Printf("Listening at %s\n", addr)
	log.Fatal(srv.ListenAndServe())
}
