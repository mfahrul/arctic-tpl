package core

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/log"
	"go.mongodb.org/mongo-driver/mongo"
	"{{.Projectpath}}/route"
)

var server = &route.Server{
	Route: router,
}

func init() {
	route.NewServer.Add(server)
}

func router(ctx context.Context, r *gin.Engine, db *mongo.Database, logger log.Logger) {

	var coreSrv Service
	{
		coreRepo := NewRepo(ctx, db, logger)
		coreSrv = NewService(coreRepo, logger)
	}

	endpoints := MakeEndpoints(coreSrv)

	r.POST("/item", endpoints.CreateItem)
	r.GET("/items", endpoints.GetAllItems)
	r.GET("/item/:ID", endpoints.GetItemByID)
	r.PUT("/item/:ID", endpoints.UpdateItem)
	r.DELETE("/item/:ID", endpoints.DeleteItem)
}
