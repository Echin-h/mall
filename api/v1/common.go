package v1

import (
	"encoding/json"
	"fmt"
	conf "gin-mall/conf/sql"
	"gin-mall/pkg/e"
	"gin-mall/pkg/util/ctl"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// 区分错误 ， 区别一下是 解组错误（json） 还是 验证错误(validator), 使用类型断言来处理
func ErrorResponse(ctx *gin.Context, err error) *ctl.TrackedErrorResponse {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range ve {
			// 这边是将字段国际化
			field := conf.T(fmt.Sprintf("Field.%s", fieldError.Field))
			tag := conf.T(fmt.Sprintf("Tag.Valid.%s", fieldError.Tag))
			return ctl.RespError(ctx, err, fmt.Sprintf("%s%s", field, tag))
		}
	}

	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return ctl.RespError(ctx, err, "JSON类型不匹配")
	}

	return ctl.RespError(ctx, err, err.Error(), e.ERROR)
}
