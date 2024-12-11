package entities

// ExerciseCard ...
type ExerciseCard struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Muscles      []int    `json:"muscles"`
	Difficulty   int      `json:"difficulty"`
	ExerciseType int      `json:"exerciseType"`
	Picture      []string `json:"picture"`
}
