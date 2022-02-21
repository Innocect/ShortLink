package model

type ShortenUrl struct {
	ShortUrl string `json:"shorturl"`
	LongUrl  string `json:"longurl"`
}

type Errors struct {
	Errors string `json:"error"`
}
