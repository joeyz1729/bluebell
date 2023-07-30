package controller

import (
	"fmt"
	"reflect"
	"strings"

	"go.uber.org/zap"

	"github.com/YiZou89/bluebell/model"

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
	// 修改gin框架validator引擎的属性
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			// 获取json tag中的第一个
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		// 可以注册自定义校验方法：
		// v.RegisterStructValidation(SignupStructLevelValidation, model.SignupForm{})

		zhT := zh.New()
		enT := en.New()
		uni := ut.New(enT, zhT, enT)

		var ok bool
		trans, ok = uni.GetTranslator(locale)
		if !ok {
			err = fmt.Errorf("uni.GetTranslator(%s) failed", locale)
			zap.L().Error("init trans err", zap.Error(err))
			return
		}

		switch locale {
		case "en":
			err = enTrans.RegisterDefaultTranslations(v, trans)
		case "zh":
			err = zhTrans.RegisterDefaultTranslations(v, trans)
		default:
			err = enTrans.RegisterDefaultTranslations(v, trans)
		}
		if err != nil {
			zap.L().Error("init trans err", zap.Error(err))
		}
		zap.L().Info("init trans success")
		return
	}
	zap.L().Info("Failed to modify gin validator")
	return
}

// removeTopStruct 定义一个去掉结构体名称前缀的自定义方法，例如SignupFrom.password去掉前面的SF
func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}

// SignupStructLevelValidation 验证两次密码相同
// 也可以直接使用validator中的eqfield字段
func SignupStructLevelValidation(sl validator.StructLevel) {
	su := sl.Current().Interface().(model.SignupForm)
	if su.Password != su.ConfirmPassword {
		sl.ReportError(su.ConfirmPassword, "re_password", "ConfirmPassword", "eqfield", "password")
	}
}
