package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"{{.Projectpath}}/config/database"
	{{ with $Ppath := .Projectpath }}
		{{ with .Modules }}
			{{ range . }}
				_ "{{$Ppath}}/src/{{.Name}}"
			{{ end }}
		{{ end }}
	{{ end }}

	"{{.Projectpath}}/route"
)

// @title Giftano Core API Docs
// @version 0.1.1
// @description Dashboard users management service.
// @termsOfService http://swagger.io/terms/

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name apikey

// @host localhost:8080
// @BasePath /

func main() {
	var httpAddr = flag.String("http", ":8080", "http listen address")

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "users",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	flag.Parse()

	ctx := context.Background()

	db := database.MongoConnection(ctx, logger)
	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		fmt.Println("listening on port", *httpAddr)
		errs <- route.Run(ctx, db, logger)
	}()

	level.Error(logger).Log("exit", <-errs)
}
