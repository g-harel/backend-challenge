package validator

import "github.com/g-harel/shopify-challenge/internal/menu"

type menuMap map[int]*menu.Menu
type flagMap map[int]bool

// ValidationResult defines the status of menus as either valid or invalid.
type ValidationResult struct {
	ValidMenus   []*menu.CheckedMenu `json:"valid_menus"`
	InvalidMenus []*menu.CheckedMenu `json:"invalid_menus"`
}

// Recursively travels down the tree of menus to check that there are no
// cycles and that no menu is located deeper than the max depth. It also
// keeps track of travelled menu IDs in the members map which is available
// by reference to the caller.
func check(mm menuMap, members flagMap, maxDepth, id, depth int) bool {
	if depth > maxDepth {
		return false
	}
	if members[id] {
		return false
	}

	members[id] = true
	depth++

	for _, childID := range mm[id].ChildIDs {
		if !check(mm, members, maxDepth, childID, depth) {
			return false
		}
	}

	return true
}

// Validate checks a slice of menus for cycles.
func Validate(menus []*menu.Menu, maxDepth int) ValidationResult {
	// Creating a map from the menus for direct access using IDs.
	mm := make(menuMap)
	for _, m := range menus {
		mm[m.ID] = m
	}

	// Marking all valid menus IDs.
	valid := make(flagMap)
	for _, m := range menus {
		// Only checking the validity from "root" menus with no parents.
		// It is assumed that menus with no root menu in their ancestry
		// are part of a cycle.
		if m.HasParent() {
			continue
		}

		members := make(flagMap)
		if !check(mm, members, maxDepth, m.ID, 1) {
			continue
		}

		// All the members of a valid menu tree are also considered valid.
		// For this assumption to hold, the menus cannot have multiple parents.
		for id := range members {
			valid[id] = true
		}
	}

	result := ValidationResult{}
	for _, m := range menus {
		cm := &menu.CheckedMenu{
			RootID:   m.ID,
			Children: m.ChildIDs,
		}

		if valid[m.ID] {
			result.ValidMenus = append(result.ValidMenus, cm)
		} else {
			result.InvalidMenus = append(result.InvalidMenus, cm)
		}
	}

	return result
}
