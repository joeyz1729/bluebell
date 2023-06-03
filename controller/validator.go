package controller

import (
	"fmt"
	"reflect"
	"strings"

	"zouyi/bluebell/model"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTrans "github.com/go-playground/validator/v10/translations/en"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
)

var trans ut.Translator

func InitTrans(locale string) (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		//v.RegisterStructValidation(SignupStructLevelValidation, model.SignupForm{})

		zhT := zh.New()
		enT := en.New()
		uni := ut.New(enT, zhT, enT)

		var ok bool

		trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}
		switch locale {
		case "en":
			err = enTrans.RegisterDefaultTranslations(v, trans)
		case "zh":
			err = zhTrans.RegisterDefaultTranslations(v, trans)
		default:
			err = enTrans.RegisterDefaultTranslations(v, trans)
		}
		return
	}
	return
}

// 定义一个去掉结构体名称前缀的自定义方法：
func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}

func SignupStructLevelValidation(sl validator.StructLevel) {
	su := sl.Current().Interface().(model.SignupForm)
	if su.Password != su.ConfirmPassword {
		sl.ReportError(su.ConfirmPassword, "re_password", "ConfirmPassword", "eqfield", "password")
	}
}
