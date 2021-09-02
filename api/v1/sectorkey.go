package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"star-server/model"
	"star-server/utils"
	"star-server/utils/errmsg"
)

func CreateSectorKey(context *gin.Context) {
	var data model.SectorKey
	_ = context.ShouldBindJSON(&data)
	//fmt.Println("获得部门名称")
	//fmt.Println(data.SectorName)
	key, e := utils.EncryptBcrypt(data.SectorName)
	if e == errmsg.ERROR {
		//fmt.Println("加密错误")
		context.JSON(http.StatusOK, gin.H{
			"code": e,
			"msg":  errmsg.GetErrMsg(e),
		})
		context.Abort()
		return
	}
	data.Key = key
	_, code := model.FindSectorKey(data)
	if code == errmsg.SUCCESS {
		code = errmsg.SectorKeyExist
		context.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  errmsg.GetErrMsg(code),
		})
		context.Abort()
		return
	}
	code = model.CreateSectKey(&data)
	context.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  errmsg.GetErrMsg(code),
	})
}
