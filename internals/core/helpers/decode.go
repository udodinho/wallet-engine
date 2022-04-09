package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/udodinho/goprojects/wallet-engine/infrastructure/servererrors"
)

func Decode(ctx *gin.Context, obj interface{}) []string {
	err := ctx.ShouldBind(obj)
	if err != nil {
		var errs []string
		errVal, ok := err.(validator.ValidationErrors)
		if ok {
			for _, fieldErr := range errVal {
				errs = append(errs, servererrors.NewFieldError(fieldErr).String())
			}
		} else {
			errs = append(errs, "Internal server error: "+err.Error())
		}
		return errs
	}
	return nil
}
