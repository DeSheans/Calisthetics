package entities

// Program ...
type Program struct {
	ID            int        `bson:"_id"`
	Name          string     `bson:"name"`
	ProgramTypeID int        `bson:"programTypeID"`
	DifficultyID  int        `bson:"difficultyID"`
	Description   string     `bson:"description"`
	Trainings     []Training `bson:"trainings"`
}

// Training ...
type Training struct {
	Name           string             `bson:"name"`
	TrainingTypeID int                `bson:"trainingTypeID"`
	Exercises      []TrainingExercise `bson:"exercises"`
}

// TrainingExercise ...
type TrainingExercise struct {
	ExerciseID int    `bson:"exerciseID"`
	Reps       string `bson:"reps"`
	Sets       string `bson:"sets"`
	Rest       string `bson:"rest"`
}
