package validator

import "github.com/g-harel/shopify-challenge/internal/menu"

type menuMap map[int]*menu.Menu
type flagMap map[int]bool

// ValidationResult defines the status of menus as either valid or invalid.
type ValidationResult struct {
	ValidMenus   []*CheckedRoot `json:"valid_menus"`
	InvalidMenus []*CheckedRoot `json:"invalid_menus"`
}

// CheckedRoot holds information about a validated menu item.
type CheckedRoot struct {
	RootID   int   `json:"root_id"`
	Children []int `json:"children"`
}

// Recursively travels down the tree of menus to check that there are no
// cycles and that no menu is located deeper than the max depth. It also
// keeps track of travelled menu IDs in the members map which is available
// by reference to the caller.
func check(menus menuMap, members flagMap, maxDepth, id, depth int) bool {
	if depth > maxDepth {
		return false
	}
	if members[id] {
		return false
	}

	members[id] = true
	depth++

	for _, childID := range menus[id].ChildIDs {
		if !check(menus, members, maxDepth, childID, depth) {
			return false
		}
	}

	return true
}

// Validate checks a list of menus for cycles and depth.
func Validate(menus []*menu.Menu, maxDepth int) ValidationResult {
	result := ValidationResult{}

	// Creating a map from the menus for direct access using IDs.
	mm := make(menuMap)
	for _, m := range menus {
		mm[m.ID] = m
	}

	for _, m := range menus {
		// Only checking root nodes' validity.
		if m.HasParent() {
			continue
		}

		members := make(flagMap)
		valid := check(mm, members, maxDepth, m.ID, 1)

		var membersList []int
		for id := range members {
			// The root's id should not be included in the list
			// of its own children even if it was traversed.
			if id == m.ID {
				continue
			}
			membersList = append(membersList, id)
		}

		cr := &CheckedRoot{
			RootID:   m.ID,
			Children: membersList,
		}
		if valid {
			result.ValidMenus = append(result.ValidMenus, cr)
		} else {
			result.InvalidMenus = append(result.InvalidMenus, cr)
		}
	}

	return result
}
