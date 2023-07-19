package service

import (
	"ApiLibrary/model/web"
	"context"
)

type BookService interface {
	GetAll(ctx context.Context) []web.BookResponse
	Create(ctx context.Context, request web.BookCreateRequest) web.BookResponse
	GetById(ctx context.Context, bookId int) web.BookResponse
	Update(ctx context.Context, request web.BookUpdateRequest) web.BookResponse
	Delete(ctx context.Context, bookId int)
}
