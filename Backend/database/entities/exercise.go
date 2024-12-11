package entities

// Exercise ...
type Exercise struct {
	ID             int      `bson:"_id"`
	Name           string   `bson:"name"`
	MusclesID      []int    `bson:"musclesID"`
	DifficultyID   int      `bson:"difficultyID"`
	ExerciseTypeID int      `bson:"exerciseTypeID"`
	EquipmentID    []int    `bson:"equipmentID"`
	Description    string   `bson:"description"`
	Tips           []string `bson:"tips"`
	Variations     []string `bson:"variations"`
	Pictures       []string `bson:"pictures"`
}
