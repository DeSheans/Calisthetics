package exporttypes

// Program ...
type Program struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	ProgramType string     `json:"programType"`
	Difficulty  string     `json:"difficulty"`
	Description string     `json:"description"`
	Trainings   []Training `json:"trainings"`
}

// Training ...
type Training struct {
	Name         string             `json:"name"`
	TrainingType string             `json:"trainingType"`
	Exercises    []TrainingExercise `json:"exercises"`
}

// TrainingExercise ...
type TrainingExercise struct {
	Exercise string `json:"exercise"`
	Reps     string `json:"reps"`
	Sets     string `json:"sets"`
	Rest     string `json:"rest"`
}
