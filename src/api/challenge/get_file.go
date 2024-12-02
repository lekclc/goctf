package challenge

import (
	"fmt"
	challenge_ "src/logic/challenge"

	"github.com/gin-gonic/gin"
)

func (s *Challenge) GetFile(c *gin.Context) {
	//token, 题目名称, name
	//如果存在附件
	//在响应中返回文件内容
	challenge_name := c.PostForm("challenge_name")
	ch := challenge_.GetChallenge()
	ch.Name = challenge_name
	filename, err := ch.GetFile()
	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
		return
	}
	fmt.Println(filename)
	//后面可以做个下载服务优化掉
	c.JSON(200, gin.H{
		"message":  "success",
		"filename": filename,
	})
	c.File(filename)

}
