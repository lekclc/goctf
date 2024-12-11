package challenge

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"path/filepath"
	challenge_ "src/logic/challenge"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *Challenge) UpdateFile(c *gin.Context) {
	// 获取上传的文件
	challenge_id := c.PostForm("challenge_id")
	challengeid, _ := strconv.Atoi(challenge_id)
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	var challenge challenge_.Challenge

	splitFilename := strings.Split(file.Filename, ".")
	end_file := splitFilename[len(splitFilename)-1]
	start_file := file.Filename[:len(file.Filename)-len(end_file)-1]
	hash := md5.Sum([]byte(file.Filename + time.Now().String()))
	filename := start_file + "_" + fmt.Sprintf("%x", hash) + "." + end_file
	savePath := filepath.Join("data/file/", filename)
	fmt.Println(savePath)
	// 保存文件到指定路径
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	err = challenge.UploadFile(uint(challengeid), filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
