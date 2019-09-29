package action

import (
	"bytes"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/mojiajuzi/forum/model"
	"github.com/mojiajuzi/forum/service"
)

type fileUploadResult struct {
	path string
	err  error
}

//WebsiteSave 站点保存
func WebsiteSave(c *gin.Context) {
	resp := service.ForumResp{}

	//验证文件是否有效
	file, _ := c.FormFile("file")
	if file == nil {
		resp.Error(http.StatusBadRequest, "请上传文件", nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	URL := c.PostForm("url")
	var wg sync.WaitGroup
	urlStatus := make(chan bool)
	fileStatus := make(chan fileUploadResult)

	wg.Add(2)
	//验证地址是否有效
	go checkURLIsOk(URL, &wg, urlStatus)

	//验证文件上传是否有效
	go checkFileIsOk(file, &wg, fileStatus)

	urlOk := <-urlStatus
	fileOk := <-fileStatus

	if !urlOk {
		resp.Error(http.StatusBadRequest, "地址无效", nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	if fileOk.path == "" {
		resp.Error(http.StatusBadRequest, fileOk.err.Error(), nil)
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	wg.Wait()

	//写入数据
	w := model.Website{}
	w.Logo = fileOk.path
	w.URL = URL

	db := model.Db()

	db.Where(model.Website{URL: w.URL}).FirstOrCreate(w)
	resp.Success("添加成功", w)
	c.JSON(http.StatusOK, resp)
	return
}

//检测提交的网站是否正常
func checkURLIsOk(u string, wg *sync.WaitGroup, res chan bool) {
	defer wg.Done()

	resp, err := http.Get(u)
	if err != nil {
		res <- false
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		res <- true
	}
	res <- false
}

//checkFileIsOk 检查文件是否上传成功
func checkFileIsOk(file *multipart.FileHeader, wg *sync.WaitGroup, res chan fileUploadResult) {
	e := fileUploadResult{}
	data, err := file.Open()
	if err != nil {
		e.err = err
		res <- e
		return
	}
	defer data.Close()
	b, _ := ioutil.ReadAll(data)
	qiniu := service.QiniuSend{}
	qiniu.Init()
	fileURL, err := qiniu.UploadByBody(file.Filename, file.Size, bytes.NewReader(b))
	if err != nil {
		e.err = err
		res <- e
		return
	}

	e.path = fileURL
	res <- e
}
