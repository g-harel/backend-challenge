package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/g-harel/shopify-challenge/internal/fetcher"
	"github.com/g-harel/shopify-challenge/internal/menu"
	"github.com/g-harel/shopify-challenge/internal/validator"
)

func main() {
	var url = "https://backend-challenge-summer-2018.herokuapp.com/challenges.json?id=2"
	var maxDepth = 4

	var menus []*menu.Menu

	menus, err := fetcher.Fetch(url)
	if err != nil {
		log.Fatal(err)
	}

	res := validator.Validate(menus, maxDepth)

	j, err := json.MarshalIndent(res, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(j))
}
