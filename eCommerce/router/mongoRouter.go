package router

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)


type MongoRouter struct {
	router *Router
}

func NewMongoRouter(router *Router) {
	m := &MongoRouter{
		router: router,
	}

	baseUri := "/mongo"

	group := m.router.engine.Group(baseUri)
	group.GET("/health", m.health)
}

func (m *MongoRouter) health(c *gin.Context) {
	if c.Request.Context().Err() == context.DeadlineExceeded {
		fmt.Println("에러가 났습니다.")
	} else {
		fmt.Println("들어옵니다.")
	}

	if !c.Writer.Written() {
		c.JSON(200, "test")
	}
}