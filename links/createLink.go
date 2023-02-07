package links

import (
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateLinkRequest struct {
	Url         string `form:"url" binding:"required"`
	Title       string `form:"title" binding:"required"`
	Description string `form:"description"`
}

type CreateLinkResponse struct {
	Id uuid.UUID
}

func CreateLinkHandler(c *gin.Context) {
	var request CreateLinkRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	url, err := url.Parse(request.Url)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Couldn't parse new link URL",
		})
	}

	newLink := Link{
		Id:          uuid.New(),
		Url:         url,
		Title:       request.Title,
		Description: request.Description,
		CreatedOn:   time.Now(),
	}

	AllLinks = append(AllLinks, newLink)

	c.JSON(200, CreateLinkResponse{
		Id: newLink.Id,
	})
}
