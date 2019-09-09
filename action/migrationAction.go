package action

import (
	"github.com/gin-gonic/gin"
	"github.com/mojiajuzi/forum/model"
)

//Migrate 数据迁移
func Migrate(c *gin.Context) {
	model.Migrate()
}
