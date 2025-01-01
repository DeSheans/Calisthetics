package database

import "gin-web-server/database/entities"

// Database ...
type Database interface {
	GetAllExerciseCards() ([]entities.ExerciseCard, error)
	GetAllProgramCards() ([]entities.ProgramCard, error)
	GetAllMuscles() ([]entities.Muscle, error)
	GetAllEquipment() ([]entities.Equipment, error)
	GetAllExerciseTypes() ([]entities.ExerciseType, error)
	GetAllDifficulties() ([]entities.Difficulty, error)
	GetAllProgramTypes() ([]entities.ProgramType, error)

	GetDifficultyByID(int) (entities.Difficulty, error)
	GetExerciseTypeByID(int) (entities.ExerciseType, error)
	GetMuscleByID(int) (entities.Muscle, error)
	GetEquipmentByID(int) (entities.Equipment, error)

	GetProgramTypeByID(int) (entities.ProgramType, error)
	GetTrainingTypeByID(int) (entities.TrainingType, error)

	GetExerciseByID(int) (entities.Exercise, error)
	GetProgramByID(int) (entities.Program, error)
}
