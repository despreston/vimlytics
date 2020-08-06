package main

import (
	"github.com/despreston/vimlytics/handlers/vimrc"
	"github.com/despreston/vimlytics/redis"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type")
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func main() {
	router := mux.NewRouter()
	router.Use(cors)
	router.Use(logger)

	router.HandleFunc("/api/vimrc", vimrc.Post).Methods("POST")

	// serve static content
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./web/dist")))

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:3001",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Ensure that the redis conn works
	redis.Client()

	log.Println("Listening at :3001")
	log.Fatal(srv.ListenAndServe())
}
