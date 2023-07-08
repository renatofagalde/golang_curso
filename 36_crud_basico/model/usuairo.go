package model

type Usuario struct {
	id    int    `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}
