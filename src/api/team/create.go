package team

import (
	"net/http"
	team_ "src/logic/team"

	"github.com/gin-gonic/gin"
)

func (s *Team) Create(c *gin.Context) {
	// token,name,team
	var info Team
	err := c.ShouldBind(&info)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
		})
		panic(err)
	}
	t := team_.GetTeam()
	t.Name = info.Name
	t.Leader = info.Leader
	t.Desc = info.Desc
	t.GameID = info.GameID
	status, key, err := t.Create()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "system error",
		})
	}
	if status == 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "database error",
		})
	} else if status == 2 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "team already exists",
		})
	} else if status == 3 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "leader not exists",
		})
	} else if status == 4 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "leader already in game",
		})
	} else if status == 0 {

		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"key":     key, //加入队伍的key
		})
	}
}
