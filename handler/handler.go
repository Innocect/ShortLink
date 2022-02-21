package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ashu/model"
	"github.com/ashu/redisDao"
)

func GetHandler(resp http.ResponseWriter, req *http.Request) {
	redisClient := redisDao.RedisConnection()
	if redisClient == nil {
		log.Fatal("Error initialising redis")
	}
	longUrl, ok := req.URL.Query()["long_url"]

	if !ok || len(longUrl) == 0 {
		log.Fatal("Invalid Long URL")

		resp.Write(getError())

	} else {

		// 1. find it in DB, If found write to response
		// 2. If not found create the URL.

		if checkInRedis(longUrl[0]) {
			log.Println("Found in DB")
			//to do
		} else {
			log.Println("Not Found in DB")
			slug, err := generateSlug()
			if err != nil {
				log.Fatal("Error in generating Slug")
				resp.Write(getError())
			}

			//2. store in Redis Todo

			//3. Create the Shorten Url
			shortUrl := "https://ashu/" + slug

			responseModel := model.ShortenUrl{
				LongUrl:  longUrl[0],
				ShortUrl: shortUrl,
			}
			response, err := json.Marshal(responseModel)
			if err != nil {
				log.Fatal("Error in Marshalling")
				resp.Write(getError())
			}

			resp.Write(response)

		}

	}
}

func getError() []byte {
	err := model.Errors{
		Errors: "Error generating URL",
	}

	resp, _ := json.Marshal(err)
	return resp
}

func generateSlug() (string, error) {
	return "abcd", nil
}

func checkInRedis(s string) bool {
	return false
}
