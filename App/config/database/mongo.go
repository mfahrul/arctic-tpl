package database

import (
	"context"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"io.giftano.api/go_core/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoConnection for mongodb
func MongoConnection(ctx context.Context, logger log.Logger) *mongo.Database {
	e := config.NewConfig()
	clientOptions := options.Client().ApplyURI("mongodb://" + e.DbUsername + ":" + e.DbPassword + "@" + e.DbHost + ":" + e.DbPort)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		level.Error(logger).Log("exit", err)
		os.Exit(-1)
	}

	// Check the connection
	err = client.Ping(ctx, nil)

	if err != nil {
		level.Error(logger).Log("exit", err)
		os.Exit(-1)
	}

	level.Info(logger).Log("msg", "Connected to MongoDB!")
	return client.Database(e.DbName)
}
