package controllers

import (
	"fmt"
	"github.com/StefanPahlplatz/bookstore_items-api/domain/items"
	"github.com/StefanPahlplatz/bookstore_items-api/services"
	"github.com/StefanPahlplatz/bookstore_oauth-go/oauth"
	"net/http"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemsController struct {
}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		// TODO return error
		return
	}

	item := items.Item{
		Seller: oauth.GetCallerId(r),
	}
	result, err := services.ItemsService.Create(item)
	if err != nil {
		// TODO return error json
	}
	fmt.Println(result)
	//TODO return success status
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
