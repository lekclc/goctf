package note

import (
	note_ "src/logic/note"

	"github.com/gin-gonic/gin"
)

func (s *Note) CreateNote(c *gin.Context) {
	name := c.PostForm("name")
	NoteName := c.PostForm("note_name")
	var n note_.Note
	path, err := n.CreateNote(name, NoteName)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		return
	}
	// 保存文件到指定路径
	if err := c.SaveUploadedFile(file, path); err != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    0,
		"message": "success",
	})
}
