package main

import (
	"sync"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"
)

type Menu struct {
	ID       int    `json:"id"`
	Data     string `json:"data"`
	ParentID int   `json:"parent_id"`
	ChildIDs []int  `json:"child_ids"`
}

type Pagination struct {
	CurrentPage int `json:"current_page"`
	PerPage     int `json:"per_page"`
	Total       int `json:"total"`
}

type APIResponse struct {
	Menus      []Menu     `json:"menus"`
	Pagination Pagination `json:"pagination"`
}

func getPage(url string, page int) (*APIResponse, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	query := req.URL.Query()
	query.Add("page", strconv.Itoa(page))
	req.URL.RawQuery = query.Encode()

	fmt.Println(req.URL.String())

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	r := &APIResponse{}
	err = json.Unmarshal(body, r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func getMenus(url string) ([]Menu, error) {
	menus := []Menu{}

	page, err := getPage(url, 1)
	if err != nil {
		return nil, err
	}
	menus = append(menus, page.Menus...)

	totalMenus := float64(page.Pagination.Total)
	perPage := float64(page.Pagination.PerPage)
	totalPages := int(math.Ceil(totalMenus / perPage))

	wg := sync.WaitGroup{}
	mux := sync.Mutex{}

	for i := 1; i < totalPages; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			page, err := getPage(url, index + 1)
			if err != nil {
				fmt.Println(err)
			}
			mux.Lock()
			menus = append(menus, page.Menus...)
			mux.Unlock()
		}(i)
	}

	wg.Wait()

	return menus, nil
}

func main() {
	menus, err := getMenus("https://backend-challenge-summer-2018.herokuapp.com/challenges.json?id=2")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", menus)
}
