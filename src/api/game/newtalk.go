package game

import "github.com/gin-gonic/gin"

func (s *Game) NewTalk(c *gin.Context) {
	// token,name,消息内容,比赛id或者0表示主界面

	//正确
	c.JSON(200, gin.H{
		"code":    0,
		"message": "success",
	})
}
