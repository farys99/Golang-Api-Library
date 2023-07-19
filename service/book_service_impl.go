package service

import (
	"ApiLibrary/model/web"
	"context"
)

type BookServiceImpl struct{}

func (service *BookServiceImpl) GetAll(ctx context.Context) []web.BookResponse {
	//TODO implement me
	panic("implement me")
}

func (service *BookServiceImpl) Create(ctx context.Context, request web.BookCreateRequest) web.BookResponse {
	//TODO implement me
	panic("implement me")
}

func (service *BookServiceImpl) GetById(ctx context.Context, bookId int) web.BookResponse {
	//TODO implement me
	panic("implement me")
}

func (service *BookServiceImpl) Update(ctx context.Context, request web.BookUpdateRequest) web.BookResponse {
	//TODO implement me
	panic("implement me")
}

func (service *BookServiceImpl) Delete(ctx context.Context, bookId int) {
	//TODO implement me
	panic("implement me")
}
