package handlers

import (
	exporttypes "gin-web-server/app/exportTypes"
	"gin-web-server/database"
	"gin-web-server/database/entities"
	"net/url"
	"slices"
	"strconv"
	"strings"
)

type concreteHandler struct {
	db database.Database
}

// NewConcreteHandler ...
func NewConcreteHandler(db database.Database) Handler {
	return concreteHandler{db: db}
}

// GetAllExerciseCards ...
func (h concreteHandler) GetAllExerciseCards(url url.Values) ([]exporttypes.ExerciseCard, error) {
	exercises, err := h.db.GetAllExerciseCards()
	if err != nil {
		return nil, err
	}

	urlValues := make(map[string][]int)
	for param, values := range url {
		vals := []int{}
		for _, temp := range values {
			for _, val := range strings.Split(temp, ",") {
				integer, err := strconv.Atoi(val)
				if err == nil {
					vals = append(vals, integer)
				}
			}
		}
		urlValues[param] = vals
	}

	var filteredExercises []entities.ExerciseCard
	for _, exercise := range exercises {
		flag := true
		for filterID, values := range urlValues {
			filterIndex := slices.IndexFunc(exerciseFilters, func(e Filter) bool { return e.FilterID == filterID })

			if filterIndex == -1 || exerciseFilters[filterIndex].Filter(exercise, values) == false {
				flag = false
				break
			}
		}
		if flag {
			filteredExercises = append(filteredExercises, exercise)
		}
	}

	result := make([]exporttypes.ExerciseCard, len(filteredExercises))
	musclesMap := make(map[int]string)
	difficultyMap := make(map[int]string)
	exerciseTypeMap := make(map[int]string)

	for i, v := range filteredExercises {
		var difficulty string
		var muscles = make([]string, len(v.MusclesID))
		var exerciseType string
		var picture string
		if len(v.Pictures) > 0 {
			picture = v.Pictures[0]
		} else {
			picture = ""
		}

		if value, ok := difficultyMap[v.DifficultyID]; ok == false {
			dbRes, err := h.db.GetDifficultyByID(v.DifficultyID)
			if err != nil {
				return nil, err
			}
			difficulty = dbRes.Name
			difficultyMap[v.DifficultyID] = dbRes.Name
		} else {
			difficulty = value
		}

		if value, ok := exerciseTypeMap[v.ExerciseTypeID]; ok == false {
			dbRes, err := h.db.GetExerciseTypeByID(v.ExerciseTypeID)
			if err != nil {
				return nil, err
			}
			exerciseType = dbRes.Name
			exerciseTypeMap[v.ExerciseTypeID] = dbRes.Name
		} else {
			exerciseType = value
		}

		for i, v := range v.MusclesID {
			if m, ok := musclesMap[v]; ok == false {
				temp, err := h.db.GetMuscleByID(v)
				if err != nil {
					return nil, err
				}
				muscles[i] = temp.Name
			} else {
				muscles[i] = m
			}
		}

		result[i] = exporttypes.ExerciseCard{
			ID:           v.ID,
			Name:         v.Name,
			Muscles:      muscles,
			Difficulty:   difficulty,
			ExerciseType: exerciseType,
			Picture:      picture,
		}
	}

	return result, nil
}

