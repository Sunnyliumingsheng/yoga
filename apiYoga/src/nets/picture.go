package nets

import (
	yangoss "github.com/Sunnyliumingsheng/yangoss"
	"github.com/gin-gonic/gin"
)

func uploadPicture(c *gin.Context) {
	handler, err := c.FormFile("file")
	if err != nil {
		c.String(400, "获取照片失败")
		return
	}
	file, err := handler.Open()
	if err != nil {
		c.String(400, "读取照片失败")
		return
	}
	url, err := yangoss.UploadFile(file, handler)
	if err != nil {
		c.String(400, "上传照片失败")
		return
	}
	c.String(200, url)
}
func deletePicture(c *gin.Context) {
	type PictureData struct {
		Url string `json:"url"`
	}
	var getData PictureData
	if err := c.ShouldBindJSON(&getData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := yangoss.RemoveFile(getData.Url)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.String(200, "删除成功")
}
