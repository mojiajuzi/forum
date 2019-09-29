package action

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mojiajuzi/forum/service"
)

//ActivityIndex 获取活动列表
func ActivityIndex(c *gin.Context) {
	resp := service.ForumResp{}
	if uid, ok := c.Get("id"); ok {
		if userID, ok := uid.(int); ok {
			fmt.Println(userID)
		}
	} else {
		resp.Error(http.StatusBadRequest, "参数缺失", nil)
		c.JSON(200, resp)
		return
	}
}

//ActivitySave 创建新活动
func ActivitySave(c *gin.Context) {

}

//ActivityUpdate 更新活动
func ActivityUpdate(c *gin.Context) {

}

//ActivityDelete 删除活动
func ActivityDelete(c *gin.Context) {

}
