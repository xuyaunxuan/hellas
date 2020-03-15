package controller

import (
	"hellas/common/e"
	"hellas/common/error"
	"hellas/dtos/common"
	"hellas/dtos/user"
	"hellas/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"log"
	"net/http"
)

// 新用户注册
func RegisterUser(c *gin.Context) {
	var registerParameter user.RegisterParameter
	//c.Query("username")
	// 参数验证
	if err := c.ShouldBindJSON(&registerParameter); err != nil {
		log.Printf("%+v", registerParameter)
		var errorDto common.ErrorDto
		errorDto.Message = e.MsgFlags[e.INVALID_PARAMS] +  err.Error()
		//errorDto.DebuggerError = err.Error()
		// 生成错误信息返回
		errorDto.Errors = error.CreateMessages(err.(validator.ValidationErrors))
		//c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.JSON(http.StatusBadRequest, errorDto)
		return
	}

	log.Printf("%+v", registerParameter)
	data := make(map[string]interface{})
	data["lists"] = models.CreateNewUser(registerParameter)
	code := e.INVALID_PARAMS
	c.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})


	//valid := validation.Validation{}

	//code := e.INVALID_PARAMS
	//if ! valid.HasErrors() {
	//	code = e.SUCCESS

	//} else {
	//	for _, err := range valid.Errors {
	//		log.Println(err.Key, err.Message)
	//	}
	//}


}