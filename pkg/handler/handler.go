package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/whyvilos/mybox/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{services: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	api := router.Group("/api", h.userIdentity)
	{

		user := api.Group("/:id_user")
		{

			user.GET("/", h.getUserProfile) //получить сраничку пользователя

			feed := user.Group("/posts")
			{
				feed.POST("/", h.createPost)
				feed.GET("/", h.getAllPosts)
				feed.GET("/:id_post", h.getPostById)   //TODO
				feed.PUT("/:id_post", h.updatePost)    //TODO
				feed.DELETE("/:id_post", h.deletePost) //TODO
			}

			catalog := user.Group("/catalog")
			{
				catalog.POST("/", h.createItem)
				catalog.GET("/", h.getAllItems)
				catalog.GET("/:id_item", h.getItemById)   //TODO
				catalog.PUT("/:id_item", h.updateItem)    //TODO
				catalog.DELETE("/:id_item", h.deleteItem) //TODO

			}
		}

		api.GET("/ping", h.testAuth)
		api.GET("/get-id", h.getId)
	}

	return router
}
