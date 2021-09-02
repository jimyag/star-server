package utils

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
	"star-server/utils/errmsg"
	"time"
)

func GetOpenid(code string) (openid string, errmsg string) {
	var url = "https://api.weixin.qq.com/sns/jscode2session"
	url += "?appid=" + AgentAppid
	url += "&secret=" + AgentSecret
	url += "&js_code=" + code
	url += "&grant_type=authorization_code"
	url += "&connect_redirect=1"
	client := &http.Client{Timeout: 5 * time.Second}
	resp, _ := client.Get(url)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var maps = make(map[string]string)
	_ = json.Unmarshal(body, &maps)
	return maps["openid"], maps["errmsg"]
	//	map[openid:opJhB5J1PKQHnSg5b3CQSxy2n9Ew session_key:mr3MtcI6/U8VT9n2kDmcyQ==]
	//	map[errcode:40163 errmsg:code been used, rid: 6128c7e8-176287d4-2100cfe5]
}

func EncryptBcrypt(pwdStr string) (string, int) {
	pwd := []byte(pwdStr)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	if err != nil {
		return "", errmsg.ERROR
	}
	return string(hash), errmsg.SUCCESS
}

func PsdMatch(hashPwd string, plainPwd string) bool {
	byteHash := []byte(hashPwd)
	bytePlain := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePlain)
	return err == nil
}
