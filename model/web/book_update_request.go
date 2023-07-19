package web

type BookUpdateRequest struct {
	Id        int    `validate:"required"`
	Title     string `validate:"required,min=1,max=200" json:"title"`
	Years     string `validate:"required,min=1,max=4" json:"years"`
	Publisher string `validate:"required,min=1,max=200" json:"publisher"`
}
