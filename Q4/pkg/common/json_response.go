package common

type JsonResponse[T any] struct {
	Type       string `json:"type"`
	Data       T      `json:"data"`
	Message    string `json:"message"`
	IsError    bool   `json:"iserror"`
	StatusCode int    `json:"statuscode"`
}