// GetAllProgramCards ...
func (h concreteHandler) GetAllProgramCards(url url.Values) ([]exporttypes.ProgramCard, error) {
	programs, err := h.db.GetAllProgramCards()
	if err != nil {
		return nil, err
	}

	urlValues := make(map[string][]int)
	for param, values := range url {
		vals := []int{}
		for _, temp := range values {
			for _, val := range strings.Split(temp, ",") {
				integer, err := strconv.Atoi(val)
				if err == nil {
					vals = append(vals, integer)
				}
			}
		}
		urlValues[param] = vals
	}

	var filteredPrograms []entities.ProgramCard
	for _, program := range programs {
		flag := true
		for filterID, values := range urlValues {
			filterIndex := slices.IndexFunc(programFilters, func(e Filter) bool { return e.FilterID == filterID })

			if filterIndex == -1 || programFilters[filterIndex].Filter(program, values) == false {
				flag = false
				break
			}
		}
		if flag {
			filteredPrograms = append(filteredPrograms, program)
		}
	}

	result := make([]exporttypes.ProgramCard, len(filteredPrograms))

	programTypeMap := make(map[int]string)
	difficultyMap := make(map[int]string)
	trainingTypeMap := make(map[int]string)
	exerciseNameMap := make(map[int]string)

	for index, card := range filteredPrograms {
		program, err := convertProgramToExportType(
			card,
			programTypeMap,
			difficultyMap,
			trainingTypeMap,
			exerciseNameMap,
			h.db,
		)
		if err != nil {
			return nil, err
		}

		result[index] = program
	}

	return result, nil
}

// GetProgramByID ...
func (h concreteHandler) GetProgramByID(id int) (exporttypes.Program, error) {
	program, err := h.db.GetProgramByID(id)
	if err != nil {
		return exporttypes.Program{}, err
	}

	trainings := make([]exporttypes.Training, len(program.Trainings))

	programType, err := h.db.GetProgramTypeByID(program.ProgramTypeID)
	if err != nil {
		return exporttypes.Program{}, err
	}
	difficulty, err := h.db.GetDifficultyByID(program.DifficultyID)
	if err != nil {
		return exporttypes.Program{}, err
	}

	for index, value := range program.Trainings {
		temp, err := convertTrainingToExportType(value, map[int]string{}, map[int]string{}, h.db)
		if err != nil {
			return exporttypes.Program{}, err
		}
		trainings[index] = temp
	}

	return exporttypes.Program{
		ID:          id,
		Name:        program.Name,
		ProgramType: programType.Name,
		Difficulty:  difficulty.Name,
		Description: program.Description,
		Trainings:   trainings,
	}, nil
}

// GetExerciseByID ...
func (h concreteHandler) GetExerciseByID(id int) (exporttypes.Exercise, error) {
	exercise, err := h.db.GetExerciseByID(id)
	if err != nil {
		return exporttypes.Exercise{}, err
	}

	muscles := make([]string, len(exercise.MusclesID))
	for index, value := range exercise.MusclesID {
		muscle, err := h.db.GetMuscleByID(value)
		if err != nil {
			return exporttypes.Exercise{}, err
		}
		muscles[index] = muscle.Name
	}

	equipments := make([]string, len(exercise.EquipmentID))
	for index, value := range exercise.EquipmentID {
		equipment, err := h.db.GetEquipmentByID(value)
		if err != nil {
			return exporttypes.Exercise{}, err
		}
		equipments[index] = equipment.Name
	}

	difficulty, err := h.db.GetDifficultyByID(exercise.DifficultyID)
	if err != nil {
		return exporttypes.Exercise{}, err
	}

	exType, err := h.db.GetExerciseTypeByID(exercise.ExerciseTypeID)
	if err != nil {
		return exporttypes.Exercise{}, err
	}

	return exporttypes.Exercise{
		ID:           id,
		Name:         exercise.Name,
		Muscles:      muscles,
		Difficulty:   difficulty.Name,
		ExerciseType: exType.Name,
		Equipment:    equipments,
		Description:  exercise.Description,
		Tips:         exercise.Tips,
		Variations:   exercise.Variations,
		Pictures:     exercise.Pictures,
	}, nil
}

