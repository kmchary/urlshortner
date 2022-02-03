package urlshortener

type Request struct {
	Url string `json:"url"`
	UserId string `json:"user_id"`
}


