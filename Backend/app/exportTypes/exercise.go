package exporttypes

// Exercise ...
type Exercise struct {
	ID           int      `json:"_id"`
	Name         string   `json:"name"`
	Muscles      []string `json:"muscles"`
	Difficulty   string   `json:"difficulty"`
	ExerciseType string   `json:"exerciseType"`
	Equipment    []string `json:"equipment"`
	Description  string   `json:"description"`
	Tips         []string `json:"tips"`
	Variations   []string `json:"variations"`
	Pictures     []string `json:"pictures"`
}
