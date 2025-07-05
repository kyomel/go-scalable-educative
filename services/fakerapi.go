package services

import (
	"net/http"

	"github.com/kyomel/go-scalable-educative/network"
)

var fakeAPIClient = network.NewClient().
	Name("fakerapi").
	Timeout(10)

func GetUsers() (res *http.Response, err error) {
	res, err = fakeAPIClient.
		Get("https://fakerapi.it/api/v1/persons")
	return
}

func GetProducts() (res *http.Response, err error) {
	res, err = fakeAPIClient.
		Get("https://fakerapi.it/api/v1/products")
	return
}

func GetBooks() (res *http.Response, err error) {
	res, err = fakeAPIClient.
		Get("https://fakerapi.it/api/v1/books")
	return
}
