package note

import (
	note_ "src/logic/note"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Note) GetNote(c *gin.Context) {
	id_ := c.PostForm("id")
	id, err := strconv.ParseUint(id_, 10, 64)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": err,
		})
	}
	var n note_.Note
	res, err := n.GetNote(uint(id))
	if err != nil {
		c.JSON(200, gin.H{
			"code":    1,
			"message": err,
		})
	}
	c.JSON(200, gin.H{
		"code":    0,
		"message": "success",
		"data":    res,
	})
}
