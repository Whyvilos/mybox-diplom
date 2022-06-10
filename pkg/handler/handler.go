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
		api.GET("/line", h.getLine)
		api.GET("/favorite", h.getFavorite)
		user := api.Group("/:id_user")
		{
			user.GET("/", h.getUserProfile)
			user.POST("/upload_avatar", h.postUploadAvatar)
			user.POST("/follow", h.postFollow)
			user.GET("/check_follow", h.getCheckFollow)
			user.POST("/unfollow", h.postUnFollow)
			user.POST("/favorite/:id_item", h.postAddFavorite)
			user.POST("/unfavorite/:id_item", h.postDeleteFavorite)
			user.GET("/favorite/:id_item/check", h.postCheckFavorite)
			feed := user.Group("/posts")
			{
				feed.POST("/", h.createPost)
				feed.GET("/", h.getAllPosts)
				feed.POST("/upload_post_media", h.postUploadPostMedia)
				feed.GET("/:id_post", h.getPostById)
				feed.PUT("/:id_post", h.updatePost)
				feed.DELETE("/:id_post", h.deletePost)
			}

			catalog := user.Group("/catalog")
			{
				catalog.POST("/", h.createItem)
				catalog.GET("/", h.getAllItems)
				catalog.POST("/upload_item_media", h.postUploadItemMedia)
				catalog.GET("/:id_item", h.getItemById)
				catalog.PUT("/:id_item", h.updateItem)
				catalog.DELETE("/:id_item", h.deleteItem)

			}
		}
		order := api.Group("/order")
		{
			order.POST("/", h.postCreateOrder)
			order.GET("/", h.getOrders)
			order.GET("/shop", h.getOrdersForYou)
			order.PUT("/:id_order/:status", h.putOrderStatus)
		}
		notice := api.Group("/notice")
		{
			notice.POST("/", h.postCreateNotice)
			notice.GET("/", h.getNotices)
			notice.GET("/check", h.getCheckNotice)
		}
		chat := api.Group("/chat") //TODO
		{
			chat.GET("/:id_chat", h.getChat)
			chat.GET("/find/:id_order", h.getFindChat)
			chat.GET("/find2/:id_order", h.getFindChat2)
			chat.POST("/:id_chat/", h.postSendMessage)
		}
		api.GET("/ping", h.ping)
		api.GET("/get-id", h.getId)
	}
	return router
}
