package models

type ErrResponse struct {
	Errors map[string]string `json:"errors"`
}