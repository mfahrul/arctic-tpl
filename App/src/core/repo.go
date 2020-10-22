package [[ with .ModuleToParse ]][[.Name]][[ end ]]

import (
	"context"
	"errors"

	"github.com/go-kit/kit/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"[[.Projectpath]]/config"
)
[[ with .ModuleToParse.Model ]]
var e = config.NewConfig()
var errFoo = errors.New("Unable to handle Repo Request")
var coll = []string{e.ServiceName + "_[[.Name | ToLower | ToPlural]]"}

type repo struct {
	collection *mongo.Collection
	logger     log.Logger
	ctx        context.Context
}

//NewRepo function
func NewRepo(ctx context.Context, db *mongo.Database, logger log.Logger) Repository {

	return &repo{
		collection: db.Collection(coll[0]),
		logger:     log.With(logger, "repo", "mongo"),
		ctx:        ctx,
	}
}

func (repo *repo) Create[[.Name | ToCamel]]([[.Name | ToLower]] [[.Name | ToCamel]]) (stringID string, err error) {

	if [[.Name | ToLower]].ID != nil {
		return stringID, errFoo
	}

	[[.Name | ToLower]]Res, err := repo.collection.InsertOne(repo.ctx, [[.Name | ToLower]])
	if err != nil {
		return stringID, err
	}
	insertedID, ok := [[.Name | ToLower]]Res.InsertedID.(primitive.ObjectID)
	if !ok {
		return stringID, errFoo
	}
	stringID = insertedID.Hex()
	return stringID, nil
}

func (repo *repo) Get[[.Name | ToCamel]]ByID(id string) ([[.Name | ToLower]] [[.Name | ToCamel]], err error) {

	if id == "" {
		return [[.Name | ToLower]], errFoo
	}

	rID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{
		"_id":       rID,
		"IsDeleted": false,
	}
	err = repo.collection.FindOne(repo.ctx, filter).Decode(&[[.Name | ToLower]])
	return
}

func (repo *repo) GetAll[[.Name | ToCamel | ToPlural]]() (results [][[.Name | ToCamel]], err error) {
	filter := []bson.M{
		{"$match": bson.M{"IsDeleted": false}},
		{
			"$sort": bson.M{
				"CreatedAt": -1,
			},
		},
	}
	cur, err := repo.collection.Aggregate(repo.ctx, filter)
	if err != nil {
		return nil, err
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	
	// for cur.Next(repo.ctx) {

	// 	// create a value into which the single document can be decoded
	// 	var elem [[.Name | ToCamel]]
	// 	err := cur.Decode(&elem)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	results = append(results, elem)
	// }

	// if err := cur.Err(); err != nil {
	// 	return nil, err
	// }

	if err := cur.All(repo.ctx, &results); err != nil {
		return nil, err
	}

	// Close the cursor once finished
	// cur.Close(repo.ctx)

	return results, nil
}

func (repo *repo) Update[[.Name | ToCamel]](id string, update interface{}) ([[.Name | ToLower]] [[.Name | ToCamel]], err error) {
	if id == "" {
		return [[.Name | ToLower]], errFoo
	}

	rID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{
		"_id":       rID,
		"IsDeleted": false,
	}
	// doc, errc := utils.ToDoc(update)

	// if err != nil {
	// 	err = errc
	// }

	updates := bson.D{{Key: "$set", Value: update}}
	errs := repo.collection.FindOneAndUpdate(repo.ctx, filter, updates, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&[[.Name | ToLower]])

	if errs != nil {
		err = errs
	}

	return [[.Name | ToLower]], err
}[[ end ]]
