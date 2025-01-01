package entities

// ProgramCard ...
type ProgramCard struct {
	ID            int        `bson:"_id"`
	Name          string     `bson:"name"`
	ProgramTypeID int        `bson:"programTypeID"`
	DifficultyID  int        `bson:"difficultyID"`
	Trainings     []Training `bson:"trainings"`
}
