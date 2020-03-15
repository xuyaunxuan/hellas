package api

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"hellas/common/e"
	"hellas/models"
	//"github.com/unknwon/com"
	"log"
	"net/http"
)
//获取多个文章标签
func GetTags(c *gin.Context) {
	data := make(map[string]interface{})
	valid := validation.Validation{}

	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		code = e.SUCCESS

		data["lists"] = models.GetUser()

	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}

//新增文章标签
func AddTag(c *gin.Context) {
}

//修改文章标签
func EditTag(c *gin.Context) {
}

//删除文章标签
func DeleteTag(c *gin.Context) {
}