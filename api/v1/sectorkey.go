package v1

import (
	"github.com/gin-gonic/gin"
	"star-server/model"
	"star-server/utils"
	"star-server/utils/errmsg"
)

func CreateSectorKey(context *gin.Context) {
	sectorName := context.Param("sector_name")
	_, code := model.UseNameGetSector(sectorName)
	if code == errmsg.ERROR {
		utils.ResponseOk(context, errmsg.SectorNotExist)
		return
	}
	key, err := utils.EncryptBcrypt(sectorName)
	if err == errmsg.ERROR {
		utils.ResponseOk(context, err)
		return
	}
	var sectorKey = model.SectorKey{Key: key, SectorName: sectorName}
	_, code = model.FindSectorKey(sectorKey)
	if code == errmsg.SUCCESS {
		utils.ResponseOk(context, errmsg.SectorKeyExist)
		return
	}
	utils.ResponseOk(context, model.CreateSectKey(&sectorKey))
}
