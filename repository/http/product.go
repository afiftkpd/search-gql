package http

import (
	"context"
	"encoding/json"
	"fmt"
	"gql/models"
	"io/ioutil"
	"net/http"
	"strconv"
)

type productRepository struct {
	// DB *sql.DB
}

func NewProductRepository() ProductRepository {
	return &productRepository{}
}

func (p *productRepository) Search(ctx context.Context, keyword string, page int, limit int) ([]models.Product, error) {
	fmt.Println("send request")
	fmt.Println(fmt.Sprintf("http://localhost:8081/search?keyword=%v&page=%v&limit=%v", keyword, page, limit))
	products := []models.Product{}
	res, err := http.Get("http://localhost:8081/search?keyword=" + keyword + "&page=" + strconv.Itoa(page) + "&limit=" + strconv.Itoa(limit))
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		return products, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body) // response body is []byte

	fmt.Println("here")
	fmt.Println(string(body))

	err = json.Unmarshal(body, &products)
	if err != nil {
		fmt.Println(string(body))
		return products, err
	}
	return products, nil
}

func (p *productRepository) AutoComplete(ctx context.Context, keyword string) ([]models.AutoComplete, error) {
	fmt.Println("send request")
	products := []models.AutoComplete{}
	res, err := http.Get("http://localhost:8081/autocomplete?keyword=" + keyword)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		return products, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body) // response body is []byte

	err = json.Unmarshal(body, &products)
	if err != nil {
		fmt.Println(string(body))
		return products, err
	}

	return products, nil
}
