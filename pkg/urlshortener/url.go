package urlshortener

type Request struct {
	Url    string `json:"url"`
	UserId string `json:"user_id"`
}

type Response struct {
	ShortUrl string `json:"short_url,omitempty" `
	Error    string `json:"error,omitempty"`
}
