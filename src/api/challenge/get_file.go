package challenge

import (
	"fmt"
	challenge_ "src/logic/challenge"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Challenge) GetFile(c *gin.Context) {
	//token, 题目名称, name
	//如果存在附件
	//在响应中返回文件内容
	challenge_id := c.PostForm("challenge_id")
	challengfid, err := strconv.Atoi(challenge_id)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "challenge_id error",
		})
	}
	var ch challenge_.Challenge
	ch.ID = uint(challengfid)
	filename, err := ch.GetFile()
	if err != nil {
		c.JSON(400, gin.H{
			"message": err,
		})
		return
	}
	fmt.Println(filename)
	c.Header("fileName", filename)
	c.Header("msg", "文件下载成功")
	c.File(filename)
}
