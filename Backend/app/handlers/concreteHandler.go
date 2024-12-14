package handlers

import (
	exporttypes "gin-web-server/app/exportTypes"
	"gin-web-server/database"
	"log"
)

// concreteHandler ...
type concreteHandler struct {
	db database.Database
}

// NewConcreteHandler ...
func NewConcreteHandler(db database.Database) Handler {
	return concreteHandler{db: db}
}

// GetAllExerciseCards ...
func (h concreteHandler) GetAllExerciseCards() ([]exporttypes.ExerciseCard, error) {
	log.Println("App handler starts GetAllExerciseCards")
	exercises, err := h.db.GetAllExerciseCards()
	if err != nil {
		log.Println("db throw error on GetAllExerciseCards. Error: ", err)

		return nil, err
	}

	result := make([]exporttypes.ExerciseCard, len(exercises))
	musclesMap := make(map[int]string)
	difficultyMap := make(map[int]string)
	exerciseTypeMap := make(map[int]string)

	for i, v := range exercises {
		var difficulty string
		var muscles = make([]string, len(v.MusclesID))
		var exerciseType string
		var picture string
		if len(v.Pictures) > 0 {
			picture = v.Pictures[0]
		} else {
			picture = ""
		}

		if d, ok := difficultyMap[v.DifficultyID]; ok == false {
			temp, err := h.db.GetDifficultyByID(v.DifficultyID)
			if err != nil {
				log.Println("db throw error on GetDifficultyByID. Error: ", err, "difficulty", v.DifficultyID)

				return nil, err
			}
			difficulty = temp.Name
		} else {
			difficulty = d
		}

		if e, ok := exerciseTypeMap[v.ExerciseTypeID]; ok == false {
			temp, err := h.db.GetExerciseTypeByID(v.ExerciseTypeID)
			if err != nil {
				log.Println("db throw error on GetExerciseTypeByID. Error: ", err, "ExerciseType: ", v.ExerciseTypeID)

				return nil, err
			}
			exerciseType = temp.Name
		} else {
			exerciseType = e
		}

		for i, v := range v.MusclesID {
			if m, ok := musclesMap[v]; ok == false {
				temp, err := h.db.GetMuscleByID(v)
				if err != nil {
					log.Println("db throw error on GetMuscleByID. Error: ", err, "muscle: ", v)

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
