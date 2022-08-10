package serv

import (
	"github.com/Doozers/debug-api/types"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/debug", func(c *gin.Context) {
		d := new(types.DebugRequest)
		err := c.ShouldBindJSON(d)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, d)
	})
	r.Run()
}
