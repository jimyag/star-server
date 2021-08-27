package verify

import "star-server/model"

func MatchIdToken(id uint, openid string) bool {
	authentication := model.UseOpenidGetUid(openid)
	if id == authentication.Uid {
		return true
	}
	return false
}
