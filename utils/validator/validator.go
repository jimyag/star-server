package validator

import (
	"fmt"
	"github.com/go-playground/locales/zh_Hans_CN"
	unTrans "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"star-server/utils/errmsg"
)

// 所有的共用的情况下
func Validate(data interface{}) (string, int) {
	validate := validator.New()
	uni := unTrans.New(zh_Hans_CN.New())
	trans, _ := uni.GetTranslator("zh_Hans_CN")

	err := zhTrans.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println("err", err)
	}
	err = validate.Struct(data)
	if err != nil {
		for _, fieldError := range err.(validator.ValidationErrors) {
			return fieldError.Translate(trans), errmsg.ERROR
		}
	}
	return "", errmsg.SUCCESS

}
