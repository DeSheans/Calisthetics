package exporttypes

// FilterGroup ...
type FilterGroup struct {
	GroupID   string         `json:"groupID"`
	Name      string         `json:"name"`
	Options   []FilterOption `json:"options"`
	GroupType string         `json:"groupType"`
}

// FilterOption ...
type FilterOption struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
