package repository

type Review struct {
	Rid          int    `json:"rid"`
	Rating       int    `json:"rating"`
	ReviewerName string `json:"reviewerName"`
	Comment      string `json:"comment"`
	Email        string `json:"email"`
}
