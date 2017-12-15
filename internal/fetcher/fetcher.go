package fetcher

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/g-harel/shopify-challenge/internal/menu"
)

type pagination struct {
	CurrentPage int `json:"current_page"`
	PerPage     int `json:"per_page"`
	Total       int `json:"total"`
}

type apiResponse struct {
	Menus      []*menu.Menu `json:"menus"`
	Pagination pagination   `json:"pagination"`
}

// Queries the provided url for a single page of menus and metadata.
func fetchPage(url string, page int) (*apiResponse, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Desired page is added as a query param to the request.
	query := req.URL.Query()
	query.Add("page", strconv.Itoa(page))
	req.URL.RawQuery = query.Encode()

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	r := &apiResponse{}
	err = json.Unmarshal(body, r)
	if err != nil {
		return nil, fmt.Errorf("error parsing response: %s", err)
	}

	return r, nil
}

// Fetch queries the provided url for all available menus.
func Fetch(url string) ([]*menu.Menu, error) {
	menus := []*menu.Menu{}

	// First page is manually fetched to use pagination information.
	page, err := fetchPage(url, 1)
	if err != nil {
		return nil, err
	}
	menus = append(menus, page.Menus...)

	totalMenus := float64(page.Pagination.Total)
	perPage := float64(page.Pagination.PerPage)
	if perPage == 0 {
		return nil, fmt.Errorf("division by zero")
	}
	totalPages := int(math.Ceil(totalMenus / perPage))

	// Remaining page fetches are done in parallel.
	ch := make(chan *apiResponse)
	errs := make(chan error)
	for i := 1; i < totalPages; i++ {
		go func(index int) {
			page, err := fetchPage(url, index+1)
			ch <- page
			errs <- err
		}(i)
	}
	for i := 1; i < totalPages; i++ {
		page := <-ch
		err := <-errs
		if err != nil {
			return nil, err
		}
		menus = append(menus, page.Menus...)
	}

	return menus, nil
}
