package controller

import (
	"blog/response"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"path"
	"time"
)

func Upload(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		response.Response(c, http.StatusInternalServerError, false, nil, "格式错误")
		return
	}
	filename := header.Filename
	ext := path.Ext(filename)
	name := "image_" + time.Now().Format("20060102150405")
	newFilename := name + ext
	out, err := os.Create("static/images/" + newFilename)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, false, nil, "保存失败")
		return
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, false, nil, "复制错误")
		return
	}
	response.Success(c, gin.H{"filePath": "/images/" + newFilename}, "上传成功")
}
