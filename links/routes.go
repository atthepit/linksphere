package links

import "github.com/gin-gonic/gin"

func RegisterHandlers(e *gin.Engine) {
	links := e.Group("/links")
	{
		links.POST("", CreateLinkHandler)
		links.GET("", GetAllLinksHandler)
	}
}
