package menu

// Menu holds information about a menu item.
type Menu struct {
	ID       int    `json:"id"`
	Data     string `json:"data"`
	ParentID int    `json:"parent_id"`
	ChildIDs []int  `json:"child_ids"`
}

type CheckedMenu struct {
	RootID   int   `json:"root_id"`
	Children []int `json:"children"`
}

// HasParent indicates whether or not the Menu has a parent.
func (m *Menu) HasParent() bool {
	// A value of zero represents an unmarshalled null value. (from json.Unmarshall)
	// Currently, the API only returns menus with an index larger than zero.
	return m.ParentID != 0
}
