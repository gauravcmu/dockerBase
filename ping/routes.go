package ping

import (
	"github.com/gin-gonic/gin"
)

type Pong struct {
	Pong string `json:"pong"`
}

func Routes(rtgroup *gin.RouterGroup) {
	rtgroup.GET("/ping", func(c *gin.Context) {
		c.JSON(200, Pong{
			Pong: "true",
		})
	})
}
