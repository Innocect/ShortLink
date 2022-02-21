package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ashu/handler"
	"github.com/ashu/redisDao"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()
	redisClient := redisDao.RedisConnection()
	if redisClient == nil {
		log.Panic("Error initialising Redis")
	}

	r.HandleFunc("/assignment/shorturl", handler.GetHandler(redisClient)).Methods(http.MethodGet)
	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"}, //you service is available and allowed for this base url
		AllowedMethods: []string{
			http.MethodGet, //http methods for your app
			http.MethodPost,
			http.MethodPatch,
			http.MethodOptions,
			http.MethodHead,
		},

		AllowedHeaders: []string{
			"*", //or you can your header key values which you are using in your application

		},
	})

	server := &http.Server{
		Handler:      corsOpts.Handler(r),
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 2 * time.Second,
		ReadTimeout:  2 * time.Second,
	}

	fmt.Println("Starting the API server on ", "https://"+server.Addr, " ........")
	log.Fatal(server.ListenAndServe())
}
