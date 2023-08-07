package middleware

import (
	"fmt"
	"github.com/KingFarGrace/CollabSearch/server/entity"
	"github.com/KingFarGrace/CollabSearch/server/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// ValidateRegisterJSON is a middleware to validate user register data.
// @See RegisterJSON for detail validation rules.
func ValidateRegisterJSON() gin.HandlerFunc {
	var jsonObj = entity.RegisterJSON{}
	return validateJSON(jsonObj)
}

// ValidateLoginJSON is a middleware to validate user login data.
// This middleware won't be handled if the request pass the token verification middleware.
// @See LoginJSON for detail validation rules.
func ValidateLoginJSON() gin.HandlerFunc {
	var jsonObj = entity.LoginJSON{}
	return validateJSON(jsonObj)
}

func ValidateUpdateJSON() gin.HandlerFunc {
	var jsonObj = entity.User{}
	return validateJSON(jsonObj)
}

// validateJSON
// An encapsulation of different middleware func.
// Pass verified json obj or abort with return code http.StatusBadRequest.
// jsonObj <T> should only be entities in package entity.
func validateJSON[T entity.RegisterJSON | entity.LoginJSON | entity.User](jsonObj T) gin.HandlerFunc {
	return func(context *gin.Context) {
		if err := context.ShouldBindBodyWith(&jsonObj, binding.JSON); err != nil {
			fmt.Printf("%v", jsonObj)
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "Invalid request json.", "jsonObj": jsonObj})
			return
		}
		if msg := valid(jsonObj); msg != "" {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": msg})
			return
		}
		context.Set("jsonObj", jsonObj)
		context.Next()
	}
}

// valid is an encapsulation of validate process.
// Return errMsg if failed to validated json data, otherwise return ok (empty string).
func valid[T entity.RegisterJSON | entity.LoginJSON | entity.User](jsonObj T) string {
	validate := validator.New()
	ok := ""
	err := validate.Struct(jsonObj)
	if err != nil {
		errMsg := "Validate failed at:"
		for _, err := range err.(validator.ValidationErrors) {
			errMsg = util.Concat(errMsg, " ", err.StructNamespace())
		}
		return errMsg
	}
	return ok
}
