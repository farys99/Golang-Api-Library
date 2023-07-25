package helper

import (
	"ApiLibrary/model/domain"
	"ApiLibrary/model/web"
)

func ToBookResponse(book domain.Book) web.BookResponse {
	return web.BookResponse{
		Id:        book.Id,
		Title:     book.Title,
		Years:     book.Years,
		Publisher: book.Publisher,
	}
}

func ToBookResponses(books []domain.Book) []web.BookResponse {
	bookResponses := make([]web.BookResponse, 0)

	for _, book := range books {
		bookResponses = append(bookResponses, ToBookResponse(book))
	}

	return bookResponses

}
