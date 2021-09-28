package v1

import (
	"github.com/gin-gonic/gin"
	"star-server/model"
	"star-server/utils"
	"star-server/utils/errmsg"
)

func CreateSectorKey(context *gin.Context) {
	var data model.SectorKey
	_ = context.ShouldBindJSON(&data)
	key, err := utils.EncryptBcrypt(data.SectorName)
	if err == errmsg.ERROR {
		utils.RequestOk(context, err)
		return
	}
	data.Key = key
	_, code := model.FindSectorKey(data)
	if code == errmsg.SUCCESS {
		utils.RequestOk(context, errmsg.SectorKeyExist)
		return
	}
	utils.RequestOk(context, model.CreateSectKey(&data))
}
