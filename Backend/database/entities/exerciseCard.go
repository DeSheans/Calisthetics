package entities

// ExerciseCard ...
type ExerciseCard struct {
	ID             int      `bson:"_id"`
	Name           string   `bson:"name"`
	MusclesID      []int    `bson:"musclesID"`
	DifficultyID   int      `bson:"difficultyID"`
	ExerciseTypeID int      `bson:"exerciseTypeID"`
	EquipmentID    []int    `bson:"equipmentID"`
	Pictures       []string `bson:"pictures"`
}
