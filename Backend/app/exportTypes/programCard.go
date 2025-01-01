package exporttypes

// ProgramCard ...
type ProgramCard struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	ProgramType string     `json:"programType"`
	Difficulty  string     `json:"difficulty"`
	Trainings   []Training `json:"trainings"`
}
