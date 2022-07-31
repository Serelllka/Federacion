package handler

import (
	"github.com/Serelllka/Federacion/pkg/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	corsConfig := cors.Config{
		AllowOrigins:     viper.GetStringSlice("cors.allowOrigins"),
		AllowMethods:     viper.GetStringSlice("cors.allowMethods"),
		AllowHeaders:     viper.GetStringSlice("cors.allowHeaders"),
		ExposeHeaders:    viper.GetStringSlice("cors.exposeHeaders"),
		AllowCredentials: viper.GetBool("cors.allowCredentials"),
		MaxAge:           viper.GetDuration("cors.maxAge"),
	}

	router.Use(cors.New(corsConfig))
	router.Use(static.Serve("/", static.LocalFile("./views", true)))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	users := router.Group("/users")
	{
		users.GET("/", h.getAllUsers)
	}

	api := router.Group("/api", h.userIdentity)
	{
		articles := api.Group("/articles")
		{
			articles.GET("/", h.GetAllArticles)
			articles.GET("/:id", h.GetArticle)
			articles.POST("/", h.AddArticle)
		}

		condem := api.Group("/condem")
		{
			condem.GET("/", h.GetAllCondemnations)
			condem.POST("/", h.AddCondemnation)
		}

		userInfo := api.Group("/user-info")
		{
			userInfo.GET("/")
			userInfo.GET("/:id")
			userInfo.POST("/")
		}
	}

	return router
}
