package entities

// Muscle ...
type Muscle struct {
	ID   int    `bson:"_id"`
	Name string `bson:"name"`
}
