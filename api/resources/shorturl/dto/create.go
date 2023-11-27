package dto

type CreateRequest struct {
	OriginalUrl string `json:"originalUrl"`
}

type CreateResponse struct {
	RedirectionUrl string `json:"redirectionUrl"`
}
