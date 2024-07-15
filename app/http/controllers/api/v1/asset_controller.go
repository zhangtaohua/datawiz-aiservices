package v1

import (
	"bytes"
	"datawiz-aiservices/pkg/config"
	"datawiz-aiservices/pkg/file"
	"datawiz-aiservices/pkg/response"
	"fmt"
	"io"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AssetsController struct {
	BaseAPIController
}

func forwardToOtherService(c *gin.Context, targetUrl, path, rawQuery string) {

	// 创建一个反向代理的目标URL
	targetURL, _ := url.Parse(targetUrl)

	// 创建一个反向代理器
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	// 修改请求的Host，以便在目标服务器上找到正确的虚拟主机
	c.Request.Host = targetURL.Host

	c.Request.URL.Path = path
	c.Request.URL.RawQuery = rawQuery

	// 使用反向代理器处理请求
	proxy.ServeHTTP(c.Writer, c.Request)
}

func (ctrl *AssetsController) Download(c *gin.Context) {
	// if !validSignature(c) {
	// 	c.AbortWithStatus(http.StatusUnauthorized)
	// 	return
	// }

	file, _ := c.GetQuery("file")
	if file == "" {
		return
	}

	// 判断是否需要进行转发
	// if strings.HasPrefix(file, "/"+common.DataSourceXR07) {
	// 	// 转发请求到其他服务
	// 	sConf := global.GVA_CONFIG.DataSources[common.DataSourceXR07]
	// 	forwardToOtherService(c, sConf.Url, strings.TrimPrefix(file, "/"+common.DataSourceXR07), "")
	// 	return
	// } else if strings.HasPrefix(file, "/"+common.DataSourceXR09) {
	// 	// 转发请求到其他服务
	// 	sConf := global.GVA_CONFIG.DataSources[common.DataSourceXR09]
	// 	forwardToOtherService(c, sConf.Url, strings.TrimPrefix(file, "/"+common.DataSourceXR09), "")
	// 	return
	// }

	if _, err := os.Open(file); err != nil {
		response.Abort404(c)
		return
	}

	name := filepath.Base(file)
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+url.QueryEscape(name))
	// //浏览器下载或预览
	// c.Header("Content-Disposition", "inline;filename="+name)
	c.Header("Content-Transfer-Encoding", "binary")
	c.File(file)
}

func (ctrl *AssetsController) Image(c *gin.Context) {
	path, _ := c.GetQuery("path")
	prdtType, _ := c.GetQuery("type")
	subType, _ := c.GetQuery("subType")
	fit, _ := c.GetQuery("fit")
	w, _ := c.GetQuery("w")
	h, _ := c.GetQuery("h")
	iw, _ := strconv.Atoi(w)
	ih, _ := strconv.Atoi(h)

	buff := bytes.NewBuffer([]byte{})
	imagePath := config.Get("app.assets_base_dir") + path
	fmt.Printf("Image path: %v", imagePath)
	imgType, err := file.GetImage(buff, imagePath, prdtType, subType, iw, ih, fit)
	if err == nil {
		c.Header("Content-Type", "image/"+imgType)
		io.Copy(c.Writer, buff)
	}
}
