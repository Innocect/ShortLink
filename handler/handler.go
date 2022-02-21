package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ashu/redisDao"
)

func GetHandler(resp http.ResponseWriter, req *http.Request) {
	redisClient := redisDao.RedisConnection()
	if redisClient == nil {
		log.Fatal("Error initialising redis")
	}
	fmt.Println("Main short Link Logic")
}
