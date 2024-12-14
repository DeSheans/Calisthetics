package handlers

import exporttypes "gin-web-server/app/exportTypes"

// Handler ...
type Handler interface {
	GetAllExerciseCards() ([]exporttypes.ExerciseCard, error)
}
