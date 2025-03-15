package note

import (
	note_ "src/logic/note"

	"github.com/gin-gonic/gin"
)

func (s *Note) GetNoteList(c *gin.Context) {
	//name 接受参数用户名,置空获取所以note
	name := c.PostForm("name")
	var n note_.Note
	res, err := n.GetNoteList(name)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    0,
		"message": "success",
		"data":    res,
	})

}
