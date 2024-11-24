package team

import (
	"net/http"
	team_ "src/logic/team"

	"github.com/gin-gonic/gin"
)

func (s *Team) Create(c *gin.Context) {
	// token,name,team
	var info struct {
		Name string `form:"name" json:"name" binding:"required"`
		Team string `form:"team" json:"team" binding:"required"`
	}
	err := c.ShouldBind(&info)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
		})
		panic(err)
	}
	t := team_.GetTeam()
	t.Create()
	// yes or no
	//if no, message
}
