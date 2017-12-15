package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/g-harel/shopify-challenge/internal/fetcher"
	"github.com/g-harel/shopify-challenge/internal/menu"
	"github.com/g-harel/shopify-challenge/internal/validator"
)

var url = "https://backend-challenge-summer-2018.herokuapp.com/challenges.json?id=2"

var dry = true

var dummyMenus1 = []*menu.Menu{
	&menu.Menu{
		ID:       1,
		Data:     "House",
		ParentID: 0,
		ChildIDs: []int{3},
	},
	&menu.Menu{
		ID:       2,
		Data:     "Company",
		ParentID: 0,
		ChildIDs: []int{4, 5, 8},
	},
	&menu.Menu{
		ID:       3,
		Data:     "Living Room",
		ParentID: 1,
		ChildIDs: []int{7},
	},
	&menu.Menu{
		ID:       4,
		Data:     "Meeting Rooms",
		ParentID: 2,
		ChildIDs: []int{},
	},
	&menu.Menu{
		ID:       5,
		Data:     "Kitchen",
		ParentID: 2,
		ChildIDs: []int{6},
	},
	&menu.Menu{
		ID:       6,
		Data:     "Oven",
		ParentID: 5,
		ChildIDs: []int{},
	},
	&menu.Menu{
		ID:       7,
		Data:     "Table",
		ParentID: 3,
		ChildIDs: []int{15},
	},
	&menu.Menu{
		ID:       8,
		Data:     "HR",
		ParentID: 2,
		ChildIDs: []int{},
	},
	&menu.Menu{
		ID:       9,
		Data:     "Computer",
		ParentID: 0,
		ChildIDs: []int{10, 11, 12},
	},
	&menu.Menu{
		ID:       10,
		Data:     "CPU",
		ParentID: 9,
		ChildIDs: []int{},
	},
	&menu.Menu{
		ID:       11,
		Data:     "Motherboard",
		ParentID: 9,
		ChildIDs: []int{},
	},
	&menu.Menu{
		ID:       12,
		Data:     "Peripherals",
		ParentID: 9,
		ChildIDs: []int{13, 14},
	},
	&menu.Menu{
		ID:       13,
		Data:     "Mouse",
		ParentID: 12,
		ChildIDs: []int{},
	},
	&menu.Menu{
		ID:       14,
		Data:     "Keyboard",
		ParentID: 12,
		ChildIDs: []int{},
	},
	&menu.Menu{
		ID:       15,
		Data:     "Chair",
		ParentID: 7,
		ChildIDs: []int{1},
	},
}

func main() {
	var menus []*menu.Menu

	if !dry {
		m, err := fetcher.Fetch(url)
		if err != nil {
			log.Fatal(err)
		}
		menus = m
	} else {
		menus = dummyMenus1
	}

	res := validator.Validate(menus, 4)
	j, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(j))
}
