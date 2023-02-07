package links

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type link struct {
	Id          uuid.UUID
	Url         string
	Title       string
	Description string
	CreatedOn   time.Time
}
type GetAllLinksResponse struct {
	Data []link
}

func GetAllLinksHandler(c *gin.Context) {
	var allLinks []link
	for _, current := range AllLinks {
		l := link{
			Id:          current.Id,
			Url:         current.Url.String(),
			Title:       current.Title,
			Description: current.Description,
			CreatedOn:   current.CreatedOn,
		}
		allLinks = append(allLinks, l)
	}

	c.JSON(200, GetAllLinksResponse{
		Data: allLinks,
	})
}
