package game

import "github.com/gin-gonic/gin"

func (s *Game) GetTalk(c *gin.Context) {
	//比赛id或者0表示主界面
	//先不做了吧
	res := []string{}

	res = append(res, "talk one")
	res = append(res, "talk two")
	c.JSON(200, gin.H{
		"code":    0,
		"message": "success",
		"data":    res,
	})
}
