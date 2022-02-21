package handler

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/ashu/model"
	"github.com/go-redis/redis"
)

func GetHandler(redisClient *redis.Client) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {

		longUrl, ok := req.URL.Query()["long_url"]

		if !ok || len(longUrl) == 0 {
			log.Fatal("Invalid Long URL")
			resp.Write(getError("Invalid Long URL"))

		} else {

			redisResult, _ := redisClient.Get(longUrl[0]).Bytes()

			var redisData *model.ShortenUrl
			_ = json.Unmarshal(redisResult, &redisData)

			if len(redisResult) > 0 && redisData.LongUrl == longUrl[0] {
				log.Println("Found in DB")

				response, err := json.Marshal(redisData)
				if err != nil {
					log.Fatal("Error in Marshalling response")
					resp.Write(getError("Error in Marshalling response"))
				}
				resp.Write(response)

			} else {
				log.Println("Not Found in DB")
				slug, err := generateSlug()
				if err != nil {
					log.Fatal("Error in generating Slug")
					resp.Write(getError("Error in generating Slug"))
				}

				shortUrl := "https://ashu/" + slug

				responseModel := model.ShortenUrl{
					LongUrl:      longUrl[0],
					ShortUrl:     shortUrl,
					AlreadyExist: false,
				}
				response, err := json.Marshal(responseModel)
				if err != nil {
					log.Fatal("Error in Marshalling response")
					resp.Write(getError("Error in Marshalling response"))
				}

				err = redisClient.Set(longUrl[0], response, 0).Err()
				if err != nil {
					log.Fatal(err)
					resp.Write(getError(err.Error()))
				}

				log.Println("Request served Successfully")
				resp.Write(response)

			}

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

func getLetters() []rune {
	var letters = []rune("abcdefghjkmnpqrtuvwxyzACDEFGHJKMNPQRTUVWXYZ")
	rand.Shuffle(len(letters), func(i, j int) {
		letters[i], letters[j] = letters[j], letters[i]
	})
	return letters
}
