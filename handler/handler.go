package handler

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

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

		resp.Write(getError("Invalid Long URL"))

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
				resp.Write(getError("Error in generating Slug"))
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
				log.Fatal("Error in Marshalling response")
				resp.Write(getError("Error in Marshalling response"))
			}

			resp.Write(response)

		}

	}
}

func getError(errName string) []byte {
	err := model.Errors{
		Errors: errName,
	}

	resp, _ := json.Marshal(err)
	return resp
}

func generateSlug() (string, error) {
	letters := getLetters()
	s := make([]rune, 4)
	x1 := rand.NewSource(time.Now().UnixNano())
	y1 := rand.New(x1)

	for i := range s {
		s[i] = letters[y1.Intn(len(letters))]
	}

	slug := string(s)

	return slug, nil
}

func checkInRedis(s string) bool {
	return false
}

func getLetters() []rune {
	var letters = []rune("23456789abcdefghjkmnpqrtuvwxyzACDEFGHJKMNPQRTUVWXYZ")
	rand.Shuffle(len(letters), func(i, j int) {
		letters[i], letters[j] = letters[j], letters[i]
	})
	return letters
}
