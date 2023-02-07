package main

import (
	"github.com/atthepit/linksphere-api/links"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	links.RegisterHandlers(r)
	r.Run()
}
