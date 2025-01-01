package model

type Usuario struct {
	ID    int    `json:"ID"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}
