package database

import "gin-web-server/database/entities"

// Database ...
type Database interface {
	GetAllExerciseCards() ([]entities.ExerciseCard, error)
	GetDifficultyByID(int) (entities.Difficulty, error)
	GetExerciseTypeByID(int) (entities.ExerciseType, error)
	GetMuscleByID(int) (entities.Muscle, error)

	// GetExerciseByID(context.Context, int) (entities.Exercise, error)
	// GetDifficultyByID(context.Context, int) (entities.Difficulty, error)
	// GetEquipmentByID(context.Context, int) (entities.Equipment, error)
	// GetMuscleByID(context.Context, int) (entities.Muscle, error)
	// GetProgramByID(context.Context, int) (entities.Program, error)
	// GetProgramTypeByID(context.Context, int) (entities.ProgramType, error)
	// GetTrainingTypeByID(context.Context, int) (entities.TrainingType, error)
}
