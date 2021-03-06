package route

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/log"

	"github.com/gin-gonic/gin" // gin-swagger middleware
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"go.mongodb.org/mongo-driver/mongo"
	"[[.Projectpath]]/config"
)

//Server struct
type Server struct {
	servers []*Server
	Route   func(ctx context.Context, r *gin.Engine, db *mongo.Database, logger log.Logger)
}

//NewServer for routing the server
var NewServer = &Server{}

//Add add one or more servers
func (s *Server) Add(servers ...*Server) {
	for _, srv := range servers {
		s.servers = append(s.servers, srv)
	}
}

//Run Server
func Run(ctx context.Context, db *mongo.Database, logger log.Logger) error {
	e := config.NewConfig()
	r := gin.New()
	r.Use(gin.Recovery())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"giftano":      "Welcome",
			"service_name": e.ServiceName + "_api",
			"api_version":  e.Version,
		})
	})

	for _, srv := range NewServer.servers {
		srv.Route(ctx, r, db, logger)
	}

	url := ginSwagger.URL("doc.json")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.NoRoute(func(c *gin.Context) {
		c.JSON(200, gin.H{"code": http.StatusNotFound, "message": "PAGE_NOT_FOUND"})
	})

	return r.Run() // listen and serve on 0.0.0.0:8080
}
