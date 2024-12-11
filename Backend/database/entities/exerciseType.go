package entities

// ExerciseType ...
type ExerciseType struct {
	ID   int    `bson:"_id"`
	Name string `bson:"name"`
}
