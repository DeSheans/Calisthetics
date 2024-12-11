package entities

// Difficulty ...
type Difficulty struct {
	ID   int    `bson:"_id"`
	Name string `bson:"name"`
}
