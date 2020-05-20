package [[ with .ModuleToParse ]][[.Name]][[ end ]]

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/log"
	"go.mongodb.org/mongo-driver/mongo"
	"[[.Projectpath]]/route"
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
	[[ with .ModuleToParse.Model ]]
	r.POST("/[[.Name | ToLower | ToSingular]]", endpoints.Create[[.Name | ToCamel]])
	r.GET("/[[.Name | ToLower | ToPlural]]", endpoints.GetAll[[.Name | ToCamel | ToPlural]])
	r.GET("/[[.Name | ToLower | ToSingular]]/:ID", endpoints.Get[[.Name | ToCamel]]ByID)
	r.PUT("/[[.Name | ToLower | ToSingular]]/:ID", endpoints.Update[[.Name | ToCamel]])
	r.DELETE("/[[.Name | ToLower | ToSingular]]/:ID", endpoints.Delete[[.Name | ToCamel]]) [[ end ]]
}
