package api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"gogofly/global"
	"gogofly/utils"
	"reflect"
)

type BaseApi struct {
	Ctx    *gin.Context
	Errors error
	Logger *zap.SugaredLogger
}

func NewBaseApi() BaseApi {
	return BaseApi{
		Logger: global.Logger,
	}
}

type BuildRequestOption struct {
	Ctx               *gin.Context
	DTO               any
	BindParamsFromUri bool
}

func (b *BaseApi) BuildRequest(option BuildRequestOption) *BaseApi {
	var errResult error

	// 绑定请求上下文
	b.Ctx = option.Ctx

	// 绑定请求数据
	if option.DTO != nil {
		if option.BindParamsFromUri {
			errResult = b.Ctx.ShouldBindUri(option.DTO)
		} else {
			errResult = b.Ctx.ShouldBind(option.DTO)
		}

		if errResult != nil {
			errResult = b.parseValidateErrors(errResult, option.DTO)
			b.AddError(errResult)
			b.ClientFail(ResponseJson{
				Msg: b.GetError().Error(),
			})
		}
	}

	return b
}

func (b *BaseApi) AddError(err error) {
	b.Errors = utils.AppendError(b.Errors, err)
}

func (b *BaseApi) GetError() error {
	return b.Errors
}

func (b *BaseApi) parseValidateErrors(errs error, target any) error {
	var errResult error

	var errValidation validator.ValidationErrors
	ok := errors.As(errs, &errValidation)
	if !ok {
		return errs
	}

	// 通过反射获取指针指向的指定元素的类型对象
	fields := reflect.TypeOf(target).Elem()
	for _, fieldErr := range errValidation {
		field, _ := fields.FieldByName(fieldErr.Field())
		errMessageTag := fmt.Sprintf("%s_err", fieldErr.Tag())
		errMessage := field.Tag.Get(errMessageTag)
		if errMessage == "" {
			errMessage = field.Tag.Get("message")
		}

		if errMessage == "" {
			errMessage = fmt.Sprintf("%s: %s Error", fieldErr.Field(), fieldErr.Tag())
		}

		errResult = utils.AppendError(errResult, errors.New(errMessage))
	}

	return errResult
}

func (b *BaseApi) Success(resp ResponseJson) {
	Success(b.Ctx, resp)
}

func (b *BaseApi) ClientFail(resp ResponseJson) {
	ClientFail(b.Ctx, resp)
}

func (b *BaseApi) ServerFail(resp ResponseJson) {
	ServerFail(b.Ctx, resp)
}
