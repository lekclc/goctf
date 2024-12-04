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
	db := con.Db.Db
	err := db.Where("name = ?", name).First(&info).Error

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":   "success",
		"name":      info.Name,
		"admin":     info.Admin,
		"Team":      info.Team,
		"CreatedAt": info.CreatedAt,
	})

}
