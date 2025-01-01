package handlers

import (
	exporttypes "gin-web-server/app/exportTypes"
	"gin-web-server/database"
	"gin-web-server/database/entities"
)

// Filter ...
type Filter struct {
	FilterID string
	Name     string
	Type     string
	Fetch    func(database.Database) ([]exporttypes.FilterOption, error)
	Filter   func(any, []int) bool
}

var exerciseFilters = []Filter{
	{
		FilterID: "muscle",
		Name:     "группы мышц",
		Type:     "checkbox",
		Fetch:    fetchMuscles,
		Filter:   filterMuscles,
	},
	{
		FilterID: "difficulty",
		Name:     "сложность",
		Type:     "radio",
		Fetch:    fetchDifficulty,
		Filter:   filterDifficulty,
	},
	{
		FilterID: "exerciseType",
		Name:     "тип упражнения",
		Type:     "radio",
		Fetch:    fetchExerciseType,
		Filter:   filterExerciseType,
	},
	{
		FilterID: "equipment",
		Name:     "оборудование",
		Type:     "checkbox",
		Fetch:    fetchEquipment,
		Filter:   filterEquipment,
	},
}

var programFilters = []Filter{
	{
		FilterID: "programType",
		Name:     "тип программы",
		Type:     "radio",
		Fetch:    fetchProgramType,
		Filter:   filterProgramType,
	},
	{
		FilterID: "difficulty",
		Name:     "сложность",
		Type:     "radio",
		Fetch:    fetchDifficulty,
		Filter:   filterProgramDifficulty,
	},
}

func fetchMuscles(db database.Database) ([]exporttypes.FilterOption, error) {
	muscles, err := db.GetAllMuscles()
	if err != nil {
		return nil, err
	}
	filterOptions := make([]exporttypes.FilterOption, len(muscles))
	for i, v := range muscles {
		filterOptions[i] = exporttypes.FilterOption(v)
	}
	return filterOptions, nil
}

func filterMuscles(e any, selectedIDs []int) bool {
	exercise := e.(entities.ExerciseCard)
	exerciseMuscles := make(map[int]bool, len(exercise.MusclesID))
	for _, v := range exercise.MusclesID {
		exerciseMuscles[v] = true
	}

	for _, v := range selectedIDs {
		if _, ok := exerciseMuscles[v]; ok == false {
			return false
		}
	}
	return true
}

func fetchDifficulty(db database.Database) ([]exporttypes.FilterOption, error) {
	difficulty, err := db.GetAllDifficulties()
	if err != nil {
		return nil, err
	}
	filterOption := make([]exporttypes.FilterOption, len(difficulty))
	for i, v := range difficulty {
		filterOption[i] = exporttypes.FilterOption(v)
	}
	return filterOption, nil
}

func filterDifficulty(e any, selectedIDs []int) bool {
	exercise := e.(entities.ExerciseCard)
	for _, v := range selectedIDs {
		if exercise.DifficultyID != v {
			return false
		}
	}
	return true
}

func fetchExerciseType(db database.Database) ([]exporttypes.FilterOption, error) {
	exerciseTypes, err := db.GetAllExerciseTypes()
	if err != nil {
		return nil, err
	}
	filterOptions := make([]exporttypes.FilterOption, len(exerciseTypes))
	for i, v := range exerciseTypes {
		filterOptions[i] = exporttypes.FilterOption(v)
	}
	return filterOptions, nil
}

func filterExerciseType(e any, selectedIDs []int) bool {
	exercise := e.(entities.ExerciseCard)
	for _, v := range selectedIDs {
		if exercise.ExerciseTypeID != v {
			return false
		}
	}
	return true
}

func fetchEquipment(db database.Database) ([]exporttypes.FilterOption, error) {
	equip, err := db.GetAllEquipment()
	if err != nil {
		return nil, err
	}
	filterOptions := make([]exporttypes.FilterOption, len(equip))
	for i, v := range equip {
		filterOptions[i] = exporttypes.FilterOption(v)
	}
	return filterOptions, nil
}

func filterEquipment(e any, selectedIDs []int) bool {
	exercise := e.(entities.ExerciseCard)
	if len(exercise.EquipmentID) == 0 {
		return true
	}
	selectedEquipment := make(map[int]bool, len(selectedIDs))
	for _, v := range selectedIDs {
		selectedEquipment[v] = true
	}
	for _, equip := range exercise.EquipmentID {
		if _, ok := selectedEquipment[equip]; ok == false {
			return false
		}
	}
	return true
}

func fetchProgramType(db database.Database) ([]exporttypes.FilterOption, error) {
	prType, err := db.GetAllProgramTypes()
	if err != nil {
		return nil, err
	}
	filterOptions := make([]exporttypes.FilterOption, len(prType))
	for i, v := range prType {
		filterOptions[i] = exporttypes.FilterOption(v)
	}
	return filterOptions, nil
}

func filterProgramType(e any, selectedIDs []int) bool {
	program := e.(entities.ProgramCard)
	for _, v := range selectedIDs {
		if program.ProgramTypeID != v {
			return false
		}
	}
	return true
}

func filterProgramDifficulty(e any, selectedIDs []int) bool {
	program := e.(entities.ProgramCard)
	for _, v := range selectedIDs {
		if program.DifficultyID != v {
			return false
		}
	}
	return true
}
