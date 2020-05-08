package core

import (
	"context"
	"errors"

	"github.com/go-kit/kit/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io.giftano.api/go_core/config"
)

var e = config.NewConfig()
var errFoo = errors.New("Unable to handle Repo Request")
var coll = []string{e.ServiceName + "_items"}

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

func (repo *repo) CreateItem(item Item) (stringID string, err error) {

	if item.ID != nil {
		return stringID, errFoo
	}

	itemRes, err := repo.collection.InsertOne(repo.ctx, item)
	if err != nil {
		return stringID, err
	}
	insertedID, ok := itemRes.InsertedID.(primitive.ObjectID)
	if !ok {
		return stringID, errFoo
	}
	stringID = insertedID.Hex()
	return stringID, nil
}

func (repo *repo) GetItemByID(id string) (item Item, err error) {

	if id == "" {
		return item, errFoo
	}

	rID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{
		"_id":       rID,
		"IsDeleted": false,
	}
	err = repo.collection.FindOne(repo.ctx, filter).Decode(&item)
	return
}

func (repo *repo) GetAllItems() (results []Item, err error) {
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
	for cur.Next(repo.ctx) {

		// create a value into which the single document can be decoded
		var elem Item
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}

		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	// Close the cursor once finished
	cur.Close(repo.ctx)

	return results, nil
}

func (repo *repo) UpdateItem(id string, update interface{}) (item Item, err error) {
	if id == "" {
		return item, errFoo
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
	errs := repo.collection.FindOneAndUpdate(repo.ctx, filter, updates, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&item)

	if errs != nil {
		err = errs
	}

	return item, err
}
