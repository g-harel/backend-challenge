package menu

// Menu holds information about a menu item.
type Menu struct {
	ID       int    `json:"id"`
	Data     string `json:"data"`
	ParentID int    `json:"parent_id"`
	ChildIDs []int  `json:"child_ids"`
}

// HasParent indicates whether or not the Menu has a parent.
func (m *Menu) HasParent() bool {
	// The endpoint never returns menus with ids of 0 at this time. This means
	// that a zero value can only be the product of the json to struct conversion
	// (json.Unmarshall) where a null integer value becomes the default int (0).
	return m.ParentID != 0
}
