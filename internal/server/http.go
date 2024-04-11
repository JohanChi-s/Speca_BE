package server

import (
	apiV1 "core_service/api/v1"
	"core_service/docs"
	"core_service/internal/handler"
	"core_service/internal/middleware"
	"core_service/pkg/jwt"
	"core_service/pkg/log"
	"core_service/pkg/server/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewHTTPServer(
	logger *log.Logger,
	conf *viper.Viper,
	jwt *jwt.JWT,
	userHandler *handler.UserHandler,
	// teamHandler *handler.TeamHandler,
	// workspaceHandler *handler.WorkspaceHandler,
	// documentHandler *handler.DocumentHandler,
	// memberHanler *handler.MemberHandler,
	// actionEventHandler *handler.ActionEventHandler,
) *http.Server {
	gin.SetMode(gin.DebugMode)
	s := http.NewServer(
		gin.Default(),
		logger,
		http.WithServerHost(conf.GetString("http.host")),
		http.WithServerPort(conf.GetInt("http.port")),
	)

	// swagger doc
	docs.SwaggerInfo.BasePath = "/v1"
	s.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerfiles.Handler,
		// ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", conf.GetInt("app.http.port"))),
		ginSwagger.DefaultModelsExpandDepth(-1),
	))

	s.Use(
		middleware.CORSMiddleware(),
		middleware.ResponseLogMiddleware(logger),
		middleware.RequestLogMiddleware(logger),
		//middleware.SignMiddleware(log),
	)
	s.GET("/", func(ctx *gin.Context) {
		logger.WithContext(ctx).Info("hello")
		apiV1.HandleSuccess(ctx, map[string]interface{}{
			":)": "Thank you for using nunu!",
		})
	})

	v1 := s.Group("/v1")
	{
		// No route group has permission
		noAuthRouter := v1.Group("/")
		{
			noAuthRouter.POST("/register", userHandler.Register)
			noAuthRouter.POST("/login", userHandler.Login)
		}
		// Non-strict permission routing group
		v1.Group("/").Use(middleware.NoStrictAuth(jwt, logger))
		{
			userRouter := v1.Group("/user").Use(middleware.NoStrictAuth(jwt, logger))
			{
				userRouter.GET("/:id", userHandler.GetUserByID)
				userRouter.PUT("/:id", userHandler.UpdateProfile)
				userRouter.DELETE("/:id", userHandler.DeleteUserByID)
			}
		}

		// Strict permission routing group
		strictAuthRouter := v1.Group("/").Use(middleware.StrictAuth(jwt, logger))
		{
			strictAuthRouter.PUT("/user", userHandler.UpdateProfile)
		}
	}

	return s
}
