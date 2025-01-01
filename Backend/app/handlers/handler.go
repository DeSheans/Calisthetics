package handlers

import (
	exporttypes "gin-web-server/app/exportTypes"
	"net/url"
)

// Handler ...
type Handler interface {
	GetAllExerciseCards(url.Values) ([]exporttypes.ExerciseCard, error)
	GetAllProgramCards(url.Values) ([]exporttypes.ProgramCard, error)
	GetProgramByID(int) (exporttypes.Program, error)
	GetExerciseByID(int) (exporttypes.Exercise, error)
	GetExerciseFilters() ([]exporttypes.FilterGroup, error)
	GetProgramFilters() ([]exporttypes.FilterGroup, error)
}
