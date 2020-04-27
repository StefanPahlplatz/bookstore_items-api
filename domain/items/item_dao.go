package items

import (
	"errors"
	"github.com/StefanPahlplatz/bookstore_items-api/clients/elasticsearch"
	"github.com/StefanPahlplatz/bookstore_utils-go/rest_errors"
)

const (
	indexItems = "items"
)

func (item *Item) Save() *rest_errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, item)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save item", errors.New("database error"))
	}
	item.Id = result.Id
	return nil
}
