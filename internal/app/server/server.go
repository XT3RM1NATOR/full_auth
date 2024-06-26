package server

import (
	"github.com/XT3RM1NATOR/full_auth/config"
	_ "github.com/XT3RM1NATOR/full_auth/docs"
	authDelivery "github.com/XT3RM1NATOR/full_auth/internal/user/delivery"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/minio/minio-go/v7"
	"github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
)

// RunHTTPServer
// @title PointAI
// @version 1.0
// @description This is the backend server for PointAI.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name .
// @license.url .
// @host petstore.swagger.io
// @externalDocs.description  OpenAPI 2.0
// @BasePath /
func RunHTTPServer(cfg *config.Config, db *mongo.Database, str *minio.Client) {
	e := echo.New()

	logger := logrus.New()
	logger.Out = os.Stdout

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339_nano} [${status}] ${method} ${uri} (${latency_human})\n",
		Output: logger.Out,
	}))

	e.Use(middleware.CORS())
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	authDelivery.RegisterAuthRoutes(e, cfg, db, str)

	if err := e.Start(cfg.Server.Port); err != nil {
		panic(err)
	}
}
