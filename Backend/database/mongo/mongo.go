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
func NewMongoDatabase(
	ctx context.Context,
	cfg Config,
) (database.Database, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(fmt.Sprintf("%s%s", cfg.Host, cfg.Port)).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, err
	}
	return mongoDatabase{client: client}, nil
}

func (db mongoDatabase) getDocumentByID(
	ctx context.Context,
	collection string,
	id int,
) *mongo.SingleResult {
	c := db.client.Database(dbName).Collection(collection)

	filter := bson.D{{Key: "_id", Value: id}}
	result := c.FindOne(ctx, filter, options.FindOne())
	return result
}

func (db mongoDatabase) GetExerciseByID(ctx context.Context, id int) (entities.Exercise, error) {
	res := db.getDocumentByID(ctx, exercise, id)
	var doc entities.Exercise
	if err := res.Decode(&doc); err != nil {
		return entities.Exercise{}, nil
	}
	return doc, nil
}
func (db mongoDatabase) GetDifficultyByID(ctx context.Context, id int) (entities.Difficulty, error) {
	res := db.getDocumentByID(ctx, difficulty, id)
	var doc entities.Difficulty
	if err := res.Decode(&doc); err != nil {
		return entities.Difficulty{}, nil
	}
	return doc, nil
}
func (db mongoDatabase) GetEquipmentByID(ctx context.Context, id int) (entities.Equipment, error) {
	res := db.getDocumentByID(ctx, equipment, id)
	var doc entities.Equipment
	if err := res.Decode(&doc); err != nil {
		return entities.Equipment{}, nil
	}
	return doc, nil
}
func (db mongoDatabase) GetExerciseTypeByID(ctx context.Context, id int) (entities.ExerciseType, error) {
	res := db.getDocumentByID(ctx, exerciseType, id)
	var doc entities.ExerciseType
	if err := res.Decode(&doc); err != nil {
		return entities.ExerciseType{}, nil
	}
	return doc, nil
}
func (db mongoDatabase) GetMuscleByID(ctx context.Context, id int) (entities.Muscle, error) {
	res := db.getDocumentByID(ctx, muscle, id)
	var doc entities.Muscle
	if err := res.Decode(&doc); err != nil {
		return entities.Muscle{}, nil
	}
	return doc, nil
}
func (db mongoDatabase) GetProgramByID(ctx context.Context, id int) (entities.Program, error) {
	res := db.getDocumentByID(ctx, program, id)
	var doc entities.Program
	if err := res.Decode(&doc); err != nil {
		return entities.Program{}, nil
	}
	return doc, nil
}
func (db mongoDatabase) GetProgramTypeByID(ctx context.Context, id int) (entities.ProgramType, error) {
	res := db.getDocumentByID(ctx, programType, id)
	var doc entities.ProgramType
	if err := res.Decode(&doc); err != nil {
		return entities.ProgramType{}, nil
	}
	return doc, nil
}
func (db mongoDatabase) GetTrainingTypeByID(ctx context.Context, id int) (entities.TrainingType, error) {
	res := db.getDocumentByID(ctx, trainingType, id)
	var doc entities.TrainingType
	if err := res.Decode(&doc); err != nil {
		return entities.TrainingType{}, nil
	}
	return doc, nil
}
