package verify

import (
	"star-server/model"
)

func MatchIdToken(id uint, openid string) bool {
	authentication, _ := model.UseOpenidGetAuth(openid)
	if id == authentication.Uid {
		return true
	}
	return false
}