func convertProgramToExportType(
	program entities.ProgramCard,
	programTypeMap map[int]string,
	difficultyMap map[int]string,
	trainingTypeMap map[int]string,
	exerciseNameMap map[int]string,
	db database.Database,
) (exporttypes.ProgramCard, error) {

	var programType string
	var difficulty string
	trainings := make([]exporttypes.Training, len(program.Trainings))

	if value, ok := programTypeMap[program.ProgramTypeID]; ok == false {
		dbRes, err := db.GetProgramTypeByID(program.ProgramTypeID)
		if err != nil {
			return exporttypes.ProgramCard{}, err
		}
		programType = dbRes.Name
		programTypeMap[program.ProgramTypeID] = dbRes.Name
	} else {
		difficulty = value
	}

	if value, ok := difficultyMap[program.DifficultyID]; ok == false {
		dbRes, err := db.GetDifficultyByID(program.DifficultyID)
		if err != nil {
			return exporttypes.ProgramCard{}, err
		}
		difficulty = dbRes.Name
		difficultyMap[program.DifficultyID] = dbRes.Name
	} else {
		difficulty = value
	}

	for trainingIndex, training := range program.Trainings {
		res, err := convertTrainingToExportType(training, trainingTypeMap, exerciseNameMap, db)
		if err != nil {
			return exporttypes.ProgramCard{}, err
		}

		trainings[trainingIndex] = res
	}

	return exporttypes.ProgramCard{
		ID:          program.ID,
		Name:        program.Name,
		ProgramType: programType,
		Difficulty:  difficulty,
		Trainings:   trainings,
	}, nil
}

func convertTrainingToExportType(
	training entities.Training,
	trainingTypeMap map[int]string,
	exerciseNameMap map[int]string,
	db database.Database,
) (exporttypes.Training, error) {

	var trainingType string
	exercises := make([]exporttypes.TrainingExercise, len(training.Exercises))

	if value, ok := trainingTypeMap[training.TrainingTypeID]; ok == false {
		dbRes, err := db.GetTrainingTypeByID(training.TrainingTypeID)
		if err != nil {
			return exporttypes.Training{}, err
		}
		trainingType = dbRes.Name
		trainingTypeMap[training.TrainingTypeID] = dbRes.Name
	} else {
		trainingType = value
	}

	for index, value := range training.Exercises {
		res, err := convertExerciseToExportType(value, exerciseNameMap, db)
		if err != nil {
			return exporttypes.Training{}, err
		}
		exercises[index] = res
	}

	return exporttypes.Training{
		Name:         training.Name,
		TrainingType: trainingType,
		Exercises:    exercises,
	}, nil
}

func convertExerciseToExportType(
	exercise entities.TrainingExercise,
	m map[int]string,
	db database.Database,
) (exporttypes.TrainingExercise, error) {
	var exerciseName string
	if value, ok := m[exercise.ExerciseID]; ok == false {
		dbRes, err := db.GetExerciseByID(exercise.ExerciseID)
		if err != nil {
			return exporttypes.TrainingExercise{}, err
		}
		exerciseName = dbRes.Name
		m[exercise.ExerciseID] = exerciseName
	} else {
		exerciseName = value
	}

	return exporttypes.TrainingExercise{
		Exercise: exerciseName,
		Reps:     exercise.Reps,
		Sets:     exercise.Sets,
		Rest:     exercise.Rest,
	}, nil
}

// GetExerciseFilters ...
func (h concreteHandler) GetExerciseFilters() ([]exporttypes.FilterGroup, error) {
	filterGroups := make([]exporttypes.FilterGroup, len(exerciseFilters))

	for i, v := range exerciseFilters {
		opts, err := v.Fetch(h.db)
		if err != nil {
			return nil, err
		}
		filterGroups[i] = exporttypes.FilterGroup{
			GroupID:   v.FilterID,
			Name:      v.Name,
			Options:   opts,
			GroupType: v.Type,
		}
	}

	return filterGroups, nil
}

// GetProgramFilters ...
func (h concreteHandler) GetProgramFilters() ([]exporttypes.FilterGroup, error) {
	filterGroups := make([]exporttypes.FilterGroup, len(programFilters))

	for i, v := range programFilters {
		opts, err := v.Fetch(h.db)
		if err != nil {
			return nil, err
		}
		filterGroups[i] = exporttypes.FilterGroup{
			GroupID:   v.FilterID,
			Name:      v.Name,
			Options:   opts,
			GroupType: v.Type,
		}
	}

	return filterGroups, nil
}
