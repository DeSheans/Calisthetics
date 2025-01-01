package mongo

import (
	"context"
	"fmt"
	"gin-web-server/database"
	"gin-web-server/database/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoDatabase struct {
	client *mongo.Client
	ctx    context.Context
}

const (
	dbName       = "calisthetic"
	exercise     = "exercise"
	muscle       = "muscle"
	exerciseType = "exerciseType"
	equipment    = "equipment"
	difficulty   = "difficulty"
	trainingType = "trainingType"
	programType  = "programType"
	program      = "program"
)

// NewMongoDatabase ...
func NewMongoDatabase(ctx context.Context, cfg Config) (database.Database, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(fmt.Sprintf(
		"%s:%s",
		cfg.Host,
		cfg.Port,
	)).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, err
	}
	return mongoDatabase{client: client, ctx: ctx}, nil
}

// getDocumentByID ...
func (db mongoDatabase) getDocumentByID(collection string, id int) *mongo.SingleResult {
	c := db.client.Database(dbName).Collection(collection)
	filter := bson.D{{Key: "_id", Value: id}}
	result := c.FindOne(db.ctx, filter, options.FindOne())
	return result
}

// GetDifficultyByID ...
func (db mongoDatabase) GetDifficultyByID(id int) (entities.Difficulty, error) {
	var doc entities.Difficulty
	if err := db.getDocumentByID(difficulty, id).Decode(&doc); err != nil {
		return entities.Difficulty{}, nil
	}
	return doc, nil
}

// GetExerciseTypeByID ...
func (db mongoDatabase) GetExerciseTypeByID(id int) (entities.ExerciseType, error) {
	var doc entities.ExerciseType
	if err := db.getDocumentByID(exerciseType, id).Decode(&doc); err != nil {
		return entities.ExerciseType{}, nil
	}
	return doc, nil
}

// GetMuscleByID ...
func (db mongoDatabase) GetMuscleByID(id int) (entities.Muscle, error) {
	var doc entities.Muscle
	if err := db.getDocumentByID(muscle, id).Decode(&doc); err != nil {
		return entities.Muscle{}, nil
	}
	return doc, nil
}

// GetProgramTypeByID ...
func (db mongoDatabase) GetProgramTypeByID(id int) (entities.ProgramType, error) {
	var doc entities.ProgramType
	if err := db.getDocumentByID(programType, id).Decode(&doc); err != nil {
		return entities.ProgramType{}, nil
	}
	return doc, nil
}

// GetTrainingTypeByID ...
func (db mongoDatabase) GetTrainingTypeByID(id int) (entities.TrainingType, error) {
	var doc entities.TrainingType
	if err := db.getDocumentByID(trainingType, id).Decode(&doc); err != nil {
		return entities.TrainingType{}, nil
	}
	return doc, nil
}

// GetExerciseByID ...
func (db mongoDatabase) GetExerciseByID(id int) (entities.Exercise, error) {
	var doc entities.Exercise
	if err := db.getDocumentByID(exercise, id).Decode(&doc); err != nil {
		return entities.Exercise{}, nil
	}
	return doc, nil
}

// GetProgramByID ...
func (db mongoDatabase) GetProgramByID(id int) (entities.Program, error) {
	var doc entities.Program
	if err := db.getDocumentByID(program, id).Decode(&doc); err != nil {
		return entities.Program{}, nil
	}
	return doc, nil
}

// GetEquipmentByID ...
func (db mongoDatabase) GetEquipmentByID(id int) (entities.Equipment, error) {
	var doc entities.Equipment
	if err := db.getDocumentByID(equipment, id).Decode(&doc); err != nil {
		return entities.Equipment{}, nil
	}
	return doc, nil
}

// GetAllProgramCards ...
func (db mongoDatabase) GetAllProgramCards() ([]entities.ProgramCard, error) {
	c := db.client.Database(dbName).Collection(program)

	filter := bson.D{}
	options := options.Find().SetProjection(bson.D{{
		Key:   "description",
		Value: 0,
	}})

	cursor, err := c.Find(db.ctx, filter, options)
	if err != nil {
		return nil, err
	}
	var result []entities.ProgramCard
	if err = cursor.All(db.ctx, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// // GetAllPrograms ...
// func (db mongoDatabase) GetAllPrograms() ([]entities.Program, error) {
// 	c := db.client.Database(dbName).Collection(program)

// 	filter := bson.D{}
// 	options := options.Find()

// 	cursor, err := c.Find(db.ctx, filter, options)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var result []entities.Program
// 	if err = cursor.All(db.ctx, &result); err != nil {
// 		return nil, err
// 	}
// 	return result, nil
// }

// GetAllExerciseCards ...
func (db mongoDatabase) GetAllExerciseCards() ([]entities.ExerciseCard, error) {
	c := db.client.Database(dbName).Collection(exercise)

	filter := bson.D{}
	options := options.Find().SetProjection(bson.D{
		{Key: "description", Value: 0},
		{Key: "tips", Value: 0},
		{Key: "variations", Value: 0},
	})

	cursor, err := c.Find(db.ctx, filter, options)
	if err != nil {
		return nil, err
	}
	var result []entities.ExerciseCard
	if err = cursor.All(db.ctx, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetAllDifficulties ...
func (db mongoDatabase) GetAllDifficulties() ([]entities.Difficulty, error) {
	c := db.client.Database(dbName).Collection(difficulty)

	cursor, err := c.Find(db.ctx, bson.D{}, options.Find())
	if err != nil {
		return nil, err
	}
	var result []entities.Difficulty
	if err = cursor.All(db.ctx, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetAllEquipment ...
func (db mongoDatabase) GetAllEquipment() ([]entities.Equipment, error) {
	c := db.client.Database(dbName).Collection(equipment)

	cursor, err := c.Find(db.ctx, bson.D{}, options.Find())
	if err != nil {
		return nil, err
	}
	var result []entities.Equipment
	if err = cursor.All(db.ctx, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetAllExerciseTypes ...
func (db mongoDatabase) GetAllExerciseTypes() ([]entities.ExerciseType, error) {
	c := db.client.Database(dbName).Collection(exerciseType)

	cursor, err := c.Find(db.ctx, bson.D{}, options.Find())
	if err != nil {
		return nil, err
	}
	var result []entities.ExerciseType
	if err = cursor.All(db.ctx, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetAllMuscles ...
func (db mongoDatabase) GetAllMuscles() ([]entities.Muscle, error) {
	c := db.client.Database(dbName).Collection(muscle)

	cursor, err := c.Find(db.ctx, bson.D{}, options.Find())
	if err != nil {
		return nil, err
	}
	var result []entities.Muscle
	if err = cursor.All(db.ctx, &result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetAllProgramTypes ...
func (db mongoDatabase) GetAllProgramTypes() ([]entities.ProgramType, error) {
	c := db.client.Database(dbName).Collection(programType)

	cursor, err := c.Find(db.ctx, bson.D{}, options.Find())
	if err != nil {
		return nil, err
	}
	var result []entities.ProgramType
	if err = cursor.All(db.ctx, &result); err != nil {
		return nil, err
	}
	return result, nil
}
