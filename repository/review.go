package repository

import "github.com/google/uuid"

type Review struct {
	Rid          uuid.UUID `json:"rid"`
	Rating       int       `json:"rating"`
	ReviewerName string    `json:"reviewerName"`
	Comment      string    `json:"comment"`
	Email        string    `json:"email"`
}
