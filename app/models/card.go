package models

// Card data structure
type Card struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}
