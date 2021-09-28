package v1

import (
	"github.com/gin-gonic/gin"
	"star-server/model"
	"star-server/utils"
	"star-server/utils/errmsg"
)

func CreateNotice(context *gin.Context) {
	var notice model.Notice
	_ = context.ShouldBindJSON(&notice)
	err := model.AddNotice(&notice)
	if err == errmsg.ERROR {
		utils.RequestOk(context, err)
		return
	}
	utils.RequestDataOk(context, err, notice)
}

func GetNotice(context *gin.Context) {
	err, notices := model.GetNotice()
	utils.RequestDataOk(context, err, notices)
}
