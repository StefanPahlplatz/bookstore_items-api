package services

import (
	"errors"
	"github.com/StefanPahlplatz/bookstore_items-api/domain/items"
	"github.com/StefanPahlplatz/bookstore_utils-go/rest_errors"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(items.Item) (*items.Item, *rest_errors.RestErr)
	Get(string) (*items.Item, *rest_errors.RestErr)
}

type itemsService struct {
}

func (s *itemsService) Create(items.Item) (*items.Item, *rest_errors.RestErr) {
	return nil, nil
}

func (s *itemsService) Get(string) (*items.Item, *rest_errors.RestErr) {
	return nil, rest_errors.NewInternalServerError("not implemented", errors.New("todo"))
}
