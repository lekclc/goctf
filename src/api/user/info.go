package user

import (
	"net/http"
	con "src/const"
	"src/database"

	"github.com/gin-gonic/gin"
)

// 获取个人信息
func (s *User) Info(c *gin.Context) {
	//token,name
	name := c.PostForm("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
		})
		return
	}
	var info database.User
	info.Name = name
	db := con.Db.Db
	db.Where("name = ?", name).First(&info)
	c.JSON(http.StatusOK, gin.H{
		"message":   "success",
		"name":      info.Name,
		"admin":     info.Admin,
		"Team":      info.Team,
		"CreatedAt": info.CreatedAt,
	})

}
