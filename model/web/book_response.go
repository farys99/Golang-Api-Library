package web

type BookResponse struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Years     string `json:"years"`
	Publisher string `json:"publisher"`
}
