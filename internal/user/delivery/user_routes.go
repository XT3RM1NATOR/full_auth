package authDelivery

import (
	"github.com/XT3RM1NATOR/full_auth/config"
	"github.com/XT3RM1NATOR/full_auth/internal/user/delivery/controller"
	"github.com/XT3RM1NATOR/full_auth/internal/user/infrastructure/client"
	"github.com/XT3RM1NATOR/full_auth/internal/user/infrastructure/repository"
	"github.com/XT3RM1NATOR/full_auth/internal/user/service"
	"github.com/XT3RM1NATOR/full_auth/middleware"
	"github.com/labstack/echo/v4"
	"github.com/minio/minio-go/v7"
	"go.mongodb.org/mongo-driver/mongo"
)

func RegisterAuthRoutes(e *echo.Echo, cfg *config.Config, db *mongo.Database, str *minio.Client) {
	ur := repository.NewUserRepositoryImpl(db, cfg)
	ec := client.NewEmailClientImpl(cfg.Email.SMTPUsername, cfg.Email.SMTPPassword, cfg.Email.SMTPHost, cfg.Email.SMTPPort)
	sc := client.NewStorageClientImpl(str)
	es := service.NewEmailServiceImpl(ec)
	us := service.NewUserServiceImpl(ur, sc, es, cfg)
	uc := controller.NewUserController(us, cfg)

	userGroup := e.Group("/user")
	userGroup.POST("/signup", uc.RegisterUser)
	userGroup.POST("/verify/:token", uc.ConfirmUser)
	userGroup.POST("/signin", uc.Login)
	userGroup.POST("/logout", uc.Logout, middleware.ValidateAccessTokenMiddleware(cfg.Auth.JWTSecretKey))
	userGroup.POST("/recover", uc.ForgotPassword)
	userGroup.POST("/reset", uc.ResetPassword)
	userGroup.PUT("/renew", uc.RenewAccessToken)
	userGroup.GET("/profile", uc.GetProfile, middleware.ValidateAccessTokenMiddleware(cfg.Auth.JWTSecretKey))
	userGroup.PUT("/profile", uc.UpdateProfile, middleware.ValidateAccessTokenMiddleware(cfg.Auth.JWTSecretKey))

	oAuth2Group := e.Group("/oauth2")
	oAuth2Group.GET("/google/callback", uc.GoogleCallback)
	oAuth2Group.GET("/google/tokens", uc.GoogleTokens)
	oAuth2Group.GET("/facebook/callback", uc.FacebookCallback)
}
