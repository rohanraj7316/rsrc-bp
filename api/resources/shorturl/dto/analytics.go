package dto

type AnalyticsResponse struct {
	TopThreeShortedDomains map[string]int `json:"topThreeShortedDomains"`
}
