package action

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mojiajuzi/forum/model"
	"github.com/mojiajuzi/forum/service"
)

//WebsiteSave 站点保存
func WebsiteSave(c *gin.Context) {
	resp := service.ForumResp{}
	//验证地址是否有效
	URL := c.PostForm("url")
	if ok := checkURLIsOk(URL); !ok {
		resp.Error(http.StatusBadRequest, "地址无效", nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	//验证文件是否有效
	file, _ := c.FormFile("file")
	if file == nil {
		resp.Error(http.StatusBadRequest, "请上传文件", nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	data, err := file.Open()
	if err != nil {
		resp.Error(http.StatusBadRequest, "文件无效", nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	defer data.Close()
	b, _ := ioutil.ReadAll(data)
	qiniu := service.QiniuSend{}
	qiniu.Init()
	fileURL, err := qiniu.UploadByBody(file.Filename, file.Size, bytes.NewReader(b))
	if err != nil {
		resp.Error(http.StatusBadRequest, err.Error(), nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	//写入数据
	w := model.Website{}
	w.Logo = fileURL
	w.URL = URL

	db := model.Db()

	db.Where(model.Website{URL: w.URL}).FirstOrCreate(w)
	resp.Success("添加成功", w)
	c.JSON(http.StatusOK, resp)
	return
}

//检测提交的网站是否正常
func checkURLIsOk(u string) bool {
	resp, err := http.Get(u)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return true
	}
	return false
}
