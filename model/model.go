package model

type ShortenUrl struct {
	ShortUrl     string `json:"shorturl"`
	LongUrl      string `json:"longurl"`
	AlreadyExist bool   `json:"exists"`
}

type Errors struct {
	Errors string `json:"error"`
}
