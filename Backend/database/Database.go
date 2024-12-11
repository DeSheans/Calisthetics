package database

import (
	"context"
	"gin-web-server/database/entities"
)

// Database ...
type Database interface {
	GetExerciseByID(context.Context, int) (entities.Exercise, error)
	GetDifficultyByID(context.Context, int) (entities.Difficulty, error)
	GetEquipmentByID(context.Context, int) (entities.Equipment, error)
	GetExerciseTypeByID(context.Context, int) (entities.ExerciseType, error)
	GetMuscleByID(context.Context, int) (entities.Muscle, error)
	GetProgramByID(context.Context, int) (entities.Program, error)
	GetProgramTypeByID(context.Context, int) (entities.ProgramType, error)
	GetTrainingTypeByID(context.Context, int) (entities.TrainingType, error)
}
