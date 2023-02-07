package links

import (
	"net/url"
	"time"

	"github.com/google/uuid"
)

type Link struct {
	Id          uuid.UUID
	Url         *url.URL
	Title       string
	Description string
	CreatedOn   time.Time
}

var AllLinks = make([]Link, 0)
