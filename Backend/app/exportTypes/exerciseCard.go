package exporttypes

// ExerciseCard ...
type ExerciseCard struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Muscles      []string `json:"muscles"`
	Difficulty   string   `json:"difficulty"`
	ExerciseType string   `json:"exerciseType"`
	Picture      string   `json:"picture"`
}
